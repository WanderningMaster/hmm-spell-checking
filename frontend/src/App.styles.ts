import { css } from "@emotion/css";

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
