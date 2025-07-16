import { bus, enumName } from "/bus.js";
import * as buspb from "/pb/bus/bus_pb.js";
import * as eventsubpb from "/m/twitch/pb/eventsub_pb.js";
import * as requestpb from "/m/twitch/pb/request_pb.js";
import * as twitchpb from "/m/twitch/pb/twitch_pb.js";
import { UpdatingControlPanel } from "/tk.js";
import { EventCfg } from './controller.js';
import { ProfileSelector } from '/m/twitch/profiles.js';

const TOPIC_TWITCH_REQUEST = enumName(twitchpb.BusTopics, twitchpb.BusTopics.TWITCH_REQUEST);
const TOPIC_TWICH_EVENTSUB_EVENT = enumName(twitchpb.BusTopics, twitchpb.BusTopics.TWITCH_EVENTSUB_EVENT);

interface Status {
    status: eventsubpb.EventSubStatus;
    detail: string;
}

class EventSubStatus extends HTMLElement {
    constructor() {
        super();
        this.title = 'Unknown';
        this.innerHTML = '?';
    }

    connectedCallback() {
        bus.subscribe(TOPIC_TWICH_EVENTSUB_EVENT, (msg) => this.handleEventSubEvent(msg));
        let msg = new buspb.BusMessage();
        msg.topic = TOPIC_TWITCH_REQUEST;
        msg.type = requestpb.MessageTypeRequest.TYPE_REQUEST_EVENT_GET_STATUS_REQ;
        msg.message = new requestpb.EventSubGetStatusRequest().toBinary();
        bus.waitForTopic(msg.topic, 5000)
            .then(() => bus.sendAnd(msg))
            .then((reply) => {
                this.status = requestpb.EventSubGetStatusResponse.fromBinary(reply.message);
            });
    }

    disconnectedCallback() {
        bus.unsubscribe(TOPIC_TWICH_EVENTSUB_EVENT);
    }

    handleEventSubEvent(msg: buspb.BusMessage) {
        switch (msg.type) {
            case eventsubpb.MessageTypeEventSub.TYPE_EVENTSUB_EVENT:
                let ev = eventsubpb.EventSubStatusEvent.fromBinary(msg.message);
                this.status = ev;
                break;
        }
    }
    set status(status: Status) {
        switch (status.status) {
            case eventsubpb.EventSubStatus.UNKNOWN:
                this.innerText = '?';
                break;
            case eventsubpb.EventSubStatus.CONNECTED:
                this.innerHTML = '&#x2705;';
                break;
            case eventsubpb.EventSubStatus.ERROR:
                this.innerHTML = '&#x274C';
                break;
            case eventsubpb.EventSubStatus.OFF:
                this.innerHTML = '&#x25CE';
        }
        this.title = status.detail;
    }
}
customElements.define('twitch-es-status', EventSubStatus);

let help = document.createElement('div');
help.innerHTML = `
<p>
The <em>Status</em> field shows whether or not AK is connected to Twitch to
receive events.
</p>

<p>
The <em>Using Account</em> selector allows you to choose which connected Twitch
account to use for receiving events and chat messages. You must hit the <em>Save</em> button to
save this setting and it won't take effect until the next time chat connects.
</p>

<p>
The <em>Enabled</em> checkbox sets whether or not receiving events is enabled.
You must hit the <em>Save</em> button for the changes to take effect.
</p>

<p>
The <em>Log Raw Events</em> checkbox sets whether or not AK will maintain a log
of all Twitch events it has received. These logs can be found by clicking AK in
the system tray, selecting <em>Open data folder</em>, going up one folder, then
to <em>Twitch</em> -> <em>logs</em>.
</p>
`;

class EventSubConfig extends UpdatingControlPanel<eventsubpb.EventSubConfig> {
    private _enabled: HTMLInputElement;
    private _log: HTMLInputElement;
    private _profile: ProfileSelector;

    constructor(cfg: EventCfg) {
        super({ title: 'EventSub Config', help, data: cfg })
        const profileTitle = 'Receive Twitch events and chat messages via this profile';
        this.innerHTML = `
<form method="dialog">
<section class="grid grid-2-col">

<label title="Connection status">Status</label>
<twitch-es-status></twitch-es-status>

<label for="select-profile" title="${profileTitle}">Using Account</label>

<label for="enabled" title="Connect to Twitch events and chat">Enabled</label>
<input type="checkbox" id="enabled" title="Connect to Twitch events and chat" />

<label for="log" title="Keep a log of all Twitch events">Log Raw Events</label>
<input type="checkbox" id="log" title="Keep a log of all Twitch events" />

<input type="submit" value="Save" />

</section>
</form>
`;

        this._enabled = this.querySelector('#enabled');
        this._log = this.querySelector('#log');
        this._profile = new ProfileSelector();
        this._profile.id = 'select-profile';
        this._profile.title = profileTitle;
        this.querySelector('label[for="select-profile"]').after(this._profile);

        this.querySelector('form').addEventListener('submit', () => {
            let cfg = this.last.clone();
            cfg.enabled = this._enabled.checked;
            cfg.logEvents = this._log.checked;
            cfg.profile = this._profile.value;
            this.save(cfg);
        });
    }

    update(cfg: eventsubpb.EventSubConfig) {
        this._enabled.checked = cfg.enabled;
        this._log.checked = cfg.logEvents;
        this._profile.selected = cfg.profile;
    }
}
customElements.define('twitch-es-config', EventSubConfig, { extends: 'fieldset' });

export { EventSubConfig };