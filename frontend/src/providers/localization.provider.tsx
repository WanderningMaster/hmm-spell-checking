/* eslint-disable @typescript-eslint/no-explicit-any */
import React from "react";
import { useStorage } from "../hooks/use-storage";
import { TextType, textEN, textUA } from "../common/text";

export interface LangContextType {
	text: TextType;
	lang: string
}

export const LangContext = React.createContext<LangContextType>({} as LangContextType);

export function useLocalization() {
  return React.useContext(LangContext as React.Context<LangContextType>);
}

type LocalizationProviderProps = {
  children: React.ReactNode;
};
export const LocalizationProvider: React.FC<LocalizationProviderProps> = ({ children }) => {
	const {getItem, setItem } = useStorage()
	const [lang, setLang] = React.useState<string>("en")
	const text = React.useMemo(() => {
		if(lang === "en") {
			return textEN
		} 	
		return textUA
	}, [lang]) 
	const handleChangeLang = (changes: any) => {
		if(changes?.lang) {
			setLang(changes.lang.newValue)
		}
	}

	React.useEffect(() => {
		if(import.meta.env.MODE !== "development") {
			chrome.storage.onChanged.addListener(handleChangeLang)
		}
		getItem("lang")
			.then((item) => {
				if(item) {
					setLang(item)
				} else {
					setItem("lang", lang)	
				}
			})
	}, [])

	const memoedData = React.useMemo(
		() => ({
			lang,
			text
		}),
	[lang, text]);

	return <LangContext.Provider value={memoedData}>{children}</LangContext.Provider>;
}
