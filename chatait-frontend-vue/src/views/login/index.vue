<!--
  - Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { LockOnIcon, UserIcon } from 'tdesign-icons-vue-next'
import { useRouter } from 'vue-router'
import { LoginForm } from '@/views/login/script/model'
import http from '@/utils/network/http'
import storage from '@/utils/storage/storage'
import tool from '@/utils/tool/tool'
import { ResponseUser } from '@/utils/model/response/user'
import logoPath from '@/assets/image/logo.png'

const router = useRouter()

const formData = ref<LoginForm>({
  username: '',
  password: '',
})

const rules = {
  username: [{ required: true, message: '用户名必填', type: 'error' }],
  password: [{ required: true, message: '密码必填', type: 'error' }],
}

const pageLoading = ref(true)
const submitLoading = ref(false)
const emailCodeEnable = ref(false)

const handleValidate = (validate: any) => {
  if (validate.validateResult === true) {
    console.log('注册')
  } else {
    console.log('失败', validate.validateResult, validate.firstError)
  }
}

const handleSignup = () => {
  router.push('/signup')
}

const handleFindPassword = () => {
  router.push('/find-password')
}

const handleSubmit = async (submitResult: any) => {
  submitResult.e.preventDefault()
  submitLoading.value = true
  if (submitResult.validateResult === true) {
    try {
      const loginResponseData = await http.postWithoutToken('oauth/login', formData.value)
      storage.setToken({
        accessToken: loginResponseData.access_token,
        accessTokenExpire: loginResponseData.access_token_expire_in + tool.getTimestamp(),
        accessTokenExpireIn: loginResponseData.access_token_expire_in,
        refreshToken: loginResponseData.refresh_token,
        refreshTokenExpire: loginResponseData.refresh_token_expire_in + tool.getTimestamp(),
        refreshTokenExpireIn: loginResponseData.refresh_token_expire_in,
      })
      const userInfoData = (await http.get('user/info')) as ResponseUser
      storage.setUserInfo({
        id: userInfoData.id,
        username: userInfoData.username,
        nickname: userInfoData.nickname,
        avatar: userInfoData.avatar,
      })
      await router.push('/conversation/0')
      submitLoading.value = false
    } catch (e) {
      submitLoading.value = false
    }
  } else {
    submitLoading.value = false
  }
}

onMounted(async () => {
  if (tool.isLogin()) {
    await router.push('/conversation/0')
  }
  http.getWithoutToken('config/options').then((response) => {
    emailCodeEnable.value = response.emailCodeEnable === '1'
  })
  pageLoading.value = false
})
</script>

<template>
  <div class="login-wrap">
    <t-card class="login-card" header-bordered :loading="pageLoading">
      <template #content>
        <div class="login-card-content">
          <div class="login-logo-space">
            <img :src="logoPath" alt="ChatAIT" />
          </div>
          <div class="login-form-space">
            <t-form
              ref="form"
              :data="formData"
              :colon="true"
              :label-width="0"
              :rules="rules"
              @submit="handleSubmit"
              @validate="handleValidate"
            >
              <t-form-item name="username">
                <t-input v-model="formData.username" clearable placeholder="请输入账户名">
                  <template #prefix-icon>
                    <user-icon />
                  </template>
                </t-input>
              </t-form-item>

              <t-form-item name="password">
                <t-input v-model="formData.password" type="password" clearable placeholder="请输入密码">
                  <template #prefix-icon>
                    <lock-on-icon />
                  </template>
                </t-input>
              </t-form-item>

              <t-form-item class="login-btn-space">
                <t-space direction="vertical" align="center">
                  <t-button theme="primary" type="submit" block :loading="submitLoading">登录</t-button>
                  <div class="login-other-btn">
                    <div class="login-btn-signup" @click="handleSignup">新用户注册</div>
                    <div v-if="emailCodeEnable" class="login-btn-find-password" @click="handleFindPassword">
                      忘记密码?
                    </div>
                  </div>
                </t-space>
              </t-form-item>
            </t-form>
          </div>
        </div>
      </template>
    </t-card>
  </div>
</template>

<style lang="scss">
.login-wrap {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 100%;
  overflow: hidden;

  .login-card {
    flex: 0 0 auto;
    width: 400px;
    height: 300px;

    .login-card-content {
      display: flex;
      flex-direction: column;
      width: 100%;
      height: 100%;
      overflow: hidden;

      .login-logo-space {
        flex: 0 0 auto;
        display: flex;
        width: 100%;
        height: 80px;
        justify-content: center;
        align-items: center;
        overflow: hidden;
      }

      .login-btn-space .t-space {
        width: 100%;

        .login-other-btn {
          display: flex;
          width: 100%;

          .login-btn-signup {
            flex: 1 1 auto;
            width: 100%;
            text-align: center;
            cursor: pointer;
          }

          .login-btn-find-password {
            flex: 1 1 auto;
            width: 100%;
            text-align: center;
            cursor: pointer;
          }
        }
      }
    }
  }

  @media screen and (max-width: 768px) {
    .login-card {
      width: 90%;
    }
  }
}
</style>
