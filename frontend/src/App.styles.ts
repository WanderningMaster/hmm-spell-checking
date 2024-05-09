import { css } from "@emotion/css";
import pattern from './assets/pattern.png'

export const rootWrapper = css`
	display: flex;
	width: 100vw;
	height: 100vh;
`

export const contentWrapper = css`
	width: 100%;
	padding: 30px;
	display: flex;
	flex-direction: column;
	row-gap: 20px;
	justify-content: center;
	height: 100vh;

	flex-basis: 70%;
`

export const sidebarText = css`
	font-size: 30px;
	font-weight: bolder;
	color: rgb(8 47 73);
`
export const checkingState = css`
	display: flex;
	column-gap: 5px;

	padding: 15px;
	align-items: center;
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

export const sidebar = css`
	display: flex;
	column-gap: 10px;
	align-items: center;
	justify-content: center;
	background-color: #ffffff;
	background-image:  repeating-radial-gradient( circle at 0 0, transparent 0, #ffffff 10px ), repeating-linear-gradient( #2c88ef55, #2c88ef );
	flex-basis: 30%;

	img {
		width: 150px;
		height: auto;
	}
`
