import { Correction } from "../hooks/use-spell-check"
import { bestCandidate, candidate, candidates, popover } from "./variations-popover.styles"

export const VariationsPopover = ({correction, setPicked}: {correction: Correction, setPicked: (variant: string) => void}) => {
	const handleClick = (variant: string) => () => {
		if(!variant) {
			return
		}
		setPicked(variant)
	}
	return (
		<div className={popover}>
			<p>Possible spelling mistake found.</p>
			<div className={candidates}>
			<div onClick={handleClick(correction?.best)} className={bestCandidate}>{correction?.best}</div>

			{correction?.variants?.map((x, idx) => {
				return (
					<div key={`var-${idx}`} onClick={handleClick(x)} className={candidate}>{x}</div>
				)
			})}
			</div>
		</div>
	)
}
