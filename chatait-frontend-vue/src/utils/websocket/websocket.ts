/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

import http from '@/utils/network/http'
import tool from '@/utils/tool/tool'

class WebsocketCli {
  constructor(url: string) {
    this.path = url
  }

  path = ''
  websocketObj: WebSocket | null = null

  static websocketInstances: Record<string, any> = {}

  static requestApiHost(): string {
    const wsHost = import.meta.env.VITE_WEBSOCKET_HOST
    if (wsHost !== undefined && wsHost !== null && wsHost !== '') {
      return wsHost
    }
    const url = window.location.href
    const p = tool.parseURL(url)
    const { protocol, domain, port } = p
    let portStr = ''
    if (port !== '') {
      portStr = `:${port}`
    }
    let protocolStr = ''
    if (protocol === 'http') {
      protocolStr = 'ws'
    } else {
      protocolStr = 'wss'
    }
    return `${protocolStr}://chat-frontend-api.${domain}${portStr}`
  }

  initWebsocketObj() {
    this.websocketObj = new WebSocket(`${WebsocketCli.requestApiHost()}/${this.path}`)
  }

  getObj(): WebSocket | null {
    return this.websocketObj
  }

  static getInstance(url: string): WebsocketCli {
    if (this.websocketInstances[url] === undefined || this.websocketInstances[url] === null) {
      this.websocketInstances[url] = new WebsocketCli(url)
      this.websocketInstances[url].initWebsocketObj()
    }
    return this.websocketInstances[url]
  }

  static sendMessage(socketObj: WebSocket, data: any) {
    http.getAccessToken().then((token) => {
      const { time, sign } = http.signParams({ token })
      data.access_token = token
      data.time = time
      data.sign = sign
      socketObj.send(JSON.stringify(data))
    })
  }

  static removeObj(url: string) {
    this.websocketInstances[url] = null
  }
}

export default WebsocketCli
