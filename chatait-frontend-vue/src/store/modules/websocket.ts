/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

import { defineStore } from 'pinia'
import { store } from '@/store'
import WebsocketCli from '@/utils/websocket/websocket'
import eventBus from '@/utils/eventBus/eventBus'

export const useWebsocketStore = defineStore('websocket', {
  state: () => ({
    websocketObj: <WebSocket | null>null,
    manualClose: false,
    heartTimer: <NodeJS.Timer | null>null,
  }),
  getters: {},
  actions: {
    connect() {
      return new Promise((resolve, reject) => {
        this.websocketObj = WebsocketCli.getInstance('websocket/index').getObj()
        if (this.websocketObj !== null) {
          this.websocketObj.onopen = (event) => {
            console.log('ws连接成功', event)
            resolve(event)
            if (this.heartTimer === null) {
              this.heartTimer = setInterval(() => {
                try {
                  this.websocketObj?.send('ping')
                } catch (e) {
                  clearInterval(Number(this.heartTimer))
                  console.log('ws监听器已关闭因为连接已断开', e)
                }
              }, 30000)
            }
          }
          this.websocketObj.onclose = (event) => {
            console.log('服务断开连接', event)
            clearInterval(Number(this.heartTimer))
            this.heartTimer = null
            WebsocketCli.removeObj('websocket/index')
            // 手动断开不在重连
            if (this.manualClose) {
              this.manualClose = false
            } else if (this.heartTimer === null) {
              this.heartTimer = setInterval(() => {
                this.connect().then(() => {
                  clearInterval(Number(this.heartTimer))
                  this.heartTimer = null
                })
              }, 5000)
            }
          }
          this.websocketObj.onerror = (event) => {
            console.log(event)
          }
          this.websocketObj.onmessage = (event) => {
            if (event.data !== 'pong') {
              eventBus.emit('conversationMj', event.data)
            }
          }
        }
      })
    },
  },
})

export function getWebsocketStore() {
  return useWebsocketStore(store)
}
