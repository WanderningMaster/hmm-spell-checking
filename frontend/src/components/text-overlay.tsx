import React, { useState } from 'react';
import { overlayStyles, textAreaStyles, highlitedText } from './text-overlay.styles';

const typos = ['hrlp']
export const TextOverlay = () => {
	const [content, setContent] = useState("");

	function splitByNewlines(inputString: string) {
		const words = inputString.split(/\n+/);

		const newlineGroups = inputString.match(/\n+/g);

		const result = [];
		for (let i = 0; i < words.length; i++) {
			result.push(words[i]);
			if (newlineGroups && i < newlineGroups.length) {
				result.push(newlineGroups[i]);
			}
		}

		return result;

	}

	function renderHiglihtedWord(x: string, key: string) {
		if(typos.includes(x.toLowerCase())) {
			return(
				<React.Fragment key={key}>
					<span className={highlitedText}>{x}</span>{" "}
				</React.Fragment>
			);
		}
		return (
			<React.Fragment key={key}>
				<span>{x} </span>
			</React.Fragment>
		)
	}

	function highlight() {
		const words = content.split(" ")

		return words.map((x, idx) => {
			if(x.includes("\n")) {
				const items = splitByNewlines(x)
				return (
					<>
						{items.map((item, part) => {
							if(item.includes("\n")) {
								return item
							}
							return renderHiglihtedWord(item, `key-${part}-${idx}`)
						})}
					</>
				)
			}
			return renderHiglihtedWord(x, `key-${idx}`)
		});
	}

	return (
		<div style={{ position: 'relative', width: '100%' }}>
			<textarea
				maxLength={100}
				rows={5}
				cols={5}
				wrap={"hard"}
				className={textAreaStyles}
				value={content}
				onChange={(e) => setContent(e.target.value)}
			/>
			{<div className={overlayStyles}>
				{highlight()}
			</div>}
		</div>
	);
};
