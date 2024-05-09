import { css } from "@emotion/css";

export const section = css`
	display: flex;
	column-gap: 5px;
	align-items: center;
`

export const checkingState = css`
	display: flex;
	justify-content: space-between;
	align-items: center;
	min-height: 77px;

	padding: 15px;
	border: 1px solid rgb(226 232 240);
	border-radius: 8px;

	font-size: 12px;

	svg {
		width: 30px;
		height: auto;
	}
	p {
		margin: 0;
	}
`

export const contentWrapper = css`
	font-size: 20px;
`

export const subtitle = css`
	font-size: 15px;
`
