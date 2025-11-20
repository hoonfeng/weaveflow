import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

export default defineConfig({
  plugins: [vue()],
  server: {
    port: 5174,
    proxy: {
      '/api': 'http://localhost:8080',
      '/docs': 'http://localhost:8080',
      '/metrics': 'http://localhost:8080',
      '/openapi.json': 'http://localhost:8080'
    }
  }
})