/* eslint-disable @typescript-eslint/no-explicit-any */
import {contentWrapper, rootWrapper, sidebar, sidebarText} from './App.styles'
import Logo from './assets/icon-without-bg.png'
import { CheckingStateCompnent } from './components/checking-state'
import { TextEditor } from './components/slate'
import { useDebounce } from './hooks/use-debounce'
import { useSpellCheck } from './hooks/use-spell-check'
import React from 'react'
import { useStorage } from './hooks/use-storage'

function App() {
	const {getItem, setItem} = useStorage()
	const [plainText, setPlainText] = React.useState("")
	const [initial, setInitial] = React.useState<string | null>(null)

	const debouncedText = useDebounce(plainText, 500)
	const { result, state, error } = useSpellCheck(debouncedText)

	const setText = (value: string) => {
		setPlainText(value)
		setItem("text", value)
	}

	async function restoreSession() {
		const text = await getItem("text")
		if(!text) {
			setItem("text", "")
			setInitial("")
		} else {
			setInitial(text)
		}
	}

	React.useEffect(() => {
		restoreSession()
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
