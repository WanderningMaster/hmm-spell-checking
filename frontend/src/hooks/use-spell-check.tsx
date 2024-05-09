import React from "react"

export interface SpellCheckResult {
	corrections: Correction[];
	totalErrors: number;
}

export interface Correction {
	typo: string;
	valid: boolean;
	best: string;
	variants: string[] | null;
}

async function fetchSpellCheckResult(text: string, signal: AbortSignal): Promise<SpellCheckResult> {
	const query = new URLSearchParams({
		text
	})
	const response = await fetch(`http://localhost:8080/api/spell-check?${query}`, {signal})
	if(!response.ok) {
		throw new Error(response.statusText)
	}
	const result: SpellCheckResult = await response.json()

	return result
}

export const enum CheckState {
	IDLE = 'idle', 
	LOADING = 'loading',
	CHECKING_ERROR = 'checking_error',
	ERRORS_FOUND = 'errors_found',
	ERRORS_NOT_FOUND = 'errors_not_found'
}
export const useSpellCheck = (text: string): {
	result: SpellCheckResult | null,
	state: CheckState,
} => {
	const [result, setResult] = React.useState<SpellCheckResult | null>(null)
	const [state, setState] = React.useState<CheckState>(CheckState.IDLE)
	
	React.useEffect(() => {
		const abortController = new AbortController()
		const signal = abortController.signal
		if(text.length) {
			setState(CheckState.LOADING)
			fetchSpellCheckResult(text, signal)
				.then((res) => {
					setResult(res)
					if(res.totalErrors === 0) {
						setState(CheckState.ERRORS_NOT_FOUND)
					} else {
						setState(CheckState.ERRORS_FOUND)
					}
				})
				.catch((err) => {
					console.error(err)
					setState(CheckState.CHECKING_ERROR)
				})
		}

		return () => {
			setState(CheckState.IDLE)
			abortController.abort()
		}
	}, [text])

	return {
		result,
		state,
	}
}
