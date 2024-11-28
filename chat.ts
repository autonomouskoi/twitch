import { bus, enumName } from "/bus.js";
import * as buspb from "/pb/bus/bus_pb.js";
import * as chatpb from "/m/twitch/pb/chat_pb.js";
import * as commandpb from "/m/twitch/pb/command_pb.js";
import * as eventsubpb from "/m/twitch/pb/eventsub_pb.js";
import * as requestpb from "/m/twitch/pb/request_pb.js";
import * as twitchpb from "/m/twitch/pb/twitch_pb.js";

const TOPIC_TWITCH_COMMAND = enumName(twitchpb.BusTopics, twitchpb.BusTopics.TWITCH_COMMAND);
const TOPIC_TWITCH_REQUEST = enumName(twitchpb.BusTopics, twitchpb.BusTopics.TWITCH_REQUEST);
const TOPIC_TWITCH_CHAT_REQUEST = enumName(twitchpb.BusTopics, twitchpb.BusTopics.TWITCH_CHAT_REQUEST);
const TOPIC_TWITCH_EVENTSUB_EVENT = enumName(twitchpb.BusTopics, twitchpb.BusTopics.TWITCH_EVENTSUB_EVENT);

interface ChatMessage {
    chatter: string;
    message: string;
}

function chatMessageDiv(cm: ChatMessage): HTMLDivElement {
    let div = document.createElement('div') as HTMLDivElement;
    div.innerHTML= `
${new Date().toLocaleString()} | ${cm.chatter}> `;
    div.appendChild(document.createTextNode(cm.message));
    return div;
}

const MESSAGE_COUNT = 100;
class LiveChat extends HTMLElement {
    private _container: HTMLElement;
    private _sendButton: HTMLButtonElement;
    private _messageInput: HTMLInputElement;

    constructor() {
        super();
        this.attachShadow({ mode: 'open' });
        this.shadowRoot.innerHTML = `
<style>
h3 {
    margin: 0px;
}
#mainContainer {
    display: flex;
    flex-direction: column;
    align-items: stretch;
}
section {
    background-color: eee;
    border: 1px solid black;
    height: 10rem;
    overflow: auto;
}
</style>
<div id="mainContainer">
    <h3>Live</h3>
    <section>
    </section>
    <div>
        <label for="input-send-message">Send message to chat</label>
        <input id="input-send-message" type="text" size=50 />
        <button id="button-send-message">Send</button>
    </div>
</div>
`;
        this._container = this.shadowRoot.querySelector('section');
        this._messageInput = this.shadowRoot.querySelector('#input-send-message');
        this._messageInput.addEventListener("keyup", (ev) => {
            if (ev.key === "Enter") {
                this._send();
            }
        });
        this._sendButton = this.shadowRoot.querySelector('#button-send-message');
        this._sendButton.onclick = () => this._send();
        bus.subscribe(TOPIC_TWITCH_EVENTSUB_EVENT, (msg: buspb.BusMessage) => this.handleChatRecv(msg));
    }

    handleChatRecv(msg: buspb.BusMessage) {
        switch (msg.type) {
            case eventsubpb.MessageTypeEventSub.TYPE_CHANNEL_CHAT_MESSAGE:
                let ccm = eventsubpb.EventChannelChatMessage.fromBinary(msg.message);
                let cm = {
                    chatter: ccm.chatter.login,
                    message: ccm.message.text,
                };
                this.appendCMI(cm);
                break;
            default:
                throw(`Unhandled message type ${msg.getType()}`);
        }
    }
    appendCMI(cm: ChatMessage) {
        this._container.appendChild(chatMessageDiv(cm));
        while (this._container.children.length > MESSAGE_COUNT) {
            this._container.removeChild(this._container.children[0]);
        }
        this._container.scrollTop = this._container.scrollHeight;
    }

    private _send() {
        this._sendButton.disabled = true;
        let text = this._messageInput.value;
        if (!text) {
            return;
        }
        let cmo = new chatpb.TwitchChatRequestSendRequest();
        cmo.text = text;
        let msg = new buspb.BusMessage();
        msg.topic = TOPIC_TWITCH_CHAT_REQUEST;
        msg.type = chatpb.MessageTypeTwitchChatRequest.TWITCH_CHAT_REQUEST_TYPE_SEND_REQ;
        msg.message = cmo.toBinary();
        bus.sendAnd(msg).then((reply) => {
            this._messageInput.value = '';
        }).catch((e) => console.log(`error sending chat message: ${JSON.stringify(e)}`))
        .finally(() => { this._sendButton.disabled = false; });
    }
}
customElements.define('twitch-chat', LiveChat);

class ChatConfig extends HTMLElement {
    private _profileSelect: HTMLSelectElement;
    private _prefixInput: HTMLInputElement;
    private _saveButton: HTMLButtonElement;
    private _cfg: chatpb.ChatConfig;

    constructor() {
        super();
        this.attachShadow({ mode: 'open' });
        this.shadowRoot.innerHTML = `
<style>
fieldset {
    display: flex;
    flex-direction: row;
    gap: 5px;
}
twitch-chat {
    flex-grow: 100;
}
</style>
<fieldset>
<legend>Chat</legend>
<div>
    <section>
            <label for="select-profile">Using Account</label>
            <select id="select-profile"></select>
    </section>
    <section>
            <label for="input-prefix">Chat Message Prefix</label>
            <input id="input-prefix" type="text" />
    </section>
    <button id="button-save" disabled>Save</button>
</div>
<twitch-chat></twitch-chat>
</fieldset>
`;
        this._profileSelect = this.shadowRoot.querySelector('#select-profile');
        this._prefixInput = this.shadowRoot.querySelector('#input-prefix');
        this._saveButton = this.shadowRoot.querySelector('#button-save');
        this._saveButton.onclick = () => this._save();
        this._cfg = new chatpb.ChatConfig();
    }

    connectedCallback() {
        bus.waitForTopic(TOPIC_TWITCH_REQUEST, 5000)
            .then(() => this._getProfiles());
    }

    get selectedProfile(): string {
        let selectedIdx = this._profileSelect.selectedIndex;
        let selected = this._profileSelect.children[selectedIdx] as HTMLOptionElement;
        return selected.value;
    }
    set selectedProfile(value: string) {
        let options = this._profileSelect.options;
        for (let i = 0; i < options.length; i++) {
            if (options[i].label === value) {
                this._profileSelect.selectedIndex = i;
                break;
            }
        }
    }

    private _getProfiles() {
        let msg = new buspb.BusMessage();
        msg.topic = TOPIC_TWITCH_REQUEST;
        msg.type = requestpb.MessageTypeRequest.TYPE_REQUEST_LIST_PROFILES_REQ;
        msg.message = new requestpb.ListProfilesRequest().toBinary();
        bus.sendWithReply(msg, (reply: buspb.BusMessage) => {
            if (reply.error) {
                throw reply.error.detail;
            }
            let lpr = requestpb.ListProfilesResponse.fromBinary(reply.message);
            this._profileSelect.textContent = '';
            lpr.names.toSorted().forEach((name) => {
                let option = document.createElement('option') as HTMLOptionElement;
                option.value = name;
                option.innerText = name;
                this._profileSelect.appendChild(option);
            });
            this._getConfig();
        });
    }

    private _getConfig() {
        let msg = new buspb.BusMessage();
        msg.topic = TOPIC_TWITCH_REQUEST;
        msg.type = requestpb.MessageTypeRequest.TYPE_REQUEST_CHAT_GET_CONFIG_REQ;
        msg.message = new requestpb.ChatGetConfigRequest().toBinary();
        bus.sendWithReply(msg, (reply: buspb.BusMessage) => {
            if (reply.error) {
                throw reply.error.detail;
            }
            this._saveButton.disabled = false;
            let gicr = requestpb.ChatGetConfigResponse.fromBinary(reply.message);
            if (!gicr.config) {
                return;
            }
            this._cfg = gicr.config;
            if (gicr.config.profile) {
                this.selectedProfile = gicr.config.profile;
            }
            this._prefixInput.value = gicr.config.messagePrefix;
        });
    }

    private _save() {
        this._saveButton.disabled = true;
        this._saveButton.innerText = 'Saving...';
        this._cfg.enabled = true;
        this._cfg.profile = this.selectedProfile;
        this._cfg.messagePrefix = this._prefixInput.value;
        let iscr = new commandpb.ChatSetConfigRequest();
        iscr.config = this._cfg;
        let msg = new buspb.BusMessage();
        msg.topic = TOPIC_TWITCH_COMMAND;
        msg.type = commandpb.MessageTypeCommand.TYPE_COMMAND_CHAT_SET_CONFIG_REQ;
        msg.message = iscr.toBinary();
        bus.waitForTopic(TOPIC_TWITCH_COMMAND, 5000)
            .then(() => bus.sendWithReply(msg, (reply: buspb.BusMessage) => {
                if (reply.error) {
                    this._saveButton.innerText = 'FAILED';
                    throw reply.error;
                }
                this._saveButton.innerText = 'Saved!';
                this._getProfiles();
            })
            )
            .finally(() => setInterval(() => {
                this._saveButton.innerText = 'Save';
                this._saveButton.disabled = false;
            }, 3000));
    }
}
customElements.define('twitch-chat-config', ChatConfig);

export { ChatConfig };