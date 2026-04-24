import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    proxy: {
      // Proxy API requests to the Go backend
      '/movies': {
        target: 'http://localhost:3000',
        changeOrigin: true,
      },
      '/sessions': {
        target: 'http://localhost:3000',
        changeOrigin: true,
      },
    },
  },
})
