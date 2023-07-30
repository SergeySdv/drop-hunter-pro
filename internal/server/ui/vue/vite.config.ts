// Plugins
import vue from '@vitejs/plugin-vue'
import vuetify, {transformAssetUrls} from 'vite-plugin-vuetify'

// Utilities
import {defineConfig} from 'vite'
import {fileURLToPath, URL} from 'node:url'
import {uglify} from "rollup-plugin-uglify";

// https://vitejs.dev/config/
export default defineConfig({
  // https://github.com/mishoo/UglifyJS/blob/master/README.md#minify-options
  plugins: [
    uglify({
      toplevel: true,
      annotations: false,
      sourcemap: false,
      output: {
        comments: "all"
      }
    }),
    vue({
      template: {transformAssetUrls}
    }),
    // https://github.com/vuetifyjs/vuetify-loader/tree/next/packages/vite-plugin
    vuetify({
      autoImport: true,
    }),
  ],
  define: {
    'process.env': {},
  },
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    },
    extensions: [
      '.js',
      '.json',
      '.jsx',
      '.mjs',
      '.ts',
      '.tsx',
      '.vue',
      '.sass'
    ],
  },
  server: {
    port: 3021,
    hmr: {
      overlay: false
    }
  },
})
