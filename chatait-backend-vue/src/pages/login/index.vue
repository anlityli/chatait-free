<!--
  - Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<template>
  <div class="login-wrapper">
    <login-header />

    <div class="login-container">
      <div class="title-container">
        <h1 class="title margin-no">登录到</h1>
        <h1 class="title">Chat AIT 后台管理系统</h1>
      </div>

      <login v-if="type === 'login'" />
      <register v-else @register-success="switchType('login')" />
      <tdesign-setting />
    </div>

    <footer class="copyright">Copyright @ 2021-2022 Chat AIT. All Rights Reserved</footer>
  </div>
</template>
<script setup lang="ts">
import { ref } from 'vue'

import TdesignSetting from '@/layouts/setting.vue'

import LoginHeader from './components/Header.vue'
import Login from './components/Login.vue'
import Register from './components/Register.vue'

const type = ref('login')
const switchType = (val: string) => {
  type.value = val
}
</script>
<script lang="ts">
export default {
  name: 'LoginIndex',
}
</script>

<style lang="less" scoped>
@import url('./index.less');
</style>
