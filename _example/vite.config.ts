import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import path from 'path'
import entriesConfig from './entries.gen.json'

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
      input: entriesConfig,
    },
  },
  server: {
    host: true,
  },
})
