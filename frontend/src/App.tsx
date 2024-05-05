import {contentWrapper, rootWrapper} from './App.styles'
import {TextOverlay} from './components/text-overlay'

function App() {
	return (
		<div className={rootWrapper}>
			<div className={contentWrapper}>
				<TextOverlay/>
			</div>
		</div>
	)
}

export default App
