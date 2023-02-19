import React from 'react'
import './assets/style/app.css'
import { BrowserRouter, Link as RouterLink, Route, Routes } from 'react-router-dom'

import Register from '@/pages/register'
import Login from '@/pages/login'
import Chat from '@/pages/chat'
import HistoryMessage from '@/pages/historyMessage'
import { Link, List, ListItem } from '@mui/material'

const routes = [
	{ label: '注册', path: '/register', element: <Register /> },
	{ label: '登录', path: '/login', element: <Login /> },
	{ label: 'IM', path: '/chat', element: <Chat /> },
	{ label: '消息', path: '/historyMessage', element: <HistoryMessage /> },
]

function App () {
	return (
		<BrowserRouter>
			<List
				sx={ {
					width: '60vw',
					display: 'flex',
					justifyContent: 'space-between',
				} }
			>
				{
					routes.map((route) => (
						<Link
							key={ route.label }
							to={ route.path }
							component={ RouterLink }
							sx={ {
								display: 'block',
								textDecoration: 'none',
							} }
						>
							<ListItem
								sx={ {
									width: '15vw',
									height: '49px',
									lineHeight: '49px',
									textAlign: 'center',
								} }
							>
								{ route.label }
							</ListItem>
						</Link>
					))
				}
			</List>

			{/* 路由组 */}
			<Routes>
				{
					routes.map((route) => (
						<Route key={ route.label } path={ route.path } element={ route.element }>{ route.label }</Route>
					))
				}
			</Routes>
		</BrowserRouter>
	)
}

export default App
