/* eslint-disable react-hooks/exhaustive-deps */
/* eslint-disable @typescript-eslint/no-explicit-any */
import React, { useMemo } from 'react'
import { createEditor, Descendant, NodeEntry, BaseRange, Text, Transforms, Path, Node } from 'slate'
import { Slate, Editable, withReact, ReactEditor } from 'slate-react'
import { withHistory } from 'slate-history'
import { higlightedLeaf, slateStyles } from './slate.styles'
import { Correction, SpellCheckResult } from '../hooks/use-spell-check'

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

type TextEditorProps = {
	setPlainText: (state: string) => void;
	result: SpellCheckResult | null
}
export const TextEditor: React.FC<TextEditorProps> = ({setPlainText, result}) => {
	const editor = useMemo(() => withHistory(withReact(createEditor())), [])

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

	const focusEditor = () => {
		ReactEditor.focus(editor)
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
				let match;

				let regex: RegExp;
				try {
					regex = new RegExp(`\\b${word}\\b`, 'g');
				}
				catch {
					return []
				}
				while((match = regex.exec(text)) !== null) {
					const start = match.index
					const end = match.index + word.length

					ranges.push({
						anchor: { path, offset: end  },
						focus: { path, offset: start },
						highlight: true,
						correction,
						replaceFn: replace(path, end, start, correction)
					});
				}
			}
		}

		return ranges
	}, [result])

	React.useEffect(() => {
		focusEditor()
	}, [])

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
