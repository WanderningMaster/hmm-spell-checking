export const textEN = {
	placeholder: "Type here...",
	checkingStateIdle: "Type or paste your text to check it",
	checkingStateLoading: "Checking Text...",
	checkingStateError: "Error while checking text",
	checkingStateErrorsFound: (totalErrors: number) => `${totalErrors} writing error${totalErrors > 1 ? 's' :''} found`,
	checkingStateErrorsFoundSubtitle: 'Click on the highlighted words to correct them.',
	checkingStateErrorsNotFound: "Errors not found!",
	possibleMistakeFound: "Possible spelling mistake found.",
	copied: "Copied!"
} as const
export type TextType = {
	placeholder: string
	checkingStateIdle: string
	checkingStateLoading: string
	checkingStateError: string
	checkingStateErrorsFound: (totalErrors: number) => string 
	checkingStateErrorsFoundSubtitle: string
	checkingStateErrorsNotFound: string
	possibleMistakeFound: string
	copied: string
}

export const textUA = {
	placeholder: "Друкуйте тут...",
	checkingStateIdle: "Введіть або вставте текст, щоб перевірити його",
	checkingStateLoading: "Перевірка тексту...",
	checkingStateError: "Помилка під час перевірки тексту",
	checkingStateErrorsFound: (totalErrors: number) => {
		return `Помилок знайдено: ${totalErrors}`
	},
	checkingStateErrorsFoundSubtitle: 'Натисніть на виділені слова, щоб виправити їх.',
	checkingStateErrorsNotFound: "Помилок не знайдено!",
	possibleMistakeFound: "Знайдено можливу орфографічну помилку.",
	copied: "Скопійовано!"
} as const
