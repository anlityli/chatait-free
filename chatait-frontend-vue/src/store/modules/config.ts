/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

import { defineStore } from 'pinia'
import http from '@/utils/network/http'
import { ResponseConfigWalletListItem } from '@/utils/model/response/config'
import { store } from '@/store'

export const useConfigStore = defineStore('config', {
  state: () => ({
    walletList: <ResponseConfigWalletListItem[]>[],
  }),
  actions: {
    async getWalletList() {
      this.walletList = (await http.get('config/wallet-list')) as ResponseConfigWalletListItem[]
    },
    walletName(field: string): string {
      for (let i = 0; i < this.walletList.length; i++) {
        if (field === this.walletList[i].field) {
          return this.walletList[i].wallet_name
        }
      }
      return ''
    },
  },
})

export function getConfigStore() {
  return useConfigStore(store)
}
