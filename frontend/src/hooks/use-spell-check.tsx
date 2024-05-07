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

async function fetchSpellCheckResult(text: string): Promise<SpellCheckResult> {
	const query = new URLSearchParams({
		text
	})
	const response = await fetch(`http://localhost:8080/api/spell-check?${query}`)
	if(!response.ok) {
		throw new Error(response.statusText)
	}
	const result: SpellCheckResult = await response.json()

	return result
}

export const useSpellCheck = (text: string) => {
	const [result, setResult] = React.useState<SpellCheckResult | null>(null)
	
	React.useEffect(() => {
		if(text.length) {
			fetchSpellCheckResult(text)
				.then(setResult)
				.catch((err) => console.error(err))
		}
	}, [text])

	return {
		result
	}
}
