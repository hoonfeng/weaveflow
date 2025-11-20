import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

export default defineConfig({
  plugins: [vue()],
  server: {
    port: 5173,
    proxy: {
      '/api': 'http://localhost:8081',
      '/admin': 'http://localhost:8081',
      '/docs': 'http://localhost:8081',
      '/metrics': 'http://localhost:8081',
      '/openapi.json': 'http://localhost:8081'
    }
  }
})