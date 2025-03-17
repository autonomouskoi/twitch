import * as eventsub from './eventsub.js';
import * as chat from './chat.js';
import * as profiles from './profiles.js';
import * as ctrl from './controller.js';

function start(mainContainer: HTMLElement) {
    let chatCfg = new ctrl.ChatCfg();
    let eventCfg = new ctrl.EventCfg();

    mainContainer.classList.add('flex-column');
    mainContainer.style.setProperty('gap', '1rem');

    mainContainer.appendChild(new profiles.Profiles());
    mainContainer.appendChild(new chat.ChatConfig(chatCfg));
    mainContainer.appendChild(new eventsub.EventSubConfig(eventCfg));

    chatCfg.refresh();
    eventCfg.refresh();
}

export { start };