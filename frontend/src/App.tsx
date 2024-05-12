/* eslint-disable @typescript-eslint/no-explicit-any */
import {contentWrapper, rootWrapper, sidebar, sidebarText} from './App.styles'
import Logo from './assets/icon-without-bg.png'
import { CheckingStateCompnent } from './components/checking-state'
import { TextEditor } from './components/slate'
import { useDebounce } from './hooks/use-debounce'
import { useSpellCheck } from './hooks/use-spell-check'
import React from 'react'

function App() {
	import.meta.env.MODE
	const [plainText, setPlainText] = React.useState("")
	const [initial, setInitial] = React.useState<string | null>(null)

	const debouncedText = useDebounce(plainText, 500)
	const { result, state, error } = useSpellCheck(debouncedText)

	const setText = (value: string) => {
		setPlainText(value)
		if(import.meta.env.MODE === "development") {
			localStorage.setItem("text", value)
		} else {
			chrome.storage.local.set({text: value})
		}
	}

	function restoreSessionProd() {
		chrome.storage.local.get()
			.then((x: any) => {
				if(!x.text) {
					chrome.storage.local.set({text: ""})
					setInitial("")	
				}
				else {
					setInitial(x.text)	
				}
			})
	}
	function restoreSession() {
		const text = localStorage.getItem("text")
		if(!text) {
			localStorage.setItem("text", "")
			setInitial("")
		} else {
			setInitial(text)
		}
	}

	React.useEffect(() => {
		if(import.meta.env.MODE === "development") {
			restoreSession()
		} else {
			restoreSessionProd()
		}
	}, [])

	if(initial === null) {
		return null
	}

	return (
		<div className={rootWrapper}>
			<div className={sidebar} >
				<img src={Logo} alt="logo" />		
				<div className={sidebarText} >Taipo</div>
			</div>
			<div className={contentWrapper}>
				<CheckingStateCompnent plainText={plainText} error={error} state={state} totalErrors={result?.totalErrors ?? 0} />
				<TextEditor initialText={initial ?? ''} setPlainText={setText} result={result} />
			</div>
		</div>
	)
}

export default App
