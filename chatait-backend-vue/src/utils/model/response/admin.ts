/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

export interface ResponseAdminInfo {
  id: string
  user_id: string
  username: string
  admin_name: string
  real_name: string
  remark: string
  role_id: string
  role_name: string
  is_enable: number
  login_nums: number
  last_login_ip: string
  last_login_at: number
  bind_ip: string
  create_admin: string
  update_admin: string
  created_at: number
  updated_at: number
  dont_del: number
  menu: any
  admin_permission: any
}

export interface ResponseAdminItem {
  id: string
  user_id: string
  username: string
  admin_name: string
  real_name: string
  remark: string
  role_id: string
  role_name: string
  is_enable: number
  login_nums: number
  last_login_ip: string
  last_login_at: number
  bind_ip: string
  create_admin: string
  update_admin: string
  created_at: number
  updated_at: number
  dont_del: number
}

export interface ResponseAdminRoleItem {
  id: string
  role_name: string
  remark: string
  permission: string
  column_permission: string
  dont_del: number
  create_admin: string
  update_admin: string
  created_at: number
  updated_at: number
}

export interface ResponseAdminRolePermissionData {
  title: string
  path: string
  is_checked: boolean
}

export interface ResponseAdminRolePermission {
  id: string
  main_permission: ResponseAdminRolePermissionData
  child_permission: ResponseAdminRolePermissionData[]
}

export interface ResponseAdminRoleColumnColumnItem {
  header: string
  index: string
  is_checked: boolean
}

export interface ResponseAdminRoleColumn {
  list_id: string
  list_name: string
  columns: ResponseAdminRoleColumnColumnItem[]
}
