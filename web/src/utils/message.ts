const notUpdatedMessage = (
	userIdentity: string,
	roomIdentity: string,
	data: any,
	updatedAt?: number,
): MessageBasic => {
	return {
		userIdentity,
		roomIdentity,
		data,
		updatedAt,
		createdAt: new Date().getTime(),
	}
}

const UpdatedMessage = (
	userIdentity: string,
	roomIdentity: string,
	data: string,
): MessageBasic => {
	return {
		userIdentity,
		roomIdentity,
		data,
		updatedAt: new Date().getTime(),
		createdAt: new Date().getTime(),
	}
}

export const messageBasic = (
	msg: MessageBasic,
): string => {
	console.log("msg:",msg)
	const { userIdentity, roomIdentity, data, createdAt, updatedAt } = msg
	if (createdAt){
		return JSON.stringify(notUpdatedMessage(userIdentity, roomIdentity, data, createdAt as number))
	}
	console.log("msg2:",JSON.stringify(UpdatedMessage(userIdentity, roomIdentity, data)))
	return JSON.stringify(UpdatedMessage(userIdentity, roomIdentity, data))
}
