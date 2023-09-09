/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

// 通用声明
declare type ClassName = { [className: string]: any } | ClassName[] | string

declare interface ImportMeta {
  // eslint-disable-next-line no-unused-vars
  glob: (url: string) => { url }
}

declare module '*.svg' {
  const CONTENT: string
  export default CONTENT
}

declare module 'sortablejs'

declare module 'marked'

declare module 'big.js'

declare module 'qartjs'
