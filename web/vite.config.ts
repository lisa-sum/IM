import { defineConfig } from 'vite'
import { resolve } from 'node:path'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
	plugins: [react()],
	resolve: {
		alias: {
			// 别名配置
			'@': resolve(__dirname, 'src'),
		},
	},
	server: {
		port: 3000,
		strictPort: true,
	},
	// 打包目录
	base: './',
})
