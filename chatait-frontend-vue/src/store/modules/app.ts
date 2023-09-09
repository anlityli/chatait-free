/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

import { defineStore } from 'pinia'
import { store } from '@/store'

const getSystemThemeMode = (): boolean => {
  return window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches
}

export const useAppStore = defineStore('app', {
  state: () => ({
    theme: 'light',
  }),
  getters: {},
  actions: {
    updateMode(mode = '') {
      if (mode === '') {
        const themeMode = getSystemThemeMode()
        if (themeMode) {
          this.theme = 'dark'
        } else {
          this.theme = 'light'
        }
      } else {
        this.theme = mode
      }
    },
  },
})

export function getAppStore() {
  return useAppStore(store)
}
