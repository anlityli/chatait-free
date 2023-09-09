/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

export interface FilterFormValue {
  symbol: string
  value: string[]
  model?: string
  attrData?: any
}

export interface TableFixedColumn {
  field: string
  direction: 'left' | 'right' // left | right
}

export interface TableRowSelect {
  enable: boolean
  type?: 'single' | 'multiple'
}
