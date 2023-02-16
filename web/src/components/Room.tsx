import { Box, Input, InputLabel, List, ListItem, Typography } from '@mui/material'
import { useEffect, useState } from 'react'
import InputAssist from '@/components/inputAssist'
import SendMessage from '@/components/SendMessage'
import { useAppSelector } from '@/utils/hooks/index'
import { RootState } from '@/store/index'

export default function Room () {
	const [chatBarList, setChatBarList] = useState<{ text: string }[]>([
		{ text: '聊天' },
		{ text: '简介' },
	])
	const { roomNumber } = useAppSelector((state: RootState) => state.room.value)
	console.log("roomNumber",roomNumber)
	const { account, nickname } = useAppSelector((state: RootState) => state.profile.value)
	console.log("account",account)
	const [inputMsg, setInputMsg] = useState<string>('')
	const [msg, setMsg] = useState<MessageBasic>({
		userIdentity: '',
		roomIdentity: '',
		data: undefined,
		createdAt: new Date().getTime(),
		updatedAt: new Date().getTime(),
	})

	// 发送消息
	const sendMessage = (event: string) => {
		if (event === 'Enter'){
			setMsg((state: MessageBasic) => {
				return {
					...state,
					userIdentity: account,
					roomIdentity: roomNumber,
					data: inputMsg,
					updatedAt: new Date().getTime(),
				}
			})
		}
	}
	return (
		<>
			<Box
				component='section'
				sx={ {
					gridArea: 'main',
					bgcolor: '#fff',
				} }>
				<List
					sx={ {
						p: 0,
						width: 'auto',
						height: '70px',
						columnCount: 6,
						borderBottom: '1px solid #ccc',
					} }
				>
					{
						chatBarList.map((item) => (<ListItem
							key={ item.text }
							sx={ {
								height: '70px',
							} }
						>
							<Typography sx={ {
								lineHeight: '70px',
								color: '#000',
								textAlign: 'center',
								fontSize: '20px',
							} }>
								{ item.text }
							</Typography>
						</ListItem>))
					}
				</List>


			</Box>
			<Box
				component='section'
				sx={ {
					gridArea: 'footer',
					bgcolor: 'green',
					position: 'relative',
				} }>
				<InputAssist />

				<InputLabel
					htmlFor='inputMessage'
					focused={ true }
				>
					<Input
						fullWidth
						id='inputMessage'
						onKeyDown={ (event) => sendMessage(event.key) }
						onChange={ (event) => setInputMsg(event.target.value) }
					/>
				</InputLabel>
				<SendMessage msg={ msg } />
			</Box>
		</>
	)
};
