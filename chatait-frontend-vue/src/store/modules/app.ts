/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

import { defineStore } from 'pinia'
import { store } from '@/store'
import { ApplicationTypeMJ } from '@/utils/constant/conversation'
import { MidjourneySpeakForm } from '@/views/conversation/script/model'

const getSystemThemeMode = (): boolean => {
  return window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches
}

export const useAppStore = defineStore('app', {
  state: () => ({
    theme: 'light',
    saveMjSpeakParams: false,
    mjSpeakParams: <MidjourneySpeakForm>{
      topic_id: '',
      content: '',
      application_type: ApplicationTypeMJ,
      no: '',
      images: '',
      seed: '',
      ar: '',
      chaos: '',
      quality: '',
      model: '',
      style: '',
      stylize: '',
      tile: '',
      iw: '',
    },
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
    switchSaveMjSpeakParams() {
      console.log(this.saveMjSpeakParams)
      this.saveMjSpeakParams = !this.saveMjSpeakParams
      console.log(this.saveMjSpeakParams)
    },
    setMjSpeakParams(formValue: MidjourneySpeakForm) {
      if (this.saveMjSpeakParams) {
        this.mjSpeakParams = formValue
      }
    },
  },
  persist: {
    key: 'app',
    storage: localStorage,
    paths: ['saveMjSpeakParams', 'mjSpeakParams'],
  },
})

export function getAppStore() {
  return useAppStore(store)
}
