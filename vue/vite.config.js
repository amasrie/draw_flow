import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  server: {
    proxy: {
      cors:false
    },
    port: 9601,
  },  plugins: [vue()],
  build: {
    outDir: 'docs'
}
})
