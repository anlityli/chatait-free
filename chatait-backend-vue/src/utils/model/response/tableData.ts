/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

export interface ResponseTableDataColumnsItem {
  field: string
  field_name: string
  field_attr: any
}

export interface ResponseTableDataListValue {
  value: string
  ori_value: any
  value_attr: any
}

export interface ResponseTableDataFilterTypesItem {
  field: string
  field_name: string
  attr: string
  attr_data: any
}

export interface ResponseTableData {
  columns: ResponseTableDataColumnsItem[]
  list: Record<string, ResponseTableDataListValue>[]
  list_id: string
  total_count: number
  page: number
  page_size: number
  total_page: number
  page_count: number
  filter_types: ResponseTableDataFilterTypesItem[]
}
