/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

import vue from '@vitejs/plugin-vue'
import vueJsx from '@vitejs/plugin-vue-jsx'
import path from 'path'
import { ConfigEnv, loadEnv, UserConfig } from 'vite'
import svgLoader from 'vite-svg-loader'

const CWD = process.cwd()

// https://vitejs.dev/config/
export default ({ mode }: ConfigEnv): UserConfig => {
  const { VITE_BASE_URL } = loadEnv(mode, CWD)
  return {
    base: VITE_BASE_URL,
    resolve: {
      alias: {
        '@': path.resolve(__dirname, './src'),
      },
    },

    css: {
      preprocessorOptions: {
        less: {
          modifyVars: {
            hack: `true; @import (reference) "${path.resolve('src/style/variables.less')}";`,
          },
          math: 'strict',
          javascriptEnabled: true,
        },
      },
    },

    plugins: [vue(), vueJsx(), svgLoader()],

    server: {
      port: 3002,
      // host: '0.0.0.0',
      // proxy: {
      //   [VITE_API_URL_PREFIX]: 'http://127.0.0.1:3000/',
      // },
    },
  }
}
