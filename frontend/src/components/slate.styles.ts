import { css } from "@emotion/css";

export const slateStyles = css`
	width: 100%;
	padding: 20px;
	font-size: 24px;
	line-height: 24px;
`


export const highlitedText = css`
	text-decoration: underline;
	text-decoration-color: red;
	text-decoration-thickness: 2px;
`

export const higlightedLeaf = (highlighted: boolean) => css`
	${highlighted && highlitedText}
`
