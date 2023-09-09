/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

import Big from 'big.js'
import storage from '@/utils/storage/storage'
import http from '@/utils/network/http'

export default {
  isLogin(): boolean {
    const token = storage.getToken()
    if (token === null) {
      return false
    }
    if (token.accessToken !== '' && token.accessTokenExpire > this.getTimestamp()) {
      return true
    }
    return token.refreshToken !== '' && token.refreshTokenExpire > this.getTimestamp()
  },
  getTimestamp(date = null): number {
    const days = 0
    let dateObj
    if (date !== null) {
      dateObj = new Date(date)
    } else {
      dateObj = new Date()
    }
    return Math.round(dateObj.getTime() / 1000 + days * 86400)
  },
  formatDate(timestamp: number, withTime = true): string {
    const newDate = new Date()
    if (timestamp) {
      newDate.setTime(timestamp * 1000)
    } else {
      return ''
    }
    const Y = `${newDate.getFullYear()}-`
    const M = `${newDate.getMonth() + 1 < 10 ? `0${newDate.getMonth() + 1}` : newDate.getMonth() + 1}-`
    const D = `${newDate.getDate() < 10 ? `0${newDate.getDate()}` : newDate.getDate()} `
    const h = `${newDate.getHours() < 10 ? `0${newDate.getHours()}` : newDate.getHours()}:`
    const m = `${newDate.getMinutes() < 10 ? `0${newDate.getMinutes()}` : newDate.getMinutes()}:`
    const s = newDate.getSeconds() < 10 ? `0${newDate.getSeconds()}` : newDate.getSeconds()
    if (withTime) return Y + M + D + h + m + s
    return Y + M + D
  },
  secondToMinute(s: number) {
    // 计算分钟
    // 算法：将秒数除以60，然后下舍入，既得到分钟数
    const h = Math.floor(s / 60)
    // 计算秒
    // 算法：取得秒%60的余数，既得到秒数
    s %= 60
    // 将变量转换为字符串
    let hStr = h.toString()
    let sStr = s.toString()
    // 如果只有一位数，前面增加一个0
    hStr = hStr.length === 1 ? `0${hStr}` : hStr
    sStr = sStr.length === 1 ? `0${sStr}` : sStr
    return `${hStr}:${sStr}`
  },
  /**
   * 获取服务器时间对象
   */
  async serverDate() {
    const datetime = await http.getWithoutToken('site/datetime')
    return new Date(datetime as string)
  },
  /**
   * 获取服务器时间戳
   * @param isSecond
   */
  async serverDatetime(isSecond = false) {
    const date = await this.serverDate()
    let time = date.getTime()
    if (isSecond) {
      time /= 1000
    }
    return time
  },
  centToYuan(cent: number): string {
    if (cent === 0) return cent.toString()
    const num = parseFloat(new Big(cent).div(100))
    return num.toFixed(2)
  },
  yuanToCent(yuan: string): number {
    if (yuan === '' || yuan === '0') return Number(yuan)
    return parseFloat(new Big(yuan).times(100))
  },
  image2Base64(img: HTMLImageElement): string {
    const canvas = document.createElement('canvas')
    canvas.width = img.width
    canvas.height = img.height
    const ctx = canvas.getContext('2d')
    ctx?.drawImage(img, 0, 0, img.width, img.height)
    return canvas.toDataURL('image/png')
  },
  getImgBase64(imgPath: string, callback: any) {
    let base64 = ''
    const img = new Image()
    img.src = imgPath
    img.onload = () => {
      base64 = this.image2Base64(img)
      callback(base64)
    }
  },
  /**
   * 解析URL地址
   * @param url
   * var myURL = parseURL('http://abc.def.xxx.com:8080/dir/index.html?id=255&m=hello#top');
   myURL.file; // = 'index.html'
   myURL.hash; // = 'top'
   myURL.host; // = 'xxx.com'
   myURL.sub; // = 'abc.def'
   myURL.subs; // = ['abc', 'def']
   myURL.query; // = '?id=255&m=hello'
   myURL.params; // = Object = { id: 255, m: hello }
   myURL.path; // = '/dir/index.html'
   myURL.segments; // = Array = ['dir', 'index.html']
   myURL.port; // = '8080'
   myURL.protocol; // = 'http'
   myURL.source; // = 'http://abc.def.xxx.com:8080/dir/index.html?id=255&m=hello#top'
   * @returns {{source: *, protocol: string, host: string, port: string, query: string, params, file: string | *, hash: string, path: string, relative: string | *, segments: string[]}}
   */
  parseURL(url: string) {
    const a = document.createElement('a')
    a.href = url
    const host = a.hostname
    const hostArr = host.split('.')
    const sub = hostArr.shift()
    const domain = hostArr.join('.')
    return {
      source: url,
      protocol: a.protocol.replace(':', ''),
      host: a.hostname,
      domain,
      sub,
      port: a.port,
      query: a.search,
      params: (function p() {
        const waitUrl = a.hash.replace(/^#[/\-_a-zA-Z0-9]+\?/, '?') || a.search
        const ret: Record<string, string> = {}
        const seg = waitUrl.replace(/^\?/, '').split('&')
        const len = seg.length
        let i = 0
        let s
        for (; i < len; i++) {
          if (!seg[i]) {
            continue
          }
          s = seg[i].split('=')
          const index = s[0] as string
          ;[, ret[index]] = s
        }
        return ret
      })(),
      file: (a.pathname.match(/\/([^/?#]+)$/i) || [undefined, ''])[1],
      hash: a.hash.replace('#', ''),
      path: a.pathname.replace(/^([^/])/, '/$1'),
      relative: (a.href.match(/tps?:\/\/[^/]+(.+)/) || [undefined, ''])[1],
      segments: a.pathname.replace(/^\//, '').split('/'),
    }
  },
}
