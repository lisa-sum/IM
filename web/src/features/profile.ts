import type { PayloadAction } from '@reduxjs/toolkit'
import { createSlice } from '@reduxjs/toolkit'

const initialState: { value: UserProfileInfo } = {
	value: {
		account: '',
		nickname: '',
		password: '',
		avatar: '',
		email: '',
		createdAt: '',
	},
}

export const profileSlice = createSlice({
	name: 'profile',
	initialState,
	reducers: {
		updateUserProfile: (state, { payload }: PayloadAction<UserProfileInfo>) => {
			console.log({ ...payload })
			state.value = { ...payload }
		},
	},
})

// Action creators are generated for each case reducer function
export const { updateUserProfile } = profileSlice.actions

export default profileSlice.reducer
