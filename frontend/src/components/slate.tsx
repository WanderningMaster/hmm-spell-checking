/* eslint-disable react-hooks/exhaustive-deps */
/* eslint-disable @typescript-eslint/no-explicit-any */
import React, { useMemo } from 'react'
import { createEditor, Descendant, NodeEntry, BaseRange, Text, Transforms, Path, Node } from 'slate'
import { Slate, Editable, withReact, ReactEditor } from 'slate-react'
import { withHistory } from 'slate-history'
import { higlightedLeaf, slateStyles } from './slate.styles'
import { Correction, SpellCheckResult } from '../hooks/use-spell-check'
import { Popover } from '@blueprintjs/core'
import { VariationsPopover } from './variations-popover'


const Leaf: React.FC<{attributes: any, children: any, leaf: any}> = ({ attributes, children, leaf }) => {
	const [open, setOpen] = React.useState(false)
	const [picked, setPicked] = React.useState<string | null>(null)
	const handleClick = () => {
		setOpen(true)
	}
	const handleClose = () => {
		setOpen(false)
	}
	React.useEffect(() => {
		if(picked !== null) {
			leaf.replaceFn(picked)
			setPicked(null)
			setOpen(false)
		}
	}, [picked])

	if(!leaf?.correction) {
		return (
			<span
				{...attributes}
				className={higlightedLeaf(leaf.highlight)}
			>
				{children}
			</span>
		)
	}
	
	return (
		<Popover position={"bottom"} onClose={handleClose} enforceFocus={false} content={
			<VariationsPopover setPicked={setPicked} correction={leaf.correction}/>
			} isOpen={open}>
			<span
				{...attributes}
				onClick={handleClick}
				className={higlightedLeaf(leaf.highlight)}
			>
				{children}
			</span>
		</Popover>
	)
}

type TextEditorProps = {
	initialText: string;
	setPlainText: (state: string) => void;
	result: SpellCheckResult | null
}
export const TextEditor: React.FC<TextEditorProps> = ({initialText, setPlainText, result}) => {
	const initialValue: Descendant[] = [
		{
			children: [{ text: initialText },],
		},
	]
	const editor = useMemo(() => withHistory(withReact(createEditor())), [])

	const replace = (path: Path, start: number, end: number) => (variant: string) => {
		Transforms.insertText(editor, variant, {
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
		const ranges: (BaseRange & {highlight: boolean, correction: Correction, replaceFn: (variant: string) => void})[] = []
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
					regex = new RegExp(word, 'ig');
				}
				catch {
					console.log("caught an error")
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
						replaceFn: replace(path, end, start)
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
