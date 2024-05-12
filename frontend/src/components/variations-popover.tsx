import { Correction } from "../hooks/use-spell-check"
import { useLocalization } from "../providers/localization.provider"
import { bestCandidate, candidate, candidates, popover } from "./variations-popover.styles"

export const VariationsPopover = ({correction, setPicked}: {correction: Correction, setPicked: (variant: string) => void}) => {
	const {text} = useLocalization()
	const handleClick = (variant: string) => () => {
		if(!variant) {
			return
		}
		setPicked(variant)
	}
	return (
		<div className={popover}>
			<p>{text.possibleMistakeFound}</p>
			<div className={candidates}>
			{correction.best && <div onClick={handleClick(correction?.best)} className={bestCandidate}>{correction?.best}</div>}

			{correction?.variants?.map((x, idx) => {
				return (
					<div key={`var-${idx}`} onClick={handleClick(x)} className={candidate}>{x}</div>
				)
			})}
			</div>
		</div>
	)
}
