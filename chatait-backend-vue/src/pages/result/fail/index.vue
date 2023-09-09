<!--
  - Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<template>
  <div class="result-success">
    <t-icon class="result-success-icon" name="error-circle" />
    <div class="result-success-title">项目创建失败</div>
    <div class="result-success-describe">企业微信联系检查创建者权限，或返回修改</div>
    <div>
      <t-button theme="default" @click="() => $router.push('/dashboard/index')">返回首页</t-button>
      <t-button @click="() => $router.push('/form/base')">返回修改</t-button>
    </div>
  </div>
</template>
<script lang="ts">
export default {
  name: 'ResultFail',
}
</script>
<style lang="less" scoped>
.result-success {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 75vh;

  &-icon {
    font-size: var(--td-comp-size-xxxxl);
    color: var(--td-text-color-secondary);
  }

  &-title {
    margin-top: var(--td-comp-margin-xxl);
    font: var(--td-font-title-large);
    color: var(--td-text-color-primary);
    text-align: center;
  }

  &-describe {
    margin: var(--td-comp-margin-s) 0 var(--td-comp-margin-xxxl);
    font: var(--td-font-body-medium);
    color: var(--td-text-color-secondary);
  }
}
</style>
