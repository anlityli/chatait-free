/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

import TDesign from 'tdesign-vue-next'
import { createApp } from 'vue'

import App from './App.vue'
import router from './router'
import { store } from './store'

import 'tdesign-vue-next/es/style/index.css'
import '@/style/index.less'
import './permission'

const app = createApp(App)

app.use(TDesign)
app.use(store)
app.use(router)

app.mount('#app')
