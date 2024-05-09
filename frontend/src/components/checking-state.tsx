import { checkingState, contentWrapper, section, subtitle } from "./checking-state.styles"
import ErrorIcon from '../assets/error.svg?react'
import CheckIcon from '../assets/check.svg?react'
import CheckBlueIcon from '../assets/check-blue.svg?react'
import { CheckState } from "../hooks/use-spell-check"

type Content = {
	title: string;
	subtitle?: string;
}
function matchText(state: CheckState, totalErrors: number): Content {
	switch(state) {
	case CheckState.IDLE:
		return {
			title: "Type or paste your text to check it"
		}
	case CheckState.LOADING:
		return {
			title: "Checking Text..."
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
		return <CheckIcon/>
	case CheckState.ERRORS_FOUND:
		return <ErrorIcon/>
	case CheckState.ERRORS_NOT_FOUND:
		return <CheckIcon/>
	default:
		return <CheckBlueIcon/>
	}
}

export const CheckingStateCompnent = ({state, totalErrors}: {
	state: CheckState,
	totalErrors: number
}) => {
	const content = matchText(state, totalErrors)
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
			<div>
			</div>
		</div>
	)
}
