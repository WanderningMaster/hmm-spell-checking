import { css } from "@emotion/css";

export const popover = css`
	width: 350px;
	padding: 20px;
`

export const candidates = css`
	display: flex;
	row-gap: 10px;
	column-gap: 5px;
	flex-wrap: wrap;
`

export const candidate = css`
	cursor: pointer;
	text-align: center;
	min-width: 50px;
	color: white;
	background: rgb(37 99 235);
	padding: 8px;
	border-radius: 5px;
	transition: all .5s ease;

	&:hover {
		background: rgb(30 58 138);
	}
`

export const bestCandidate = css`
	cursor: pointer;
	text-align: center;
	min-width: 50px;
	color: white;
	background: rgb(34 197 94);
	padding: 8px;
	border-radius: 5px;
	transition: all .5s ease;

	&:hover {
		background: rgb(22 101 52);
	}
`
