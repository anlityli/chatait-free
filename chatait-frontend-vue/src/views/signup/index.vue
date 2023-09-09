<!--
  - Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { LockOnIcon, UserIcon } from 'tdesign-icons-vue-next'
import { useRouter } from 'vue-router'
import { SignupForm } from '@/views/signup/script/model'
import http from '@/utils/network/http'
import storage from '@/utils/storage/storage'
import tool from '@/utils/tool/tool'
import { ResponseUser } from '@/utils/model/response/user'
import logoPath from '@/assets/image/logo.png'

const router = useRouter()

const formData = ref<SignupForm>({
  username: '',
  password: '',
  confirm_password: '',
  code: '',
  nickname: '',
})

const submitLoading = ref(false)
const sendCodeLoading = ref(false)
const sendCodeBtnText = ref('发送验证码')
const sendCodeBtnDisable = ref(false)
const emailCodeEnable = ref(false)

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
  nickname: [{ required: true, message: '昵称必填', type: 'error' }],
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
    const sendCodeResponseData = await http.postWithoutToken('oauth/signup-send-code', formData.value)
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
      const signupResponseData = await http.postWithoutToken('oauth/signup-finish', formData.value)
      storage.setToken({
        accessToken: signupResponseData.access_token,
        accessTokenExpire: signupResponseData.access_token_expire_in + tool.getTimestamp(),
        accessTokenExpireIn: signupResponseData.access_token_expire_in,
        refreshToken: signupResponseData.refresh_token,
        refreshTokenExpire: signupResponseData.refresh_token_expire_in + tool.getTimestamp(),
        refreshTokenExpireIn: signupResponseData.refresh_token_expire_in,
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
  http.getWithoutToken('config/options').then((response) => {
    emailCodeEnable.value = response.emailCodeEnable === '1'
  })
})
</script>

<template>
  <div class="signup-wrap">
    <t-card class="signup-card" header-bordered>
      <template #content>
        <div class="signup-card-content">
          <div class="signup-logo-space">
            <img :src="logoPath" alt="ChatAIT" />
          </div>
          <div class="signup-form-space">
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
                  <template v-if="emailCodeEnable" #suffix>
                    <t-button :loading="sendCodeLoading" :disabled="sendCodeBtnDisable" @click="handleSendCode">
                      {{ sendCodeBtnText }}
                    </t-button>
                  </template>
                </t-input>
              </t-form-item>

              <t-form-item v-if="emailCodeEnable" name="code">
                <t-input v-model="formData.code" type="text" clearable placeholder="请输入验证码">
                  <template #prefix-icon>
                    <lock-on-icon />
                  </template>
                </t-input>
              </t-form-item>

              <t-form-item name="nickname">
                <t-input v-model="formData.nickname" type="text" clearable placeholder="请输入昵称">
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

              <t-form-item class="signup-btn-space">
                <t-space direction="vertical" align="center">
                  <t-button theme="primary" type="submit" block :loading="submitLoading">注册</t-button>
                  <div class="signup-btn-login" @click="handleLogin">返回登录</div>
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
.signup-wrap {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 100%;
  overflow: hidden;

  .signup-card {
    flex: 0 0 auto;
    width: 400px;

    .signup-card-content {
      display: flex;
      flex-direction: column;
      width: 100%;
      height: 100%;
      overflow: hidden;

      .signup-logo-space {
        flex: 0 0 auto;
        display: flex;
        width: 100%;
        height: 80px;
        justify-content: center;
        align-items: center;
        overflow: hidden;
      }

      .signup-btn-space .t-space {
        width: 100%;
      }

      .signup-btn-login {
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
    .signup-card {
      width: 90%;
    }
  }
}
</style>
