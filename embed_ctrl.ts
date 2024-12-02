import * as eventsub from "./eventsub.js";
import * as chat from "./chat.js";
import * as profiles from "./profiles.js";

function start(mainContainer: HTMLElement) {
    mainContainer.appendChild(new profiles.Profiles());
    mainContainer.appendChild(new eventsub.EventSubConfig());
    mainContainer.appendChild(new chat.ChatConfig());
}

export { start };