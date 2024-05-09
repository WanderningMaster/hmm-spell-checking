import { css } from "@emotion/css";

export const slateStyles = css`
	width: 100%;
	font-size: 20px;
	line-height: 20px;
	height: 100%;
	overflow-y: auto;
	overflow-x: none;

	&::-webkit-scrollbar {
		display: none;
	}

	&:focus {
		outline: none;
	}
`


export const highlitedText = css`
	text-decoration: underline;
	text-decoration-color: red;
	text-decoration-thickness: 2px;
`

export const higlightedLeaf = (highlighted: boolean) => css`
	${highlighted && highlitedText}
`
