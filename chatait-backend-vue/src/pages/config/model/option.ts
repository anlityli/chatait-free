/*
 * Copyright 2024 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

import { ConfigAllOption } from '@/pages/config/model/allOption'

export interface OptionListItem {
  group_name: string
  group_items: ConfigAllOption[]
}
