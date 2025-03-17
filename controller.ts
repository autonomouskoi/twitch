import { bus, enumName } from '/bus.js';
import * as buspb from '/pb/bus/bus_pb.js';
import * as twitchpb from '/m/twitch/pb.js';
import { ValueUpdater } from '/vu.js';

const TOPIC_REQUEST = enumName(twitchpb.BusTopics, twitchpb.BusTopics.TWITCH_REQUEST);
const TOPIC_COMMAND = enumName(twitchpb.BusTopics, twitchpb.BusTopics.TWITCH_COMMAND);

class ChatCfg extends ValueUpdater<twitchpb.chat.ChatConfig> {
    constructor() {
        super(new twitchpb.chat.ChatConfig());
    }

    refresh() {
        bus.waitForTopic(TOPIC_REQUEST, 5000)
            .then(() => {
                return bus.sendAnd(new buspb.BusMessage({
                    topic: TOPIC_REQUEST,
                    type: twitchpb.requests.MessageTypeRequest.TYPE_REQUEST_CHAT_GET_CONFIG_REQ,
                    message: new twitchpb.requests.ChatGetConfigRequest().toBinary(),
                }));
            }).then((reply) => {
                let resp = twitchpb.requests.ChatGetConfigResponse.fromBinary(reply.message);
                this.update(resp.config);
            });
    }

    save(cfg: twitchpb.chat.ChatConfig): Promise<void> {
        let cscr = new twitchpb.command.ChatSetConfigRequest({
            config: cfg,
        });
        return bus.sendAnd(new buspb.BusMessage({
            topic: TOPIC_COMMAND,
            type: twitchpb.command.MessageTypeCommand.TYPE_COMMAND_CHAT_SET_CONFIG_REQ,
            message: new twitchpb.command.ChatSetConfigRequest({
                config: cfg,
            }).toBinary(),
        })).then((reply) => {
            let resp = twitchpb.command.ChatSetConfigResponse.fromBinary(reply.message);
            this.update(resp.config);
        });
    }
}

class EventCfg extends ValueUpdater<twitchpb.event.EventSubConfig> {
    constructor() {
        super(new twitchpb.event.EventSubConfig());
    }

    refresh() {
        bus.waitForTopic(TOPIC_REQUEST, 5000)
            .then(() => {
                return bus.sendAnd(new buspb.BusMessage({
                    topic: TOPIC_REQUEST,
                    type: twitchpb.requests.MessageTypeRequest.TYPE_REQUEST_EVENT_GET_CONFIG_REQ,
                    message: new twitchpb.requests.EventSubGetConfigRequest().toBinary(),
                }));
            }).then((reply) => {
                let resp = twitchpb.requests.EventSubGetConfigResponse.fromBinary(reply.message);
                this.update(resp.config);
            });
    }

    save(cfg: twitchpb.event.EventSubConfig): Promise<void> {
        let cscr = new twitchpb.command.EventSubSetConfigRequest({
            config: cfg,
        });
        return bus.sendAnd(new buspb.BusMessage({
            topic: TOPIC_COMMAND,
            type: twitchpb.command.MessageTypeCommand.TYPE_COMMAND_EVENT_SET_CONFIG_REQ,
            message: new twitchpb.command.EventSubSetConfigRequest({
                config: cfg,
            }).toBinary(),
        })).then((reply) => {
            let resp = twitchpb.command.EventSubSetConfigResponse.fromBinary(reply.message);
            this.update(resp.config);
        });
    }
}
export { ChatCfg, EventCfg };