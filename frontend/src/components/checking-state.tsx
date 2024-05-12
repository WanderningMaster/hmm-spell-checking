import { checkingState, contentWrapper, section, subtitle } from "./checking-state.styles"
import ErrorIcon from '../assets/error.svg?react'
import CheckIcon from '../assets/check.svg?react'
import CheckYellowIcon from '../assets/check-yellow.svg?react'
import CheckBlueIcon from '../assets/check-blue.svg?react'
import { CheckState } from "../hooks/use-spell-check"
import { ClipboardBtn } from "./clipboard-btn"

type Content = {
	title: string;
	subtitle?: string;
}
function matchText(state: CheckState, totalErrors: number, error: string | null): Content {
	switch(state) {
	case CheckState.IDLE:
		return {
			title: "Type or paste your text to check it"
		}
	case CheckState.LOADING:
		return {
			title: "Checking Text..."
		}
	case CheckState.CHECKING_ERROR:
		return {
			title: "Error while checking text.",
			subtitle: error ?? ''
		}
	case CheckState.ERRORS_FOUND:
		return {
			title: `${totalErrors} writing error${totalErrors > 1 ? 's' :''} found`, 
			subtitle: 'Click on the highlighted words to correct them.'
		}
	case CheckState.ERRORS_NOT_FOUND:
		return {
			title: `Errors not found!`, 
		}
	default:
		return {
			title: "Type or paste your text to check it"
		}
	}
}

function matchIcon(state: CheckState) {
	switch(state) {
	case CheckState.IDLE:
		return <CheckBlueIcon/>
	case CheckState.LOADING:
		return <CheckYellowIcon/>
	case CheckState.CHECKING_ERROR:
		return <ErrorIcon/>
	case CheckState.ERRORS_FOUND:
		return <ErrorIcon/>
	case CheckState.ERRORS_NOT_FOUND:
		return <CheckIcon/>
	default:
		return <CheckBlueIcon/>
	}
}

export const CheckingStateCompnent = ({state, totalErrors, plainText, error}: {
	state: CheckState,
	error: string | null,
	totalErrors: number
	plainText: string
}) => {
	const content = matchText(state, totalErrors, error)
	const icon = matchIcon(state)	
	return (
		<div className={checkingState}>
			<div className={section}>
				{icon}	
				<div className={contentWrapper}>
					<p>{content.title}</p>
					{content.subtitle && <p className={subtitle}>{content.subtitle}</p>}
				</div>
			</div>
			<div className={section}>
				<ClipboardBtn disabled={totalErrors !== 0 || plainText.length === 0} text={plainText}/>
			</div>
		</div>
	)
}
