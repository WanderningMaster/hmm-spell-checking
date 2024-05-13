import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import './index.css'
import { LocalizationProvider } from './providers/localization.provider'

ReactDOM.createRoot(document.getElementById('root')!).render(
	<React.StrictMode>
		<LocalizationProvider> 
			<App />
		</LocalizationProvider>
	</React.StrictMode>,
)
