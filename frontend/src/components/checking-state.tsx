import { checkingState, contentWrapper, section, subtitle } from "./checking-state.styles"
import ErrorIcon from '../assets/error.svg?react'
import CheckIcon from '../assets/check.svg?react'
import CheckYellowIcon from '../assets/check-yellow.svg?react'
import CheckBlueIcon from '../assets/check-blue.svg?react'
import { CheckState } from "../hooks/use-spell-check"
import { ClipboardBtn } from "./clipboard-btn"
import { TextType } from "../common/text"
import { useLocalization } from "../providers/localization.provider"

type Content = {
	title: string;
	subtitle?: string;
}
function matchText(state: CheckState, totalErrors: number, error: string | null, text: TextType): Content {
	switch(state) {
	case CheckState.IDLE:
		return {
			title: text.checkingStateIdle
		}
	case CheckState.LOADING:
		return {
			title: text.checkingStateLoading
		}
	case CheckState.CHECKING_ERROR:
		return {
			title: text.checkingStateError,
			subtitle: error ?? ''
		}
	case CheckState.ERRORS_FOUND:
		return {
			title: text.checkingStateErrorsFound(totalErrors), 
			subtitle: text.checkingStateErrorsFoundSubtitle
		}
	case CheckState.ERRORS_NOT_FOUND:
		return {
			title: text.checkingStateErrorsNotFound,
		}
	default:
		return {
			title: text.checkingStateIdle
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
	const {text} = useLocalization()
	const content = matchText(state, totalErrors, error, text)
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
