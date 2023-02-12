import { useEffect, useState } from 'react'
import './assets/style/app.css'
import Login from './pages/login'
import Send from './pages/send'
import HistoryMessage from './pages/historyMessage'
import { Route, Routes } from 'react-router-dom'

function App () {
	const [account, setAcc] = useState<string>('test1 Summer')
	const [password, setPwd] = useState<string>('bc6dc48b743dc5d013b1abaebd2faed2')
	const [token, setToken] = useState<string>('')
	const [messageList, setMessageList] = useState<MessageBasic[]>([
		{
			user_identity: 'visitant',
			room_identity: '1',
			data: 'visitant message',
			created_at: new Date().toLocaleString(),
			updated_at: new Date().toLocaleString(),
		},
	])
	const [historyMessage, setHistoryMessage] = useState<MessageBasic[]>([
		{
			user_identity: 'string', // 用户id 一对一
			room_identity: 'string', // 群聊/房间id 一对多
			data: 'string', // 消息数据
			created_at: 'string', // 创建时间
			updated_at: 'string', // 更新时间
		},
	])
	useEffect(() => {
		// 获取消息列表
		fetch('http://127.0.0.1:4000/historyMessage', {
			method: 'GET',
			headers: { 'Content-Type': 'application/json' },
		})
		.then(res => {
			// 判断服务器返回的状态码, ok即200
			if (res.ok){
				return res.json()
			}
			throw new Error('请求失败')
		})
		.then((res) => {
			// 对Unix时间戳转换成本地时间展示
			const data = res.body.map((item: any) => {
				item.created_at = new Date(item.created_at * 1000).toLocaleString()
				item.updated_at = new Date(item.updated_at * 1000).toLocaleString()
				return item
			})
			// 更新消息列表
			setMessageList(data)
		})
		// 异常处理
		.catch(err => {
			window.alert(err)
			console.error(err)
		})
	}, [])

	return (
		<article>


			<Send />
			<HistoryMessage historyMessage={ historyMessage } />

			<section className='message_list'>
				<ol>
					{
						messageList &&
						messageList.map((item, index) => {
							return <li key={ index }>
								<span>用户ID:{ item.user_identity }</span>
								<p>房间ID:{ item.room_identity }</p>
								<p>消息:{ item.data }</p>
								<p>创建时间:{ item.created_at }</p>
								<p>更新时间:{ item.updated_at }</p>
							</li>
						})
					}
				</ol>
			</section>
		</article>
	)
}

export default App
