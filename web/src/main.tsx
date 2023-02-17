import React from 'react'
import ReactDOM from 'react-dom/client'

import { store } from '@/store/index'
import { Provider } from 'react-redux'

import App from './App'

ReactDOM.createRoot(document.getElementById('root') as HTMLElement)
.render(
	<Provider store={ store }>
		<App/>
	</Provider>
)
