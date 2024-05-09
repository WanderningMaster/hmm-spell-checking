import React from "react";
import { Button } from "@blueprintjs/core";
import Clipboard from "../assets/clipboard.svg?react"
import { btn } from "./clipboard-btn.styles";

type ClipboardBtnProps = {
	text: string;
	disabled?: boolean
}
export const ClipboardBtn: React.FC<ClipboardBtnProps> = ({text, disabled}) => {
	const [clicked, setClicked] = React.useState(false)
	const handleClick = () => {
		navigator.clipboard.writeText(text)
		setClicked(true)
		const timer = setTimeout(() => {
			setClicked(false)
			clearTimeout(timer)
		}, 1000)
	}

	const BtnContent = React.useMemo(() => {
		if(clicked) {
			return <p>Copied!</p>
		}
		return <Clipboard/>
	}, [clicked])

	return <Button
		disabled={disabled}
		onClick={handleClick}
		className={btn}
		icon={BtnContent}
	/>
}
