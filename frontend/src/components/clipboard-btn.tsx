import React from "react";
import { Button } from "@blueprintjs/core";
import Clipboard from "../assets/clipboard.svg?react"
import { btn } from "./clipboard-btn.styles";
import { useLocalization } from "../providers/localization.provider";

type ClipboardBtnProps = {
	text: string;
	disabled?: boolean
}
export const ClipboardBtn: React.FC<ClipboardBtnProps> = ({text: textToCopy, disabled}) => {
	const {text} = useLocalization()
	const [clicked, setClicked] = React.useState(false)
	const handleClick = () => {
		navigator.clipboard.writeText(textToCopy)
		setClicked(true)
		const timer = setTimeout(() => {
			setClicked(false)
			clearTimeout(timer)
		}, 1000)
	}

	const BtnContent = React.useMemo(() => {
		if(clicked) {
			return <p>{text.copied}</p>
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
