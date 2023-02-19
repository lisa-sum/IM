import { Box, Button } from '@mui/material'
import { Send as SendIcon } from '@mui/icons-material'
import { messageBasic } from '@/utils/message'
import { useAppDispatch } from '@/utils/hooks/index'
import { updateMessage } from '@/features/message'

export default function SendMessage ({ msg }: { msg: MessageBasic }) {
	const dispatch = useAppDispatch()
	const sendMessage = (msg: MessageBasic) => {
		const ws: WebSocket = new WebSocket('ws://127.0.0.1:4000/chat')
		ws.onopen = () => {
			ws.send(messageBasic(msg))
		}
		ws.onmessage = (event: MessageEvent<any>) => {
			dispatch(updateMessage(event.data))
		}
		ws.onerror = () => {
			console.info('WebSocket连接异常!')
		}
		ws.onclose = () => {
			console.info('WebSocket连接关闭!')
		}
	}
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
				onClick={ () => sendMessage(msg) }
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
}
