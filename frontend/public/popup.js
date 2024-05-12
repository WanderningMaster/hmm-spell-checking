const textEN = {
	openEditor: 'Open Editor',
	interfaceLang: 'Interface Language',
	ua: 'Ukrainian',
	en: 'English'
}
const textUA = {
	openEditor: "Відкрити редактор",
	interfaceLang: "Мова інтерфейсу",
	ua: "Українська",
	en: "Англійська"
}

function setEnglish() {
	const openTab = document.getElementById("openTab")
	const title = document.getElementById("title")
	const ua = document.getElementById("ua")
	const en = document.getElementById("en")

	openTab.innerText = textEN.openEditor
	title.innerText = textEN.interfaceLang
	ua.innerText = textEN.ua
	en.innerText = textEN.en
}

function setUkrainian() {
	const openTab = document.getElementById("openTab")
	const title = document.getElementById("title")
	const ua = document.getElementById("ua")
	const en = document.getElementById("en")

	openTab.innerText = textUA.openEditor
	title.innerText = textUA.interfaceLang
	ua.innerText = textUA.ua
	en.innerText = textUA.en
}

async function setActiveLangDefault() {
	let lang = (await chrome.storage.local.get())?.lang
	if(!lang) {
		lang = "en"
		await chrome.storage.local.set({lang: "en"})
	}
	// const lang = localStorage.getItem("lang")
	let btn;
	if(lang === "en") {
		btn = document.getElementById("en")
		setEnglish()
	} else {
		btn = document.getElementById("ua")
		setUkrainian()
	}
	btn.classList.add("active")
}

async function setActiveLang(event) {
	let ua = document.getElementById("ua")
	let en = document.getElementById("en")
	let lang = event.target.id	

	if(lang === "ua") {
		// localStorage.setItem("lang", "ua")
		await chrome.storage.local.set({lang: "ua"})
		ua.classList.add("active")
		en.classList.remove("active")
		setUkrainian()
	} else {
		// localStorage.setItem("lang", "en")
		await chrome.storage.local.set({lang: "en"})
		en.classList.add("active")
		ua.classList.remove("active")
		setEnglish()
	}
}

document.getElementById('openTab').addEventListener('click', function() {
  chrome.tabs.create({url: 'index.html'});
});

setActiveLangDefault()

const btns = document.getElementsByClassName("lang")
for(const btn of btns) {
	btn.addEventListener('click', setActiveLang);
}
