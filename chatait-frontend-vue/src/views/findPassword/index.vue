<!--
  - Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { LockOnIcon, UserIcon } from 'tdesign-icons-vue-next'
import { useRouter } from 'vue-router'
import http from '@/utils/network/http'
import storage from '@/utils/storage/storage'
import tool from '@/utils/tool/tool'
import { ResponseUser } from '@/utils/model/response/user'
import logoPath from '@/assets/image/logo.png'
import { FindPasswordForm } from '@/views/findPassword/script/model'

const router = useRouter()

const formData = ref<FindPasswordForm>({
  username: '',
  password: '',
  confirm_password: '',
  code: '',
})

const submitLoading = ref(false)
const sendCodeLoading = ref(false)
const sendCodeBtnText = ref('发送验证码')
const sendCodeBtnDisable = ref(false)

const confirmPasswordValidator = (val: string): boolean => {
  return formData.value.password === val
}

const rules = {
  username: [{ required: true, message: '用户名必填', type: 'error' }],
  password: [{ required: true, message: '密码必填', type: 'error' }],
  confirm_password: [
    { required: true, message: '确认密码必填', type: 'error' },
    { validator: confirmPasswordValidator, message: '两次密码不一致' },
  ],
  code: [{ required: true, message: '验证码必填', type: 'error' }],
}

const handleValidate = (validate: any) => {
  if (validate.validateResult === true) {
    console.log('注册')
  } else {
    console.log('失败', validate.validateResult, validate.firstError)
  }
}

const handleLogin = () => {
  router.push('/login')
}

const handleServerRemainSecond = async (targetTime: number) => {
  const serverTime = await tool.serverDatetime(true)
  return targetTime - serverTime
}

const handleSendCode = async () => {
  sendCodeLoading.value = true
  try {
    const sendCodeResponseData = await http.postWithoutToken('oauth/find-password-send-code', formData.value)
    sendCodeLoading.value = false
    sendCodeBtnDisable.value = true
    const serverDatetime = await tool.serverDatetime(true)
    let second = sendCodeResponseData.interval_second
    const targetTime = serverDatetime + second

    const secondTimer = setInterval(async () => {
      try {
        if (second % 5 === 0) {
          second = await handleServerRemainSecond(targetTime)
        }
        if (second <= 0) {
          sendCodeBtnDisable.value = false
          sendCodeBtnText.value = '发送验证码'
          clearInterval(secondTimer)
          return
        }
        sendCodeBtnText.value = `${second.toString()}秒`
        second--
      } catch (e) {
        clearInterval(secondTimer)
      }
    }, 1000)
  } catch (e) {
    sendCodeLoading.value = false
  }
}

const handleSubmit = async (submitResult: any) => {
  submitResult.e.preventDefault()
  submitLoading.value = true
  if (submitResult.validateResult === true) {
    console.log('注册')
    try {
      const findPasswordResponseData = await http.postWithoutToken('oauth/find-password-finish', formData.value)
      storage.setToken({
        accessToken: findPasswordResponseData.access_token,
        accessTokenExpire: findPasswordResponseData.access_token_expire_in + tool.getTimestamp(),
        accessTokenExpireIn: findPasswordResponseData.access_token_expire_in,
        refreshToken: findPasswordResponseData.refresh_token,
        refreshTokenExpire: findPasswordResponseData.refresh_token_expire_in + tool.getTimestamp(),
        refreshTokenExpireIn: findPasswordResponseData.refresh_token_expire_in,
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
    } catch (error) {
      console.log(error)
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
})
</script>

<template>
  <div class="find-password-wrap">
    <t-card class="find-password-card" header-bordered>
      <template #content>
        <div class="find-password-card-content">
          <div class="find-password-logo-space">
            <img :src="logoPath" alt="ChatAIT" />
          </div>
          <div class="find-password-form-space">
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
                <t-input v-model="formData.username" class="input-username" placeholder="请输入邮箱">
                  <template #prefix-icon>
                    <user-icon />
                  </template>
                  <template #suffix>
                    <t-button :loading="sendCodeLoading" :disabled="sendCodeBtnDisable" @click="handleSendCode">
                      {{ sendCodeBtnText }}
                    </t-button>
                  </template>
                </t-input>
              </t-form-item>

              <t-form-item name="code">
                <t-input v-model="formData.code" type="text" clearable placeholder="请输入验证码">
                  <template #prefix-icon>
                    <lock-on-icon />
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

              <t-form-item name="confirm_password">
                <t-input v-model="formData.confirm_password" type="password" clearable placeholder="请重复密码">
                  <template #prefix-icon>
                    <lock-on-icon />
                  </template>
                </t-input>
              </t-form-item>

              <t-form-item class="find-password-btn-space">
                <t-space direction="vertical" align="center">
                  <t-button theme="primary" type="submit" block :loading="submitLoading">提交</t-button>
                  <div class="find-password-btn-login" @click="handleLogin">返回登录</div>
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
.find-password-wrap {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 100%;
  overflow: hidden;

  .find-password-card {
    flex: 0 0 auto;
    width: 400px;

    .find-password-card-content {
      display: flex;
      flex-direction: column;
      width: 100%;
      height: 100%;
      overflow: hidden;

      .find-password-logo-space {
        flex: 0 0 auto;
        display: flex;
        width: 100%;
        height: 80px;
        justify-content: center;
        align-items: center;
        overflow: hidden;
      }

      .find-password-btn-space .t-space {
        width: 100%;
      }

      .find-password-btn-login {
        width: 100%;
        text-align: center;
        cursor: pointer;
      }

      .input-username .t-input {
        padding-right: 0 !important;
      }
    }
  }

  @media screen and (max-width: 768px) {
    .find-password-card {
      width: 90%;
    }
  }
}
</style>
