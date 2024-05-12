const isDev = import.meta.env.MODE === "development"

type Key = 'text' | 'lang'
export const useStorage = () => {	
	async function getItem(key: Key): Promise<string | null> {
		if(isDev) {
			return localStorage.getItem(key)
		} else {
			const store = await chrome.storage.local.get()
			return store[key]
		}
	}
	async function setItem(key: Key, value: string) {
		if(isDev) {
			localStorage.setItem(key, value)
		} else {
			await chrome.storage.local.set({[key]: value})
		}
	}


	return {
		setItem,
		getItem
	}
}
