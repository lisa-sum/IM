import { Avatar, Box, Input, InputLabel, List, ListItem, Typography } from '@mui/material'
import { useState } from 'react'
import InputAssist from '@/components/inputAssist'
import SendMessage from '@/components/SendMessage'
import { useAppSelector } from '@/utils/hooks/index'
import { RootState } from '@/store/index'

export default function Room () {
	const messageList = useAppSelector((state: RootState) => state.message.value.messageList)
	const [chatBarList, setChatBarList] = useState<{ text: string }[]>([
		{ text: '聊天' },
		{ text: '简介' },
	])

	console.log('messageList', messageList)
	const { roomNumber } = useAppSelector((state: RootState) => state.room.value)
	const { account, nickname, avatar } = useAppSelector((state: RootState) => state.profile.value)
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

				<List
					sx={ {
						overflowY: 'auto',
						minHeight: '410px',
						mt: '0 40px',
					} }
				>
					{
						messageList.map((message: MessageBasic, index: number) => (
							<ListItem
								key={ index }
								sx={ {
									width: 'auto',
									height: '85px',
									mb: '40px',
									position: 'relative',
								} }
							>
								<Typography
									sx={ {
										position: 'absolute',
										left: '85px',
										top: '10px',
										color: '#ccc',
										fontSize: 20,
									} }
								>{ message.userIdentity }</Typography>
								<Avatar
									src={ avatar }
									alt={ message.userIdentity }
									sx={ {
										position: 'absolute',
										left: '20px',
										top: '20px',
										width: '50px',
										height: '50px',
										borderRadius: '50%',
									} }
								/>
								<Box
									sx={ {
										position: 'absolute',
										top: '40px',
										left: '85px',
										bgcolor: '#e5e5e5',
										borderRadius: '12px',
									} }
								>
									<Typography
										sx={ {
											m: '15px 15px',
											color: '#000',
											fontSize: 18,
										} }
									>{ message.data }</Typography>
								</Box>
							</ListItem>
						))
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
