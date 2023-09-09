<!--
  - Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { EditIcon } from 'tdesign-icons-vue-next'
import http from '@/utils/network/http'
import { ResponseUser } from '@/utils/model/response/user'
import tool from '@/utils/tool/tool'
import storage from '@/utils/storage/storage'
import { ResponseFinanceWalletInfo } from '@/utils/model/response/finance'
import { WalletType } from '@/utils/constant/config'
import { useConfigStore } from '@/store'
import AvatarImage from '@/components/avatarImage/AvatarImage.vue'

const configStore = useConfigStore()

const infoData = ref<ResponseUser>({
  id: '',
  username: '',
  nickname: '',
  avatar: '',
  level_id: 0,
  level_name: '',
  level_expire_date: '',
  created_at: 0,
  last_login_at: 0,
})

const walletInfo = ref<ResponseFinanceWalletInfo>({
  user_id: '',
  balance: 0,
  gpt3: 0,
  gpt4: 0,
  midjourney: 0,
})

const userInfoData = storage.getUserInfo()
const isEditNickname = ref(false)
const editNicknameForm = ref({
  nickname: '',
})

const handleGetData = async () => {
  infoData.value = (await http.get('user/info')) as ResponseUser
  walletInfo.value = (await http.get('finance/wallet-info')) as ResponseFinanceWalletInfo
}

const handleShowEditNickname = () => {
  isEditNickname.value = true
  editNicknameForm.value.nickname = infoData.value.nickname
}

const handleSubmitEditNickname = async () => {
  try {
    await http.post('user/edit-nickname', editNicknameForm.value)
    infoData.value.nickname = editNicknameForm.value.nickname
    isEditNickname.value = false
    editNicknameForm.value.nickname = ''
  } catch (e) {
    console.log(e)
  }
}

onMounted(() => {
  handleGetData()
})
</script>
<template>
  <div class="user-profile-wrap">
    <t-card title="个人信息">
      <div class="info-avatar">
        <t-avatar size="60px">
          <template #content>
            <avatar-image :url="userInfoData?.avatar"></avatar-image>
          </template>
        </t-avatar>
      </div>
      <div class="info-block">
        <div class="info-item">
          <div class="info-item-label">用户名:</div>
          <div class="info-item-value">{{ userInfoData?.username }}</div>
        </div>
        <div class="info-item">
          <div class="info-item-label">昵称:</div>
          <div class="info-item-value">
            <template v-if="!isEditNickname">{{ infoData.nickname }}</template>
            <template v-else>
              <t-input
                v-model="editNicknameForm.nickname"
                type="text"
                clearable
                placeholder="请输入昵称"
                @blur="handleSubmitEditNickname"
              ></t-input>
            </template>
            <edit-icon style="margin-left: 10px" @click="handleShowEditNickname"></edit-icon>
          </div>
        </div>
      </div>
      <div class="info-block">
        <div class="info-item">
          <div class="info-item-label">加入时间:</div>
          <div class="info-item-value">{{ tool.formatDate(infoData.created_at) }}</div>
        </div>
        <div class="info-item">
          <div class="info-item-label">最后登陆时间:</div>
          <div class="info-item-value">{{ tool.formatDate(infoData.last_login_at) }}</div>
        </div>
      </div>
      <div class="info-block">
        <div class="info-item">
          <div class="info-item-label">级别:</div>
          <div class="info-item-value">{{ infoData.level_name }}</div>
        </div>
        <div v-if="infoData.level_id > 1" class="info-item">
          <div class="info-item-label">会员到期:</div>
          <div class="info-item-value">{{ infoData.level_expire_date }}</div>
        </div>
      </div>
      <div class="info-block">
        <div class="info-item">
          <div class="info-item-label">剩余{{ configStore.walletName(WalletType.balance) }}:</div>
          <div class="info-item-value">{{ tool.centToYuan(walletInfo.balance) }}</div>
        </div>
        <div class="info-item">
          <div class="info-item-label">剩余{{ configStore.walletName(WalletType.gpt3) }}:</div>
          <div class="info-item-value">{{ tool.centToYuan(walletInfo.gpt3) }}</div>
        </div>
      </div>
      <div class="info-block">
        <div class="info-item">
          <div class="info-item-label">剩余{{ configStore.walletName(WalletType.gpt4) }}:</div>
          <div class="info-item-value">{{ tool.centToYuan(walletInfo.gpt4) }}</div>
        </div>
        <div class="info-item">
          <div class="info-item-label">剩余{{ configStore.walletName(WalletType.midjourney) }}:</div>
          <div class="info-item-value">{{ tool.centToYuan(walletInfo.midjourney) }}</div>
        </div>
      </div>
    </t-card>
  </div>
</template>

<style lang="scss" scoped>
.user-profile-wrap {
  box-sizing: border-box;
  padding: 10px;

  .info-block {
    margin-bottom: 10px;
  }

  .info-avatar {
    text-align: center;
    padding: 15px;
  }
}
</style>
