import { useState } from 'react'

export default function Send () {
	const [userIdentity, setuserIdentity] = useState<string>('student')
	const [roomIdentity, setroomIdentity] = useState<string>('1')
	const [message, setMessage] = useState<string>('')
	const [messageList, setMessageList] = useState<MessageBasic[]>([
		{
			userIdentity: 'student',
			roomIdentity: '1',
			data: 'student message',
			createdAt: new Date().toLocaleString(),
			updatedAt: new Date().toLocaleString(),
		},
	])
	const ms = () => {
		return JSON.stringify({
			'userIdentity': userIdentity,
			'roomIdentity': roomIdentity,
			'message': message,
			'create_at': Math.round(new Date().getTime() / 1000),
			'update_at': Math.round(new Date().getTime() / 1000),
		})
	}
	const send = () => {
		const ws = new WebSocket('ws://127.0.0.1:4000/')
		ws.onopen = () => {
			ws.send(ms())
		}
		ws.onmessage = (event) => {
			const messageData = {
				userIdentity: userIdentity,
				roomIdentity: roomIdentity,
				data: event.data,
				createdAt: new Date().toLocaleString(),
				updatedAt: new Date().toLocaleString(),
			}
			setMessageList([...messageList, messageData])
		}
	}
	return (
		<section>
			<label htmlFor='userIdentity'>your nickname:
				<input
					type='text'
					id='userIdentity'
					value={ userIdentity }
					onChange={ (event) => setuserIdentity(event.target.value) }
				/>
			</label>
			<label htmlFor='roomIdentity'>your roomIdentity:
				<input
					type='number'
					id='roomIdentity'
					value={ roomIdentity }
					onChange={ (event) => setroomIdentity(event.target.value) }
				/>
			</label>
			<label htmlFor='message'>message:
				<input
					type='search'
					id='message'
					value={ message }
					onChange={ (event) => setMessage(event.target.value) }
				/>
			</label>
			<button type='button' id='send' onClick={ send }>send</button>

			<ol>
				{
					messageList && messageList.map((item, index) => {
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
	)
};
