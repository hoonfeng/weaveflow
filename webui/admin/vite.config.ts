import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

export default defineConfig({
  plugins: [vue()],
  server: {
    port: 5175,
    strictPort: true,
    proxy: {
      '/api': { target: 'http://localhost:8081', changeOrigin: true },
      '/admin': { target: 'http://localhost:8081', changeOrigin: true },
      '/docs': { target: 'http://localhost:8081', changeOrigin: true },
      '/metrics': { target: 'http://localhost:8081', changeOrigin: true },
      '/openapi.json': { target: 'http://localhost:8081', changeOrigin: true }
    }
  }
})