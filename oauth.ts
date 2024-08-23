import { bus, enumName } from "/bus.js";
import * as buspb from "/pb/bus/bus_pb.js";
import * as commandpb from "/m/twitch/pb/command_pb.js";
import * as twitchpb from "/m/twitch/pb/twitch_pb.js";

const TOPIC_COMMAND = enumName(twitchpb.BusTopics, twitchpb.BusTopics.TWITCH_COMMAND);

let mainContainer = document.querySelector("#mainContainer") as HTMLElement;
mainContainer.innerText = "Validated response...";

function handleOAuth() {
    let params = new URLSearchParams(document.location.hash.slice(1));
    let access_token = params.get('access_token');
    if (!access_token) {
        mainContainer.innerText = "Missing access token";
        return;
    }
    if (params.get('token_type') !== 'bearer') {
        mainContainer.innerText = "Expecting bearer token";
        return;
    }
    let scopes = params.get('scope').split(' ');
    if (!scopes.length) {
        mainContainer.innerText = "No scopes returned";
        return;
    }
    let state = params.get('state');
    let stateParts = state.split(':');
    if (stateParts.length != 2) {
        mainContainer.innerText = "Invalid state";
        return;
    }
    // precheck passed
    mainContainer.innerText = "Saving profile...";

    let wpr = new commandpb.WriteProfileRequest();
    wpr.profile = new commandpb.Profile();
    wpr.profile.name = stateParts[0];
    wpr.profile.token = new commandpb.Token();
    wpr.profile.token.clientId = state;
    wpr.profile.token.access = access_token;
    wpr.profile.token.scopes = scopes;
    let msg = new buspb.BusMessage();
    msg.topic = TOPIC_COMMAND;
    msg.type = commandpb.MessageTypeCommand.TYPE_COMMAND_WRITE_PROFILE_REQ;
    msg.message = wpr.toBinary();

    bus.sendWithReply(msg, (reply) => {
        if (reply.error) {
            mainContainer.innerText = reply.error.detail;
            return;
        }
        mainContainer.innerText = "Profile saved!";
    });
}

bus.waitForTopic(TOPIC_COMMAND, 5000).then(() => handleOAuth());