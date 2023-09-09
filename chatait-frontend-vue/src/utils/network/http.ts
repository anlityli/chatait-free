/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

import axios, { AxiosResponse } from 'axios'
import shaJs from 'sha.js'
import { Base64 } from 'js-base64'
import { DialogPlugin, MessagePlugin } from 'tdesign-vue-next'
import router from '@/router'
import { RequestSign } from '@/utils/network/model'
import storage from '@/utils/storage/storage'
import tool from '@/utils/tool/tool'

class ResponseError extends Error {
  private code: number

  constructor(message: string, code: number) {
    super(message)
    this.code = code
  }
}

export default {
  refreshTokenLock: <Promise<string> | null>null,
  /**
   * 请求API的Host
   */
  requestApiHost(): string {
    const host = import.meta.env.VITE_HTTP_HOST
    if (host !== undefined && host !== null && host !== '') {
      return host
    }
    const url = window.location.href
    const p = tool.parseURL(url)
    const { protocol, domain, port } = p
    let portStr = ''
    if (port !== '') {
      portStr = `:${port}`
    }
    return `${protocol}://chat-frontend-api.${domain}${portStr}`
  },
  /**
   * get请求数据
   * @param path
   * @param params
   */
  async get(path: string, params?: any): Promise<any> {
    try {
      const accessToken = await this.getAccessToken()
      const signObj = this.signParams({ token: accessToken })
      const config = {
        headers: {
          Authorization: `Bearer ${accessToken}`,
          'x-site-sign': signObj.sign,
          'x-site-time': signObj.time,
        },
        params: null,
      }
      if (params !== undefined) {
        config.params = params
      }
      return await this.axiosGet(path, config)
    } catch (e: any) {
      if (e.code !== 405) {
        await MessagePlugin.error(e.message)
      }
      throw new ResponseError(`error:${e.message}`, e.code)
    }
  },
  /**
   * post请求数据
   * @param path
   * @param data
   */
  async post(path: string, data: any): Promise<any> {
    try {
      const accessToken = await this.getAccessToken()
      const signObj = this.signParams({ token: accessToken })
      const config = {
        headers: {
          Authorization: `Bearer ${accessToken}`,
          'x-site-sign': signObj.sign,
          'x-site-time': signObj.time,
        },
        params: null,
      }
      return await this.axiosPost(path, data, config)
    } catch (e: any) {
      if (e.code !== 405) {
        await MessagePlugin.error(e.message)
      }
      throw new ResponseError(`error:${e.message}`, e.code)
    }
  },
  /**
   * 不加token的get请求
   * @param path
   * @param params
   */
  async getWithoutToken(path: string, params?: any): Promise<any> {
    try {
      const signObj = this.signParams({ token: '' })
      const config = {
        headers: {
          Authorization: '',
          'x-site-sign': signObj.sign,
          'x-site-time': signObj.time,
        },
        params: null,
      }
      if (params !== undefined) {
        config.params = params
      }
      return await this.axiosGet(path, config)
    } catch (e: any) {
      if (e.code !== 405) {
        await MessagePlugin.error(e.message)
      }
      throw new ResponseError(`error:${e.message}`, e.code)
    }
  },
  /**
   * 不加token的post请求
   * @param path
   * @param data
   */
  async postWithoutToken(path: string, data: any): Promise<any> {
    try {
      const signObj = this.signParams({ token: '' })
      const config = {
        headers: {
          Authorization: '',
          'x-site-sign': signObj.sign,
          'x-site-time': signObj.time,
        },
        params: null,
      }
      return await this.axiosPost(path, data, config)
    } catch (e: any) {
      if (e.code !== 405) {
        await MessagePlugin.error(e.message)
      }
      throw new ResponseError(`error:${e.message}`, e.code)
    }
  },
  async axiosGet<T = any, R = AxiosResponse<T>>(path: string, config: any): Promise<R> {
    const responseData = await axios.get(`${this.requestApiHost()}/${path}`, config)
    await this.requestErrorHandle(responseData)
    return responseData.data.message
  },
  async axiosPost<T = any, R = AxiosResponse<T>>(path: string, data: any, config: any): Promise<R> {
    const responseData = await axios.post(`${this.requestApiHost()}/${path}`, data, config)
    await this.requestErrorHandle(responseData)
    return responseData.data.message
  },
  async requestErrorHandle(responseData: any) {
    if (responseData.status !== 200) {
      throw new ResponseError(`${responseData.status.toString()}`, responseData.status)
    }
    if (responseData.data.error !== 0) {
      if (responseData.data.error === 401) {
        storage.clearToken()
        storage.clearUserInfo()
        await router.push('/login')
        throw new ResponseError(`${responseData.data.message.toString()}`, responseData.data.error)
      } else if (responseData.data.error === 405) {
        const dialogInstance = DialogPlugin.confirm({
          body: responseData.data.message.data,
          confirmBtn: responseData.data.message.confirm_jump !== '' ? responseData.data.message.confirm_text : null,
          onConfirm: async () => {
            dialogInstance.destroy()
            await router.push(responseData.data.message.confirm_jump)
          },
        })
        throw new ResponseError(`${responseData.data.message.data}`, responseData.data.error)
      } else {
        throw new ResponseError(`${responseData.data.message.toString()}`, responseData.data.error)
      }
    }
  },
  async getAccessToken(): Promise<string> {
    const token = storage.getToken()
    if (token === null) {
      return ''
    }
    if (token.accessToken !== '' && token.accessTokenExpire > tool.getTimestamp()) {
      return token.accessToken
    }
    if (token.refreshToken !== '' && token.refreshTokenExpire > tool.getTimestamp()) {
      if (this.refreshTokenLock === null) {
        this.refreshTokenLock = this.refreshToken()
      }
      const tokenStr = await this.refreshTokenLock
      this.refreshTokenLock = null
      return tokenStr
    }
    storage.clearToken()
    storage.clearUserInfo()
    await router.push('/login')
    return ''
  },
  async refreshToken(): Promise<string> {
    const token = storage.getToken()
    if (token === null) {
      return ''
    }
    try {
      const responseData = (await this.postWithoutToken('oauth/refresh-token', {
        refresh_token: token.refreshToken,
      })) as any
      storage.setToken({
        accessToken: responseData.access_token,
        accessTokenExpire: responseData.access_token_expire_in + tool.getTimestamp(),
        accessTokenExpireIn: responseData.access_token_expire_in,
        refreshToken: responseData.refresh_token,
        refreshTokenExpire: responseData.refresh_token_expire_in + tool.getTimestamp(),
        refreshTokenExpireIn: responseData.refresh_token_expire_in,
      })
      return responseData.access_token as string
    } catch (error: any) {
      if (error.code === 401) {
        storage.clearToken()
        storage.clearUserInfo()
        await router.push('/login')
      }
      return ''
    }
  },
  signParams(params: any): RequestSign {
    const time = Math.round(Date.now() / 1000).toString()
    const signParamsData = { ...params, time, private_key: import.meta.env.VITE_SERVER_API_PRIVATE_KEY }
    const newKey = Object.keys(signParamsData).sort()
    const resArr = []
    for (let i = 0; i < newKey.length; i++) {
      resArr.push(`${newKey[i]}=${encodeURI(signParamsData[newKey[i]])}`)
    }
    // console.log(resArr.join('&'))
    // console.log(CryptoJS.enc.Base64.stringify)
    return {
      time,
      sign: shaJs('sha1')
        .update(Base64.encode(resArr.join('&')))
        .digest('hex'),
    }
  },
}
