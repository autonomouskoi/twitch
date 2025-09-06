import { bus, enumName } from "/bus.js";
import * as buspb from "/pb/bus/bus_pb.js";
import * as commandpb from "/m/twitch/pb/command_pb.js";
import * as requestpb from "/m/twitch/pb/request_pb.js";
import * as twitchpb from "/m/twitch/pb/twitch_pb.js";
import { ControlPanel, VoidPromise } from "/tk.js";

const TOPIC_COMMAND = enumName(twitchpb.BusTopics, twitchpb.BusTopics.TWITCH_COMMAND);
const TOPIC_REQUEST = enumName(twitchpb.BusTopics, twitchpb.BusTopics.TWITCH_REQUEST);

const TWO_WEEKS = 60 * 60 * 24 * 14;

function profileSymbol(profile: requestpb.ListProfilesResponse_ProfileListing): string {
    let expires = Number(profile.expires);
    if (expires == 0) {
        return '¯\\_(ツ)_/¯';
    }
    let now = new Date().getTime() / 1000;
    if (expires <= now) {
        return '&#x274C';
    }
    if ((expires - now) < TWO_WEEKS) {
        return '&#x23F0';
    }
    return '&#x2705;';
}

// create a button that will prompt to delete a profile
function profileButton(profile: requestpb.ListProfilesResponse_ProfileListing, onsuccess: (name: string) => void): HTMLButtonElement {
    let button = document.createElement('button') as HTMLButtonElement;
    button.innerHTML = `&#x1F5D1; ${profile.name} ${profileSymbol(profile)}`;
    button.addEventListener('click', () => {
        if (!window.confirm(`Really delete connection to ${name}?`)) {
            return;
        }
        let msg = new buspb.BusMessage();
        msg.topic = TOPIC_COMMAND;
        msg.type = commandpb.MessageTypeCommand.TYPE_COMMAND_DELETE_PROFILE_REQ;
        let dpr = new commandpb.DeleteProfileRequest();
        dpr.name = profile.name;
        msg.message = dpr.toBinary();
        bus.sendWithReply(msg, (reply) => {
            if (reply.error) {
                alert(`error deleting ${name}: ${reply.error.detail}`);
                return;
            }
            onsuccess(profile.name);
        });
    });
    return button;
}

// NewProfile generates an OAuth enrollment link
class NewProfile extends HTMLElement {
    private _button: HTMLButtonElement;
    private _anchor: HTMLAnchorElement;

    constructor() {
        super();
        this.attachShadow({ mode: 'open' });
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

        this._button.addEventListener('click', () => {
            let msg = new buspb.BusMessage();
            msg.topic = TOPIC_COMMAND;
            msg.type = commandpb.MessageTypeCommand.TYPE_COMMAND_GET_OAUTH_URL_REQ;
            msg.message = new commandpb.GetOAuthURLRequest().toBinary();
            bus.sendWithReply(msg, (reply) => {
                if (reply.error) {
                    this._anchor.innerText = `error: ${reply.error.detail}`;
                    return;
                }
                let gour = commandpb.GetOAuthURLResponse.fromBinary(reply.message);
                this._anchor.href = gour.url;
                this._anchor.innerText = 'Click this, or copy this into a browser logged into the desired Twitch account';
            });
        });
    }
}
customElements.define('twitch-new-profile', NewProfile);

let profilesHelp = document.createElement('div');
profilesHelp.innerHTML = `
<p>
The <code>twitch</code> module can have tokens from multiple accounts. The account
used to talk to the Twitch API can be different than the account used to interact
with chat.
</p>

<p>
The <em>New Connection</em> button will create a link to get a token from Twitch.
The link can be clicked or copyed and pasted into a different browser. The link
will request a token from Twitch tied to the Twitch account the browser is logged
into. To get a token tied to a different account, you may need to log into that
account in a different browser or a Private/Incognito window. The link will become
invalid if AutonomousKoi or the <code>twitch</code> module is restarted.
</p>

<p>
A button will be shown for each account the <code>twitch</code> module has a token
for. Clicking the button will show a confirmation to remove the token. If you confirm
the deletion, the <code>twitch</code> module will tell Twitch to invalidate the
token and then it will delete its copy of the token.
</p>

<p>
A symbol is displayed next to the profile name indicating its status:
<dl>
    <dt>${profileSymbol(new requestpb.ListProfilesResponse_ProfileListing())}</dt>
    <dd>The profile was created before AK saved the expiration :(</dd>

    <dt>${profileSymbol(new requestpb.ListProfilesResponse_ProfileListing({ name: 'soon', expires: BigInt(Math.floor(new Date().getTime() / 1000) + 5) }))}</dt>
    <dd>The token will expire in less than two weeks</dd>

    <dt>${profileSymbol(new requestpb.ListProfilesResponse_ProfileListing({ name: 'expired', expires: BigInt(Math.floor(new Date().getTime() / 1000) - 5) }))}</dt>
    <dd>The token is expired</dd>

    <dt>${profileSymbol(new requestpb.ListProfilesResponse_ProfileListing({ name: 'good', expires: BigInt(Math.floor(new Date().getTime() / 1000) + 5 + TWO_WEEKS) }))}</dt>
    <dd>The token isn't expiring soon</dd>
</dl>
</p>
`;

class Profiles extends ControlPanel {
    constructor() {
        super({ title: 'Connected Twitch Accounts', help: profilesHelp });
        bus.waitForTopic(TOPIC_REQUEST, 5000).then(() => this.update());
    }

    update() {
        this.innerHTML = `
<div id="connected"></div>
<twitch-new-profile></twitch-new-profile>
`;
        let msg = new buspb.BusMessage();
        msg.topic = TOPIC_REQUEST;
        msg.type = requestpb.MessageTypeRequest.TYPE_REQUEST_LIST_PROFILES_REQ;
        msg.message = new requestpb.ListProfilesRequest().toBinary();
        bus.sendWithReply(msg, (reply) => {
            let connectedContainer = this.querySelector("#connected") as HTMLElement;
            if (reply.error) {
                connectedContainer.innerText = `error getting profiles: ${reply.error.detail}`;
                return;
            }
            let lpr = requestpb.ListProfilesResponse.fromBinary(reply.message);
            lpr.profiles
                .toSorted((a, b) => a.name.localeCompare(b.name))
                .forEach((profile) => {
                    let button = profileButton(profile, () => this.update());
                    connectedContainer.appendChild(button);
                });
        });
    }
}
customElements.define('twitch-profiles', Profiles, { extends: 'fieldset' });

class ProfileSelector extends HTMLSelectElement {

    private _ready = new VoidPromise();

    constructor() {
        super();

        this.update();
    }

    update() {
        bus.waitForTopic(TOPIC_REQUEST, 5000)
            .then(() => bus.sendAnd(new buspb.BusMessage({
                topic: TOPIC_REQUEST,
                type: requestpb.MessageTypeRequest.TYPE_REQUEST_LIST_PROFILES_REQ,
                message: new requestpb.ListProfilesRequest().toBinary(),
            }))).then((reply) => {
                let resp = requestpb.ListProfilesResponse.fromBinary(reply.message);
                this.textContent = '';
                resp.profiles
                    .toSorted((a, b) => a.name.localeCompare(b.name))
                    .forEach((profile) => {
                        let option = document.createElement('option');
                        option.value = profile.name;
                        option.innerText = profile.name;
                        option.innerHTML += ` ${profileSymbol(profile)}`;
                        this.appendChild(option);
                    });
                this._ready.resolve();
            }).catch((e) => this._ready.reject(e));
    }

    set selected(value: string) {
        this._ready.wait.then(() => this.value = value);
    }
}
customElements.define('twitch-profile-select', ProfileSelector, { extends: 'select' });

export { Profiles, ProfileSelector };