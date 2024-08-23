import { bus, enumName } from "/bus.js";
import * as buspb from "/pb/bus/bus_pb.js";
import * as chatpb from "/m/twitch/pb/chat_pb.js";
import * as commandpb from "/m/twitch/pb/command_pb.js";
import * as requestpb from "/m/twitch/pb/request_pb.js";
import * as twitchpb from "/m/twitch/pb/twitch_pb.js";
import { Timestamp } from "@bufbuild/protobuf";

const TOPIC_TWITCH_COMMAND = enumName(twitchpb.BusTopics, twitchpb.BusTopics.TWITCH_COMMAND);
const TOPIC_TWITCH_REQUEST = enumName(twitchpb.BusTopics, twitchpb.BusTopics.TWITCH_REQUEST);
const TOPIC_TWICH_CHAT_EVENT = enumName(twitchpb.BusTopics, twitchpb.BusTopics.TWITCH_CHAT_EVENT);
const TOPIC_TWITCH_CHAT_RECV = enumName(twitchpb.BusTopics, twitchpb.BusTopics.TWITCH_CHAT_RECV);
const TOPIC_TWITCH_CHAT_SEND = enumName(twitchpb.BusTopics, twitchpb.BusTopics.TWITCH_CHAT_SEND);

interface Status {
    status: chatpb.ChatStatus;
    detail: string;
}

class IRCStatus extends HTMLElement {
    private _span: HTMLSpanElement;

    constructor() {
        super();
        this.attachShadow({ mode: 'open' });
        this.shadowRoot.innerHTML = '<span title="Unknown">?</span>';
        this._span = this.shadowRoot.querySelector("span") as HTMLSpanElement;
    }

    connectedCallback() {
        bus.subscribe(TOPIC_TWICH_CHAT_EVENT, (msg) => this.handleChatEvent(msg));
        let msg = new buspb.BusMessage();
        msg.topic = TOPIC_TWITCH_REQUEST;
        msg.type = requestpb.MessageTypeRequest.TYPE_REQUEST_IRC_GET_STATUS_REQ;
        msg.message = new requestpb.IRCGetStatusRequest().toBinary();
        bus.waitForTopic(msg.topic, 5000)
            .then(() => bus.sendWithReply(msg, (reply: buspb.BusMessage) => {
                if (reply.type !== requestpb.MessageTypeRequest.TYPE_REQUEST_IRC_GET_STATUS_RESP) {
                    return;
                }
                this.status = requestpb.IRCGetStatusResponse.fromBinary(reply.message);
            }));
    }
    disconnectedCallback() {
        bus.unsubscribe(TOPIC_TWICH_CHAT_EVENT);
    }

    handleChatEvent(msg: buspb.BusMessage) {
        let ev = chatpb.ChatEvent.fromBinary(msg.message);
        switch (ev.type) {
            case chatpb.ChatEventType.EVENT_TYPE_STATUS:
                this.status = ev;
                break;
        }
    }
    set status(status: Status) {
        switch (status.status) {
            case chatpb.ChatStatus.UNKNOWN:
                this._span.innerText = '?';
                break
            case chatpb.ChatStatus.CONNECTED:
                this._span.innerHTML = '&#x2705;';
                break
            case chatpb.ChatStatus.ERROR:
                this._span.innerHTML = '&#x274C';
                break
            case chatpb.ChatStatus.OFF:
                this._span.innerHTML = '&#x25CE';
        }
        this._span.title = status.detail;
    }
}
customElements.define('twitch-irc-status', IRCStatus);

function chatMessageDiv(cmi: chatpb.ChatMessageIn): HTMLDivElement {
    let div = document.createElement('div') as HTMLDivElement;
    div.innerHTML= `
${new Date().toLocaleString()} | ${cmi.nick}> `;
    div.appendChild(document.createTextNode(cmi.text));
    return div;
}

const MESSAGE_COUNT = 100;
class IRCChat extends HTMLElement {
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
    }

    connectedCallback() {
        bus.subscribe(TOPIC_TWITCH_CHAT_RECV, (msg: buspb.BusMessage) => this.handleChatRecv(msg));
    }

    handleChatRecv(msg: buspb.BusMessage) {
        this.appendCMI(chatpb.ChatMessageIn.fromBinary(msg.message))
    }
    appendCMI(cmi: chatpb.ChatMessageIn) {
        this._container.appendChild(chatMessageDiv(cmi));
        while (this._container.children.length > MESSAGE_COUNT) {
            this._container.removeChild(this._container.children[0]);
        }
        this._container.scrollTop = this._container.scrollHeight;
    }

    private _send() {
        let text = this._messageInput.value;
        if (!text) {
            return;
        }
        let cmo = new chatpb.ChatMessageOut();
        cmo.text = text;
        let msg = new buspb.BusMessage();
        msg.topic = TOPIC_TWITCH_CHAT_SEND;
        msg.message = cmo.toBinary();
        bus.send(msg);
        this._messageInput.value = '';
        let cmi = new chatpb.ChatMessageIn();
        cmi.nick = '(You)';
        cmi.text = text;
        this.appendCMI(cmi);
    }
}
customElements.define('twitch-irc-chat', IRCChat);

class IRCConfig extends HTMLElement {
    private _enabledCheckbox: HTMLInputElement;
    private _profileSelect: HTMLSelectElement;
    private _prefixInput: HTMLInputElement;
    private _saveButton: HTMLButtonElement;
    private _cfg: chatpb.IRCConfig;

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
twitch-irc-chat {
    flex-grow: 100;
}
</style>
<fieldset>
<legend>Chat</legend>
<div>
    <section>
            <label for="twitch-irc-status">Status: </label>
            <twitch-irc-status id="twitch-irc-status"></twitch-irc-status>
    </section>
    <section>
            <label for="checkbox-enabled">Enabled</label>
            <input id="checkbox-enabled" type="checkbox" />
    </section>
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
<twitch-irc-chat></twitch-irc-chat>
</fieldset>
`;
        this._enabledCheckbox = this.shadowRoot.querySelector("#checkbox-enabled");
        this._profileSelect = this.shadowRoot.querySelector('#select-profile');
        this._prefixInput = this.shadowRoot.querySelector('#input-prefix');
        this._saveButton = this.shadowRoot.querySelector('#button-save');
        this._saveButton.onclick = () => this._save();
        this._cfg = new chatpb.IRCConfig();
    }

    connectedCallback() {
        bus.waitForTopic(TOPIC_TWITCH_REQUEST, 5000)
            .then(() => this._getProfiles());
    }

    get enabled(): boolean {
        return this._enabledCheckbox.checked;
    }
    set enabled(value: boolean) {
        this._enabledCheckbox.checked = value;
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
        msg.type = requestpb.MessageTypeRequest.TYPE_REQUEST_IRC_GET_CONFIG_REQ;
        msg.message = new requestpb.IRCGetConfigRequest().toBinary();
        bus.sendWithReply(msg, (reply: buspb.BusMessage) => {
            if (reply.error) {
                throw reply.error.detail;
            }
            this._saveButton.disabled = false;
            let gicr = requestpb.IRCGetConfigResponse.fromBinary(reply.message);
            if (!gicr.config) {
                return;
            }
            this._cfg = gicr.config;
            this.enabled = gicr.config.enabled;
            if (gicr.config.profile) {
                this.selectedProfile = gicr.config.profile;
            }
            this._prefixInput.value = gicr.config.messagePrefix;
        });
    }

    private _save() {
        this._saveButton.disabled = true;
        this._saveButton.innerText = 'Saving...';
        this._cfg.enabled = this.enabled;
        this._cfg.profile = this.selectedProfile;
        this._cfg.messagePrefix = this._prefixInput.value;
        let iscr = new commandpb.IRCSetConfigRequest();
        iscr.config = this._cfg;
        let msg = new buspb.BusMessage();
        msg.topic = TOPIC_TWITCH_COMMAND;
        msg.type = commandpb.MessageTypeCommand.TYPE_COMMAND_IRC_SET_CONFIG_REQ;
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
customElements.define('twitch-irc-config', IRCConfig);

export { IRCConfig };