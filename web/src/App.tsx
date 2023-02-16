import { useEffect, useState } from 'react'
import './assets/style/app.css'
import Send from './pages/send'
import HistoryMessage from './pages/historyMessage'

function App () {
	const [messageList, setMessageList] = useState<MessageBasic[]>([
		{
			userIdentity: 'student',
			roomIdentity: '2',
			data: 'student test text message',
			createdAt: new Date().toLocaleString() as string,
			updatedAt: new Date().toLocaleString() as string,
		},
	])
	const [historyMessage, setHistoryMessage] = useState<MessageBasic[]>([
		{
			userIdentity: 'string', // 用户id 一对一
			roomIdentity: 'string', // 群聊/房间id 一对多
			data: 'string', // 消息数据
			createdAt:  new Date().getTime(), // 创建时间
			updatedAt:  new Date().getTime(), // 更新时间
		},
	])
	useEffect(() => {
		// 获取消息列表
		fetch('http://127.0.0.1:4000/historyMessage', {
			method: 'GET',
			headers: { 'Content-Type': 'application/json' },
		})
		.then(async (res) => {
			const result = await res.json()
			// 判断服务器返回的状态码, ok即200
			if (res.ok){

				return result
			}
			throw new Error('请求失败')
		})
		.then((res) => {
			// 对Unix时间戳转换成本地时间展示
			const data = res.body.map((item: any) => {
				item.createdAt = new Date(item.createdAt).toLocaleString()
				item.updatedAt = new Date(item.updatedAt).toLocaleString()
				return item
			})
			console.log(JSON.parse(data))
			// 更新消息列表
//			console.log(JSON.parse(res.body))
//			setMessageList(JSON.parse(res.body))
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
								<span>用户ID:{ item.userIdentity }</span>
								<p>房间ID:{ item.roomIdentity }</p>
								<p>消息:{ item.data }</p>
								<p>创建时间:{ item.createdAt }</p>
								<p>更新时间:{ item.updatedAt }</p>
							</li>
						})
					}
				</ol>
			</section>
		</article>

	)
}

export default App
