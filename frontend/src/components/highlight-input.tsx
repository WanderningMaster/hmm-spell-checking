/* eslint-disable @typescript-eslint/no-explicit-any */
import React from "react"
import { canvas, inputContainer } from "./highlight-input.styles"
import ContentEditable from "react-contenteditable"

type InputHighlighterProps = {
	toHighlight: string[]
}



export const InputHighlighter: React.FC<InputHighlighterProps> = ({toHighlight}) => {
	const [text, setText] = React.useState("")
	const [modifiedText, setModifiedText] = React.useState("")
	const inputRef = React.useRef<HTMLDivElement | null>(null)

	function highlight() {
		const words = text.split("")
		let html = ""
		for(const word of words) {
			if(toHighlight.includes(word.toLowerCase())) {
				html += `<span><u>${word}</u> </span>`
			} else {
				html += `<span>${word} </span>`
			}
		}
		console.log(html)

		return html
	}

	return (
		<div className={inputContainer}>
			<ContentEditable
				innerRef={inputRef}
				className={inputContainer}
				spellCheck={false}
				inputMode="text"
				onChange={(e: any) => {
					setText(e.nativeEvent.target.textContent)
				}}
				html={highlight()}
			/>
		</div>
	)
}
