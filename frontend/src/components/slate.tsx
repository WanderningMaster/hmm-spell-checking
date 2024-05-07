/* eslint-disable @typescript-eslint/no-explicit-any */
import React, { useMemo } from 'react'
import { createEditor, Descendant, NodeEntry, BaseRange, Text, Transforms, Path, Node } from 'slate'
import { Slate, Editable, withReact } from 'slate-react'
import { withHistory } from 'slate-history'
import { higlightedLeaf, slateStyles } from './slate.styles'
import { useDebounce } from '../hooks/use-debounce'
import { Correction, useSpellCheck } from '../hooks/use-spell-check'

const initialValue: Descendant[] = [
  {
    children: [
      { text: '' },
    ],
  },
]

const Leaf: React.FC<{attributes: any, children: any, leaf: any}> = ({ attributes, children, leaf }) => {
	const handleClick = () => {
		if(leaf.replaceFn) {
			leaf.replaceFn()
		}
	}

	return (
		<span
			{...attributes}
			onClick={handleClick}
			className={higlightedLeaf(leaf.highlight)}
		>
			{children}
		</span>
	)
}

export const TextEditor = () => {
	const editor = useMemo(() => withHistory(withReact(createEditor())), [])
	const [plainText, setPlainText] = React.useState("")

	const debouncedText = useDebounce(plainText, 500)
	const { result } = useSpellCheck(debouncedText)

	const replace = (path: Path, start: number, end: number, correction: Correction) => () => {
		Transforms.insertText(editor, correction.best, {
			at: {
				anchor: { path, offset: end  },
				focus: { path, offset: start },
			}
		})
	};
	const handleChange = (value: Descendant[]) => {
		const text = value.map((x) => {
			return Node.string(x)
		}).join(" ").replace(/\s+/g, ' ').trim()
		setPlainText(text)
	}

	const decorate = React.useCallback(([node, path]: NodeEntry): BaseRange[] => {
		const ranges: (BaseRange & {highlight: boolean, correction: Correction, replaceFn: () => void})[] = []
		if(Text.isText(node) && result) {
			const { text } = node
			for(const correction of result.corrections) {
				if(correction.valid) {
					continue
				}
				const word = correction.typo
				const indexes = [];
				let match;

				const regex = new RegExp(`\\b${word}\\b`, 'g');
				while((match = regex.exec(text)) !== null) {
					indexes.push({ start: match.index, end: match.index + word.length, match: match[0] });
				}
				indexes.forEach(({start, end }) => {
					ranges.push({
						anchor: { path, offset: end  },
						focus: { path, offset: start },
						highlight: true,
						correction,
						replaceFn: replace(path, end, start, correction)
					});
				});
			}
		}

		return ranges
	}, [result])

	return (
		<Slate onChange={handleChange} editor={editor} initialValue={initialValue}>
			<Editable 
				spellCheck={false}
				decorate={decorate}
				className={slateStyles}
				placeholder="Type here..."
				renderLeaf={props => <Leaf {...props} />}
			/>
		</Slate>
	)
}
