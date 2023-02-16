import React from 'react'
import ReactDOM from 'react-dom/client'
import { BrowserRouter, Route, Routes } from 'react-router-dom'

import { store } from '@/store/index'
import { Provider } from 'react-redux'

import App from './App'
import Login from './pages/login'
import Chat from './pages/chat'
import Register from './pages/register'

ReactDOM.createRoot(document.getElementById('root') as HTMLElement)
.render(
	<Provider store={ store }>
		<BrowserRouter>
			<Routes>
				<Route path='/login' element={ <Login /> }>Login</Route>
				<Route path='/register' element={ <Register /> }>register</Route>
				<Route path='/chat' element={ <Chat /> }>register</Route>
				<Route path='/' element={ <App /> }>index</Route>
			</Routes>
		</BrowserRouter>
	</Provider>,
)
