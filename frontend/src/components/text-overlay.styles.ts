import { css } from "@emotion/css"

export const textAreaStyles = css`
	  width: 100%;
          position: absolute;
	  padding: 0;
	  border: 0;
	  height: 100%;
          top: 0;
          left: 0;
	  font-size: 24px;
	  line-height: 24px;
	  z-index: 1;
	  resize: none;

	  &:focus {
		  outline: none;
	  }
`

export const highlitedText = css`
	text-decoration: underline;
	text-decoration-color: red;
	text-decoration-thickness: 2px;
`

export const overlayStyles = css`
	  width: 100%;
          position: absolute;
	  padding: 0;
	  border: 0;
	  outline: 0;
	  height: 100%;
          top: 0;
          left: 0;
	  font-size: 24px;
	  line-height: 24px;
	  z-index: 2;

          pointer-events: none;
          white-space: pre-wrap;
`
