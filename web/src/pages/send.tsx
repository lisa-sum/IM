import { useState } from 'react'

export default function Send () {
	const [user_identity, setUser_identity] = useState<string>('visitant')
	const [room_identity, setRoom_identity] = useState<string>('1')
	const [message, setMessage] = useState<string>('')
	const [messageList, setMessageList] = useState<MessageBasic[]>([
		{
			user_identity: 'visitant',
			room_identity: '1',
			data: 'visitant message',
			created_at: new Date().toLocaleString(),
			updated_at: new Date().toLocaleString(),
		},
	])
	const ms = () => {
		return JSON.stringify({
			'user_identity': user_identity,
			'room_identity': room_identity,
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
				user_identity: user_identity,
				room_identity: room_identity,
				data: event.data,
				created_at: new Date().toLocaleString(),
				updated_at: new Date().toLocaleString(),
			}
			setMessageList([...messageList, messageData])
		}
	}
	return (
		<section>
			<label htmlFor='user_identity'>your nickname:
				<input
					type='text'
					id='user_identity'
					value={ user_identity }
					onChange={ (event) => setUser_identity(event.target.value) }
				/>
			</label>
			<label htmlFor='room_identity'>your room_identity:
				<input
					type='number'
					id='room_identity'
					value={ room_identity }
					onChange={ (event) => setRoom_identity(event.target.value) }
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
	)
};
