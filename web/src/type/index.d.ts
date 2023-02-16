type RoomBasic = {
	user_identity: string
	avatar: string
	info: string
	name: string
	number: string
	created_at: number
	updated_at: number
}

// WebSocket发送消息的格式
type MessageBasic = {
	userIdentity: string
	roomIdentity: string
	data: any
	createdAt: number | string
	updatedAt?: number | string
}

type UserProfileInfo = {
	account: string
	nickname: string
	password?: string
	avatar: string

	email: string
	createdAt: string
}
