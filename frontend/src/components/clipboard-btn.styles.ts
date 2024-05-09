import { css } from "@emotion/css";

export const btn = css`
	background: white !important;
	border: 1px solid black;
	border-radius: 10px !important;
	outline: none !important;

	svg {
		width: 20px;
		height: auto;
	}

	&:active {
		&:disabled {
			background: white !important;
		}
		background: rgb(203 213 225) !important;
	}

	&:disabled {
		box-shadow: inset 0 0 0 1px rgba(17, 20, 24, 0.2), 0 1px 2px rgba(17, 20, 24, 0.1) !important;
	}
`
