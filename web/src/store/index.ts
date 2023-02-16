import { configureStore } from '@reduxjs/toolkit'
import roomReducer from '@/features/room'
import profileReducer from '@/features/profile'
import messageReducer from '@/features/message'

export const store = configureStore({
	reducer: {
		room: roomReducer,
		profile: profileReducer,
		message: messageReducer,
	},
})

export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch
