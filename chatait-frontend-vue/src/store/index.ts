/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

import { createPinia } from 'pinia'
import { createPersistedState } from 'pinia-plugin-persistedstate'
// import { provideEventBus } from '@/utils/eventBus/eventBus'

const store = createPinia()
store.use(createPersistedState())
// provideEventBus()
export { store }

export * from './modules/app'
export * from './modules/config'
export * from './modules/topic'
export * from './modules/websocket'

export default store
