/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

import { createApp } from 'vue'
import TDesign from 'tdesign-vue-next'
import App from './App.vue'
import router from './router'

// 引入组件库全局样式资源
import 'tdesign-vue-next/es/style/index.css'

createApp(App).use(TDesign).use(router).mount('#app')
