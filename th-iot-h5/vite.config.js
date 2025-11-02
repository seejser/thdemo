import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

export default defineConfig({
  plugins: [vue()],
  server: {
    host: '0.0.0.0', // 允许通过 IP 访问
    port: 5173,      // 你可以改成任何端口号
    open: false,     // 是否自动打开浏览器
  },
})
