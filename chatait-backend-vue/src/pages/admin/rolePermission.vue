<!--
  - Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<template>
  <div class="main-content-container">
    <div class="role-permission-list">
      <div v-for="(item, index) in permissionList" :key="index" class="role-permission-list-item">
        <div class="list-left">
          <t-checkbox
            :checked="mainPermissionCheckoutStatus[index].checkout_all"
            :indeterminate="mainPermissionCheckoutStatus[index].indeterminate"
            @change="
            (checked: boolean) => {
                    handleCheckoutMainAll(checked, item, index)
                  }

"
          >
            {{ item.main_permission.title }}
          </t-checkbox>
        </div>
        <div class="list-right">
          <t-checkbox-group
            v-model="mainPermissionCheckoutStatus[index].checked_value"
            @change="(checkedItems: string[]) => {
                    handleCheckoutChild(checkedItems, item, index)
                  }"
          >
            <t-checkbox
              v-for="(childItem, childIndex) in item.child_permission"
              :key="childIndex"
              class="child-checkbox"
              :value="childItem.path"
              >{{ childItem.title }}
            </t-checkbox>
          </t-checkbox-group>
        </div>
      </div>
      <t-button @click="handleSubmit">提交</t-button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import http from '@/utils/network/http'
import { ResponseAdminRolePermission } from '@/utils/model/response/admin'
import { RolePermissionMainCheckoutItem } from '@/pages/admin/model/model'

const router = useRouter()
const route = useRoute()

const permissionList = ref<ResponseAdminRolePermission[]>([])
const mainPermissionCheckoutStatus = ref<RolePermissionMainCheckoutItem[]>([])

const handleCheckoutMainAll = (checked: boolean, mainItem: ResponseAdminRolePermission, index: number) => {
  mainPermissionCheckoutStatus.value[index].checked_value = []
  for (let i = 0; i < permissionList.value[index].child_permission.length; i++) {
    permissionList.value[index].child_permission[i].is_checked = checked
    if (checked) {
      mainPermissionCheckoutStatus.value[index].checked_value.push(permissionList.value[index].child_permission[i].path)
    }
  }
  mainPermissionCheckoutStatus.value[index].checkout_all = checked
  mainPermissionCheckoutStatus.value[index].indeterminate =
    mainPermissionCheckoutStatus.value[index].checked_value.length > 0 &&
    !mainPermissionCheckoutStatus.value[index].checkout_all
}

const handleCheckoutChild = (checkedItems: string[], item: ResponseAdminRolePermission, index: number) => {
  mainPermissionCheckoutStatus.value[index].checked_value = checkedItems
  mainPermissionCheckoutStatus.value[index].checkout_all =
    mainPermissionCheckoutStatus.value[index].checked_value.length ===
    permissionList.value[index].child_permission.length
  mainPermissionCheckoutStatus.value[index].indeterminate =
    mainPermissionCheckoutStatus.value[index].checked_value.length > 0 &&
    !mainPermissionCheckoutStatus.value[index].checkout_all
}

const handleRefreshMainCheckout = () => {
  for (let i = 0; i < permissionList.value.length; i++) {
    mainPermissionCheckoutStatus.value[i].checked_value = []
    let childCheckoutCount = 0
    for (let j = 0; j < permissionList.value[i].child_permission.length; j++) {
      if (permissionList.value[i].child_permission[j].is_checked) {
        childCheckoutCount += 1
        mainPermissionCheckoutStatus.value[i].checked_value.push(permissionList.value[i].child_permission[j].path)
      }
    }
    mainPermissionCheckoutStatus.value[i].checkout_all =
      childCheckoutCount === permissionList.value[i].child_permission.length
    mainPermissionCheckoutStatus.value[i].indeterminate =
      childCheckoutCount > 0 && !mainPermissionCheckoutStatus.value[i].checkout_all
  }
}

const handleSubmit = async () => {
  const requestData = {
    id: route.query.id,
    permission: <string[]>[],
  }
  for (let i = 0; i < mainPermissionCheckoutStatus.value.length; i++) {
    requestData.permission.push(...mainPermissionCheckoutStatus.value[i].checked_value)
  }
  await http.post('admin/role-permission-edit', requestData)
  router.back()
}

onMounted(async () => {
  permissionList.value = (await http.get(`admin/role-permission?id=${route.query.id}`)) as ResponseAdminRolePermission[]
  for (let i = 0; i < permissionList.value.length; i++) {
    mainPermissionCheckoutStatus.value.push({
      key: '',
      checkout_all: false,
      indeterminate: false,
      checked_value: [],
    })
  }
  handleRefreshMainCheckout()
})
</script>

<style lang="less" scoped>
.role-permission-list {
  width: 100%;
  border: 1px solid var(--td-border-level-2-color);
  border-bottom: none;

  .role-permission-list-item {
    display: flex;
    width: 100%;
    border-bottom: 1px solid var(--td-border-level-2-color);

    .list-left {
      flex: 0 0 auto;
      width: 200px;
      padding: 10px;
      box-sizing: border-box;
      border-right: 1px solid var(--td-border-level-2-color);
    }

    .list-right {
      flex: 1 1 auto;
      width: 100%;
      padding: 10px;
      box-sizing: border-box;

      .child-checkbox {
        margin-right: 10px;
      }
    }
  }
}
</style>
