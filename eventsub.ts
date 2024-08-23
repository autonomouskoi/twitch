import { bus, enumName } from "/bus.js";
import * as buspb from "/pb/bus/bus_pb.js";
import * as eventsubpb from "/m/twitch/pb/eventsub_pb.js";
import * as commandpb from "/m/twitch/pb/command_pb.js";
import * as requestpb from "/m/twitch/pb/request_pb.js";
import * as twitchpb from "/m/twitch/pb/twitch_pb.js";

const TOPIC_TWITCH_COMMAND = enumName(twitchpb.BusTopics, twitchpb.BusTopics.TWITCH_COMMAND);
const TOPIC_TWITCH_REQUEST = enumName(twitchpb.BusTopics, twitchpb.BusTopics.TWITCH_REQUEST);
const TOPIC_TWICH_EVENTSUB_EVENT = enumName(twitchpb.BusTopics, twitchpb.BusTopics.TWITCH_EVENTSUB_EVENT);

interface Status {
    status: eventsubpb.EventSubStatus;
    detail: string;
}

class EventSubStatus extends HTMLElement {
    private _span: HTMLSpanElement;

    constructor() {
        super();
        this.attachShadow({ mode: 'open' });
        this.shadowRoot.innerHTML = '<span title="Unknown">?</span>';
        this._span = this.shadowRoot.querySelector("span") as HTMLSpanElement;
    }

    connectedCallback() {
        bus.subscribe(TOPIC_TWICH_EVENTSUB_EVENT, (msg) => this.handleEventSubEvent(msg));
        let msg = new buspb.BusMessage();
        msg.topic = TOPIC_TWITCH_REQUEST;
        msg.type = requestpb.MessageTypeRequest.TYPE_REQUEST_EVENT_GET_STATUS_REQ;
        msg.message = new requestpb.EventSubGetStatusRequest().toBinary();
        bus.waitForTopic(msg.topic, 5000)
            .then(() => bus.sendWithReply(msg, (reply: buspb.BusMessage) => {
                if (reply.type !== requestpb.MessageTypeRequest.TYPE_REQUEST_EVENT_GET_STATUS_RESP) {
                    return;
                }
                this.status = requestpb.EventSubGetStatusResponse.fromBinary(reply.message);
            }));
    }
    disconnectedCallback() {
        bus.unsubscribe(TOPIC_TWICH_EVENTSUB_EVENT);
    }

    handleEventSubEvent(msg: buspb.BusMessage) {
        switch (msg.type) {
            case eventsubpb.MessageTypeEventSub.TYPE_EVENTSUB_EVENT:
                let ev = eventsubpb.EventSubStatusEvent.fromBinary(msg.message);
                this.status = ev;
                break
        }
    }
    set status(status: Status) {
        switch (status.status) {
            case eventsubpb.EventSubStatus.UNKNOWN:
                this._span.innerText = '?';
                break
            case eventsubpb.EventSubStatus.CONNECTED:
                this._span.innerHTML = '&#x2705;';
                break
            case eventsubpb.EventSubStatus.ERROR:
                this._span.innerHTML = '&#x274C';
                break
            case eventsubpb.EventSubStatus.OFF:
                this._span.innerHTML = '&#x25CE';
        }
        this._span.title = status.detail;
    }
}
customElements.define('twitc-es-status', EventSubStatus);

class EventSubConfig extends HTMLElement {
    private _enabledCheckbox: HTMLInputElement;
    private _profileSelect: HTMLSelectElement;
    private _saveButton: HTMLButtonElement;
    private _cfg: eventsubpb.EventSubConfig;

    constructor() {
        super();
        this.attachShadow({ mode: 'open' });
        this.shadowRoot.innerHTML = `
<fieldset>
<legend>Channel Events (Subs, Follows, etc)</legend>

<section>
        <label for="twitc-es-status">Status: </label>
        <twitc-es-status id="twitc-es-status"></twitc-es-status>
</section>
<section>
        <label for="checkbox-enabled">Enabled</label>
        <input id="checkbox-enabled" type="checkbox" />
</section>
<section>
        <label for="select-profile">Using Account</label>
        <select id="select-profile"></select>
</section>
<button id="button-save" disabled>Save</button>

</fieldset>
`;
        this._enabledCheckbox = this.shadowRoot.querySelector("#checkbox-enabled");
        this._profileSelect = this.shadowRoot.querySelector('#select-profile');
        this._saveButton = this.shadowRoot.querySelector('#button-save');
        this._saveButton.onclick = () => this._save();
        this._cfg = new eventsubpb.EventSubConfig();
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
        msg.type = requestpb.MessageTypeRequest.TYPE_REQUEST_EVENT_GET_CONFIG_REQ;
        msg.message = new requestpb.EventSubGetConfigRequest().toBinary();
        bus.sendWithReply(msg, (reply: buspb.BusMessage) => {
            if (reply.error) {
                throw reply.error.detail;
            }
            this._saveButton.disabled = false;
            let gicr = requestpb.EventSubGetConfigResponse.fromBinary(reply.message);
            if (!gicr.config) {
                return;
            }
            this._cfg = gicr.config;
            this.enabled = gicr.config.enabled;
            if (gicr.config.profile) {
                this.selectedProfile = gicr.config.profile;
            }
        });
    }

    private _save() {
        this._saveButton.disabled = true;
        this._saveButton.innerText = 'Saving...';
        this._cfg.enabled = this.enabled;
        this._cfg.profile = this.selectedProfile;
        let iscr = new commandpb.EventSubSetConfigRequest();
        iscr.config = this._cfg;
        let msg = new buspb.BusMessage();
        msg.topic = TOPIC_TWITCH_COMMAND;
        msg.type = commandpb.MessageTypeCommand.TYPE_COMMAND_EVENT_SET_CONFIG_REQ;
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
customElements.define('twitc-es-config', EventSubConfig);

export { EventSubConfig };