import { Box, Button } from '@mui/material'
import { Send as SendIcon } from '@mui/icons-material'
import { messageBasic } from '@/utils/message'

const sendMessage = (msg:string) => {
	console.log("msg3:",msg)
	const ws: WebSocket = new WebSocket('ws://127.0.0.1:4000/chat')
	ws.onopen = () => {
		ws.send(msg)
	}
	ws.onmessage = (event: MessageEvent<any>) => {
		console.log("event.data",event.data)
	}
	ws.onerror = () => {
		console.info('WebSocket连接异常!')
	}
	ws.onclose = () => {
		console.info('WebSocket连接关闭!')
	}
}

export default function SendMessage ({ msg }: { msg: MessageBasic }) {

	return (
		<Box sx={ {
			position: 'absolute',
			width: '130px',
			height: '40px',
			right: '20px',
			bottom: '20px',
		} }>
			<Button
				variant='contained'
				endIcon={ <SendIcon /> }
				onClick={ () => sendMessage(messageBasic(msg)) }
				sx={ {
					color: '#fff',
					height: '40px',
					lineHeight: '40px',
				} }
			>
				发送(S)
			</Button>
		</Box>
	)
};
