/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

import { ResponseConfigPayListItemParams, ResponseConfigPayListItemPayChannel } from '@/utils/model/response/config'

export interface FormPayEdit {
  id: string
  params: ResponseConfigPayListItemParams[]
  pay_channel: ResponseConfigPayListItemPayChannel[]
  frontend_description: string
  status: number
}
