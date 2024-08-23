import * as eventsub from "./eventsub.js";
import * as irc from "./irc.js";
import * as profiles from "./profiles.js";

function start(mainContainer: HTMLElement) {
    mainContainer.appendChild(new profiles.Profiles());
    mainContainer.appendChild(new irc.IRCConfig());
    mainContainer.appendChild(new eventsub.EventSubConfig());
}

export { start };