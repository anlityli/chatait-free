/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

export interface ResponsePage {
  list_data: any
  total_count: number
  page: number
  page_size: number
  total_page: number
  page_count: number
}
