import { Suspense, useEffect, useState } from 'react'
import { Avatar, Box, List, ListItem, Typography } from '@mui/material'
import { useAppDispatch, useAppSelector } from '@/utils/hooks/index'
import { activeRoom } from '@/features/room'
import { RootState } from '@/store/index'
import Room from '@/components/Room'

export default function Chat () {
	const dispatch = useAppDispatch()
	const { roomName } = useAppSelector((state: RootState) => state.room.value)
	const [roomList, setRoomList] = useState<RoomBasic[]>([])
	// 进入房间
	const intoRoom = (name: string, number: string) => {
		dispatch(activeRoom({ roomName: name, roomNumber: number }))
	}

	useEffect(() => {
		fetch('http://127.0.0.1:4000/roomList?number=all')
		.then(async(response) => {
			const res = await response.json()
			if (res.code === 406){
				throw new Error(res.message || '请求失败')
			}
			return res
		})
		.then((res) => {
			console.log(res)
			setRoomList(res.body)
		})
		.catch((err) => {
			console.error(err)
		})
	}, [])
	return (
		<>
			<Box
				sx={ {
					width: '1440px',
					height: '80vh',
					display: 'grid',
					bgcolor: '#ccc',
					gridTemplateRows: '60px 480px 200px',
					gridTemplateColumns: 'minmax(280px, 380px) minmax(740px, 1060px)',
					gridTemplateAreas: `"sidebar header"
				"sidebar main"
				"sidebar footer"`,
				} }>

				<Box
					component='section'
					sx={ {
						gridArea: 'sidebar',
						bgcolor: '#000',
					} }>
					<List
						sx={ {
							width: '380px',
						} }
					>
						<Suspense fallback={ '正在查询ing...' }>
							{
								roomList && roomList.map((item: RoomBasic) => (<ListItem
									sx={ {
										position: 'relative',
										width: '380px',
										height: '90px',
										py: '15px',
										':hover': {
											bgcolor: '#555',
										},
									} }
									onClick={ () => intoRoom(item.name, item.number) }
									key={ item.number }>
									<Avatar
										src={ item.avatar }
										alt={ item.name }
										sx={ {
											width: '60px',
											height: '60px',
										} } />
									<Typography
										sx={ {
											ml: '15px',
										} }
									>{ item.name }</Typography>
									<Typography sx={ {
										position: 'absolute',
										right: '20px',
									} }>{ `${ new Date(item.updated_at).getHours() }:${ new Date(item.updated_at).getMinutes() }` }</Typography>
								</ListItem>))
							}
						</Suspense>
					</List>
				</Box>

				<Box
					component='section'
					sx={ {
						gridArea: 'header',
						textAlign: 'center',
						bgcolor: '#000',
					} }>
					<Typography
						sx={ {
							fontSize: '20px',
							color: '#fff',
							lineHeight: '60px',
						} }>{ roomName }
					</Typography>
				</Box>

				{
					roomName === '' ? <></> : <Room />
				}
			</Box>
		</>
	)
};
