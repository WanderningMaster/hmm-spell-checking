/* eslint-disable no-undef */
const CONTEXT_MENU_ID = "MY_CONTEXT_MENU";
async function getword(info,) {
	if (info.menuItemId !== CONTEXT_MENU_ID) {
		return;
	}
	await chrome.storage.local.set({text: info.selectionText})

	chrome.tabs.create({url: 'index.html'});
}
chrome.contextMenus.create({
  title: "Check spelling of selected text...", 
  contexts:["selection"], 
  id: CONTEXT_MENU_ID
});
chrome.contextMenus.onClicked.addListener(getword)
