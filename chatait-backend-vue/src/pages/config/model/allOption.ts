/*
 * Copyright 2024 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

import { ResponseConfigAllOptionOptionsItem } from '@/utils/model/response/config'

export interface ConfigAllOption {
  config_name: string
  created_at: number
  input_type: number
  options: ResponseConfigAllOptionOptionsItem[]
  sort: number
  title: string
  type: string
  unit: string
  updated_at: number
  value: any
}
