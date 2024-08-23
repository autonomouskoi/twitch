import { bus, enumName } from "/bus.js";
import * as buspb from "/pb/bus/bus_pb.js";
import * as commandpb from "/m/twitch/pb/command_pb.js";
import * as requestpb from "/m/twitch/pb/request_pb.js";
import * as twitchpb from "/m/twitch/pb/twitch_pb.js";

const TOPIC_COMMAND = enumName(twitchpb.BusTopics, twitchpb.BusTopics.TWITCH_COMMAND); 
const TOPIC_REQUEST = enumName(twitchpb.BusTopics, twitchpb.BusTopics.TWITCH_REQUEST);

function profileButton(name: string, onsuccess: (name: string) => void): HTMLButtonElement {
    let button = document.createElement('button') as HTMLButtonElement;
    button.innerHTML = '&#x1F5D1; ' + name;
    button.onclick = () => {
        if (!window.confirm(`Really delete connection to ${name}?`)) {
            return;
        }
        let msg = new buspb.BusMessage();
        msg.topic = TOPIC_COMMAND;
        msg.type = commandpb.MessageTypeCommand.TYPE_COMMAND_DELETE_PROFILE_REQ;
        let dpr = new commandpb.DeleteProfileRequest();
        dpr.name = name;
        msg.message = dpr.toBinary();
        bus.sendWithReply(msg, (reply: buspb.BusMessage) => {
            if (reply.error) {
                alert(`error deleting ${name}: ${reply.error.detail}`);
                return;
            }
            onsuccess(name);
        });
    };
    return button;
}

class NewProfile extends HTMLElement {
    private _button: HTMLButtonElement;
    private _anchor: HTMLAnchorElement;

    constructor() {
        super();
        this.attachShadow({mode:'open'});
        this.shadowRoot.innerHTML = `
<style>
</style>
<section>
    <button>New Connection</button>
    <a href="#"></a>
</section>
`;
        this._button = this.shadowRoot.querySelector("button") as HTMLButtonElement;
        this._anchor = this.shadowRoot.querySelector("a") as HTMLAnchorElement;

        this._button.onclick = () => {
            let msg = new buspb.BusMessage();
            msg.topic = TOPIC_COMMAND;
            msg.type = commandpb.MessageTypeCommand.TYPE_COMMAND_GET_OAUTH_URL_REQ;
            msg.message = new commandpb.GetOAuthURLRequest().toBinary();
            bus.sendWithReply(msg, (reply: buspb.BusMessage) => {
                if (reply.error) {
                    this._anchor.innerText = `error: ${reply.error.detail}`;
                    return;
                }
                let gour = commandpb.GetOAuthURLResponse.fromBinary(reply.message);
                this._anchor.href = gour.url;
                this._anchor.innerText = 'Click this, or copy this into a browser logged into the desired Twitch account';
            });
        };
    }
}
customElements.define('twitch-new-profile', NewProfile);

class Profiles extends HTMLElement {
    constructor() {
        super();
        this.attachShadow({mode:'open'});
        bus.waitForTopic(TOPIC_REQUEST, 5000).then(() => this.update());
    }

    update() {
        this.shadowRoot.innerHTML = `
<fieldset>
<legend>Connected Twitch Accounts</legend>
<div id="connected"></div>
<twitch-new-profile></twitch-new-profile>
</fieldset>
`;
        let msg = new buspb.BusMessage();
        msg.topic = TOPIC_REQUEST;
        msg.type = requestpb.MessageTypeRequest.TYPE_REQUEST_LIST_PROFILES_REQ;
        msg.message = new requestpb.ListProfilesRequest().toBinary();
        bus.sendWithReply(msg, (reply: buspb.BusMessage) => {
            let connectedContainer = this.shadowRoot.querySelector("#connected") as HTMLElement;
            if (reply.error)  {
                connectedContainer.innerText = `error getting profiles: ${reply.error.detail}`;
                return;
            }
            let lpr = requestpb.ListProfilesResponse.fromBinary(reply.message);
            lpr.names.toSorted().forEach((name: string) => {
                let button = profileButton(name, () => this.update());
                connectedContainer.appendChild(button);
            });
        });
    }
}
customElements.define('twitch-profiles', Profiles);

export { Profiles };