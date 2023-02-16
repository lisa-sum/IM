import type { PayloadAction } from '@reduxjs/toolkit'
import { createSlice } from '@reduxjs/toolkit'

const initialState: { value: { messageList: MessageBasic[] } } = {
	value: {
		messageList: [
			{
				userIdentity: 'userIdentity',
				roomIdentity: 'roomIdentity',
				data: 'data',
				createdAt: new Date().toLocaleString(),
				updatedAt: new Date().toLocaleString(),
			},
		],
	},
}

export const profileSlice = createSlice({
	name: 'message',
	initialState,
	reducers: {
		updateMessage: (state, { payload }: PayloadAction<any>) => {
			console.log('payload', payload)
//			state.value.messageList.push(payload)
			state.value.messageList = [...state.value.messageList, JSON.parse(payload)]
		},
	},
})

// Action creators are generated for each case reducer function
export const { updateMessage } = profileSlice.actions

export default profileSlice.reducer
