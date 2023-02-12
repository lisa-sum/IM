import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import Login from './pages/login'
import Register from './pages/register'

ReactDOM.createRoot(document.getElementById('root') as HTMLElement)
.render(
	<BrowserRouter>
		<Routes>
			<Route path='/login' element={ <Login /> }>Login</Route>
			<Route path='/register' element={ <Register /> }>register</Route>
			<Route path='/' element={ <App /> }>index</Route>
		</Routes>
	</BrowserRouter>,
)
