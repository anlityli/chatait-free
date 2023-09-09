/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

export interface FormAdminEdit {
  id: string
  user_id: string
  admin_name: string
  real_name: string
  remark: string
  role_id: string
  is_enable: number
  password: string
  bind_ip: string[]
}

export interface FormAdminRoleEdit {
  id: string
  role_name: string
  remark: string
}

export interface RolePermissionMainCheckoutItem {
  key: string
  checkout_all: boolean
  indeterminate: boolean
  checked_value: string[]
}

export interface FormAdminResetPassword {
  password: string
  re_password: string
}
