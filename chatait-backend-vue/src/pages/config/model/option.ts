/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

import { ResponseConfigAllOption } from '@/utils/model/response/config'

export interface OptionListItem {
  group_name: string
  group_items: ResponseConfigAllOption[]
}
