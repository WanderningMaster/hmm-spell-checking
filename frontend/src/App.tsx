import {contentWrapper, rootWrapper, sidebar, sidebarText} from './App.styles'
import Logo from './assets/icon-without-bg.png'
import { CheckingStateCompnent } from './components/checking-state'
import { TextEditor } from './components/slate'
import { useDebounce } from './hooks/use-debounce'
import { useSpellCheck } from './hooks/use-spell-check'
import React from 'react'

function App() {
	const [plainText, setPlainText] = React.useState("")

	const debouncedText = useDebounce(plainText, 500)
	const { result, state } = useSpellCheck(debouncedText)
	return (
		<div className={rootWrapper}>
			<div className={sidebar} >
				<img src={Logo} alt="logo" />		
				<div className={sidebarText} >Taipo</div>
			</div>
			<div className={contentWrapper}>
				<CheckingStateCompnent state={state} totalErrors={result?.totalErrors ?? 0} />
				<TextEditor setPlainText={setPlainText} result={result} />
			</div>
		</div>
	)
}

export default App
