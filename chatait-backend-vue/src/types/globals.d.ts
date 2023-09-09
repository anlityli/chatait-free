/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

// 通用声明

// Vue
declare module '*.vue' {
  import { DefineComponent } from 'vue'

  const component: DefineComponent<{}, {}, any>
  export default component
}

declare type ClassName = { [className: string]: any } | ClassName[] | string

declare module '*.svg' {
  const CONTENT: string
  export default CONTENT
}

declare type Recordable<T = any> = Record<string, T>
