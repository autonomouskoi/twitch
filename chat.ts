import { bus, enumName } from "/bus.js";
import * as buspb from "/pb/bus/bus_pb.js";
import * as chatpb from "/m/twitch/pb/chat_pb.js";
import * as eventsubpb from "/m/twitch/pb/eventsub_pb.js";
import * as twitchpb from "/m/twitch/pb/twitch_pb.js";
import { UpdatingControlPanel } from '/tk.js';
import { ChatCfg } from './controller.js';
import { ProfileSelector } from '/m/twitch/profiles.js';

const TOPIC_TWITCH_CHAT_REQUEST = enumName(twitchpb.BusTopics, twitchpb.BusTopics.TWITCH_CHAT_REQUEST);
const TOPIC_TWITCH_EVENTSUB_EVENT = enumName(twitchpb.BusTopics, twitchpb.BusTopics.TWITCH_EVENTSUB_EVENT);

// a chat message to be displayed. May come from the bus or be generated by the UI
interface ChatMessage {
    chatter: string;
    message: string;
}

// generate a div to display a chat message
function chatMessageDiv(cm: ChatMessage): HTMLDivElement {
    let div = document.createElement('div') as HTMLDivElement;
    div.innerHTML = `
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
        this._sendButton.addEventListener('click', () => this._send());
        bus.subscribe(TOPIC_TWITCH_EVENTSUB_EVENT, (msg) => this.handleChatRecv(msg));
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

let help = document.createElement('div');
help.innerHTML = `
<p>
The <em>Using Account</em> selector allows you to choose which connected Twitch
account to use for chat. You must hit the <em>Save</em> button to save this setting
and it won't take effect until the next time chat connects.
</p>

<p>
The <em>Chat Message Prefix</em> field allows you to specify a prefix for every
message the <code>twitch</code> module sends to chat. For example, if the prefix
is set to <code>[bot] </code> and a module sends the message <code>how are you?</code>,
The message will show in chat as <code>[bot] how are you?</code>. This can help
chat participants understand that the message came from a bot and not a human.
This setting is saved by hitting the <em>Save</em> button and takes effect immediately.
</p>
`;

class ChatConfig extends UpdatingControlPanel<chatpb.ChatConfig> {
    private _profile: ProfileSelector;
    private _prefix: HTMLInputElement;

    constructor(cfg: ChatCfg) {
        super({ title: 'Chat Config', help, data: cfg })

        this.innerHTML = `
<form method="dialog">
<section class="grid grid-2-col">

<label for="select-profile">Using Account</label>
<twitch-profile-select id="select-profile"></twitch-profile-select>

<label for="input-prefix">Chat Message Prefix</label>
<input id="input-prefix" type="text" />

<input type="submit" value="Save" />

</section>
</form>

<section>
<twitch-chat></twitch-chat>
</section>
`;

        this._profile = this.querySelector('#select-profile');
        this._prefix = this.querySelector('input');

        this.querySelector('form').addEventListener('submit', () => {
            let cfg = this.last.clone();
            cfg.messagePrefix = this._prefix.value;
            cfg.profile = this._profile.value;
            this.save(cfg);
        });
    }

    update(cfg: chatpb.ChatConfig) {
        this._prefix.value = cfg.messagePrefix;
        this._profile.value = cfg.profile;
    }
}
customElements.define('twitch-chat-config', ChatConfig, { extends: 'fieldset' });

export { ChatConfig };