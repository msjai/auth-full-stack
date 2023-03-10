import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  base: '/account/',
  plugins: [vue()],
  server: {
    host: true,
    port: 5173
  }
})


