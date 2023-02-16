import { configureStore } from '@reduxjs/toolkit'
import roomReducer from '@/features/room'
import profileReducer from '@/features/profile'

export const store = configureStore({
	reducer: {
		room: roomReducer,
		profile: profileReducer,
	},
})

export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch
