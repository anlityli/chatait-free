/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

import { createPinia } from 'pinia'
import { createPersistedState } from 'pinia-plugin-persistedstate'

const store = createPinia()
store.use(createPersistedState())

export { store }

export * from './modules/admin'
export * from './modules/config'
export * from './modules/notification'
export * from './modules/permission'
export * from './modules/setting'
export * from './modules/tabs-router'

export default store
