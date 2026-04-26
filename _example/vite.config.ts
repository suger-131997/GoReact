import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import path from 'path'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: {
      '~': path.resolve(__dirname, "./"),
    },
  },
  build: {
    manifest: true,
    rolldownOptions: {
      input: 'tmp/src/main.tsx',
    },
  },
  server: {
    host: true,
  },
})
