<!--
  - Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<template>
  <div class="main-content-container page-goods-edit">
    <t-form class="edit-form" :data="formData" :colon="true" @submit="handleSubmit">
      <t-tabs :default-value="'info'">
        <t-tab-panel :value="'info'" label="基本信息" style="padding: 20px">
          <t-form-item label="标题" name="title">
            <t-input v-model="formData.title" class="edit-input"></t-input>
          </t-form-item>
          <t-form-item label="市场价" name="market_price_yuan">
            <t-input v-model="formData.market_price_yuan" class="edit-input">
              <template #suffix>
                {{ '元' }}
              </template>
            </t-input>
          </t-form-item>
          <t-form-item label="实际价" name="real_price_yuan">
            <t-input v-model="formData.real_price_yuan" class="edit-input">
              <template #suffix>
                {{ '元' }}
              </template>
            </t-input>
          </t-form-item>
          <t-form-item label="上架" name="status">
            <t-switch v-model="formData.status" :custom-value="[1, 0]" />
          </t-form-item>
        </t-tab-panel>
        <t-tab-panel :value="'featItems'" label="特色功能介绍" style="padding: 20px">
          <div class="feat-items-list">
            <div v-for="(item, index) in formData.feat_items" :key="index" class="feat-item">
              <div class="feat-item-icon">
                <t-select
                  v-model="item.icon"
                  placeholder="请选择"
                  :popup-props="{ overlayInnerStyle: { width: '400px' } }"
                  @change="handleFeatItems"
                >
                  <t-option
                    v-for="iconItem in iconOptions"
                    :key="iconItem.stem"
                    :value="iconItem.stem"
                    class="overlay-options"
                  >
                    <div>
                      <t-icon :name="iconItem.stem"></t-icon>
                    </div>
                  </t-option>
                  <template #valueDisplay>
                    <t-icon :name="item.icon" :style="{ marginRight: '8px' }"></t-icon>
                  </template>
                </t-select>
              </div>
              <div class="feat-item-text">
                <t-input v-model="item.text" @change="handleFeatItems"></t-input>
              </div>
              <div class="feat-item-tool">
                <t-button
                  shape="circle"
                  theme="danger"
                  :disabled="item.icon === '' && item.text === ''"
                  @click="handleDeleteFeatItem(index)"
                >
                  <template #icon>
                    <delete-icon></delete-icon>
                  </template>
                </t-button>
              </div>
            </div>
          </div>
        </t-tab-panel>
        <t-tab-panel :value="'other'" label="购买开通" style="padding: 20px">
          <t-form-item label="购买类型" name="buy_type">
            <t-select v-model="formData.buy_type" class="edit-input">
              <t-option v-for="(item, index) in allBuyType" :key="index" :label="item.label" :value="item.value" />
            </t-select>
          </t-form-item>
          <t-form-item v-if="formData.buy_type === 1" label="购买级别" name="active_level_id">
            <t-select v-model="formData.active_level_id" class="edit-input">
              <t-option v-for="(item, index) in allLevel" :key="index" :label="item.level_name" :value="item.id" />
            </t-select>
          </t-form-item>
          <t-form-item
            v-if="formData.buy_type === 1 && formData.active_level_id > 1"
            label="有效期类型"
            name="active_expire_type"
          >
            <t-select v-model="formData.active_expire_type" class="edit-input">
              <t-option
                v-for="(item, index) in allActiveExpireType"
                :key="index"
                :label="item.label"
                :value="item.value"
              />
            </t-select>
          </t-form-item>
          <t-form-item
            v-if="formData.buy_type === 1 && formData.active_level_id > 1"
            label="有效期"
            name="active_expire_value"
          >
            <t-input-number v-model="formData.active_expire_value" theme="normal" class="edit-input">
              <template #suffix>
                {{ formData.active_expire_type === 1 ? '天' : formData.active_expire_type === 2 ? '月' : '年' }}
              </template>
            </t-input-number>
          </t-form-item>
          <t-form-item
            v-if="formData.buy_type !== ShopGoodsBuyTypeLevel"
            :label="`购买${
              formData.buy_type === ShopGoodsBuyTypeBalance
                ? configStore.walletName(WalletType.balance)
                : formData.buy_type === ShopGoodsBuyTypeGpt3
                ? configStore.walletName(WalletType.gpt3)
                : formData.buy_type === ShopGoodsBuyTypeGpt4
                ? configStore.walletName(WalletType.gpt4)
                : configStore.walletName(WalletType.midjourney)
            }`"
            name="buy_value_yuan"
          >
            <t-input v-model="formData.buy_value_yuan" class="edit-input">
              <template #suffix>
                {{ '元' }}
              </template>
            </t-input>
          </t-form-item>
        </t-tab-panel>
      </t-tabs>
      <t-divider></t-divider>
      <t-button theme="primary" type="submit">提交</t-button>
    </t-form>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { MessagePlugin } from 'tdesign-vue-next'
import { useRoute, useRouter } from 'vue-router'
import { DeleteIcon, manifest } from 'tdesign-icons-vue-next'
import http from '@/utils/network/http'
import { FormShopGoodsEdit, FormShopGoodsEditFeatItem } from '@/pages/shop/model/model'
import { ResponseShopGoods } from '@/utils/model/response/shop'
import tool from '@/utils/tool/tool'
import { ResponseConfigLevelListItem } from '@/utils/model/response/config'
import { useConfigStore } from '@/store'
import { WalletType } from '@/constants/config'
import {
  ShopGoodsBuyTypeBalance,
  ShopGoodsBuyTypeGpt3,
  ShopGoodsBuyTypeGpt4,
  ShopGoodsBuyTypeLevel,
  ShopGoodsBuyTypeMidjourney,
} from '@/constants/shop'
// 获取全部图标的列表
const iconOptions = ref(manifest)

const router = useRouter()
const route = useRoute()
const configStore = useConfigStore()

const formData = ref<FormShopGoodsEdit>({
  id: '',
  title: '',
  content: '',
  feat_items: [],
  buy_type: 1,
  active_level_id: 1,
  active_expire_type: 1,
  active_expire_value: 0,
  buy_value: 0,
  buy_value_yuan: '0.00',
  market_price: 0,
  market_price_yuan: '0.00',
  real_price: 0,
  real_price_yuan: '0.00',
  status: 0,
  sort: 0,
})

const allBuyType = ref([
  { label: '开通级别', value: ShopGoodsBuyTypeLevel },
  { label: `充值${configStore.walletName(WalletType.balance)}`, value: ShopGoodsBuyTypeBalance },
  // { label: `购买${configStore.walletName(WalletType.gpt3)}`, value: ShopGoodsBuyTypeGpt3 },
  // { label: `购买${configStore.walletName(WalletType.gpt4)}`, value: ShopGoodsBuyTypeGpt4 },
  // { label: `购买${configStore.walletName(WalletType.midjourney)}`, value: ShopGoodsBuyTypeMidjourney },
])

const allActiveExpireType = ref([
  { label: '按天', value: 1 },
  { label: '按月', value: 2 },
  { label: '按年', value: 3 },
])

const allLevel = ref<ResponseConfigLevelListItem[]>([])

const handleFeatItems = () => {
  for (let i = 0; i < formData.value.feat_items.length; i++) {
    if (formData.value.feat_items[i].icon === '' && formData.value.feat_items[i].text === '') {
      formData.value.feat_items.splice(i, 1)
    }
  }
  if (
    formData.value.feat_items.length <= 0 ||
    formData.value.feat_items[formData.value.feat_items.length - 1].icon !== '' ||
    formData.value.feat_items[formData.value.feat_items.length - 1].text !== ''
  ) {
    formData.value.feat_items.push({
      icon: '',
      text: '',
    })
  }
}

const handleDeleteFeatItem = (index: number) => {
  formData.value.feat_items.splice(index, 1)
  handleFeatItems()
}

const handleSubmit = async () => {
  try {
    const requestData = {
      id: formData.value.id,
      title: formData.value.title,
      content: formData.value.content,
      feat_items: JSON.stringify(formData.value.feat_items),
      buy_type: formData.value.buy_type,
      active_level_id: formData.value.active_level_id,
      active_expire_type: formData.value.active_expire_type,
      active_expire_value: formData.value.active_expire_value,
      buy_value: tool.yuanToCent(formData.value.buy_value_yuan),
      market_price: tool.yuanToCent(formData.value.market_price_yuan),
      real_price: tool.yuanToCent(formData.value.real_price_yuan),
      status: formData.value.status,
      sort: formData.value.sort,
    }
    if (route.name === 'shopGoodsEdit') {
      await http.post('shop/goods-edit', requestData)
    } else {
      await http.post('shop/goods-add', requestData)
    }
    await MessagePlugin.success('操作成功')
    await router.back()
  } catch (e) {
    await MessagePlugin.success(`发生错误: ${e.toString()}`)
  }
}

onMounted(async () => {
  allLevel.value = (await http.get(`config/level-list`)) as ResponseConfigLevelListItem[]
  if (route.name === 'shopGoodsEdit') {
    const goodsData = (await http.get(`shop/goods-one?id=${route.query.id}`)) as ResponseShopGoods
    formData.value.id = route.query.id as string
    formData.value.title = goodsData.title
    formData.value.content = goodsData.content
    formData.value.buy_type = goodsData.buy_type
    formData.value.active_level_id = goodsData.active_level_id
    formData.value.active_expire_type = goodsData.active_expire_type
    formData.value.active_expire_value = goodsData.active_expire_value
    formData.value.buy_value = goodsData.buy_value
    formData.value.buy_value_yuan = tool.centToYuan(goodsData.buy_value)
    formData.value.market_price = goodsData.market_price
    formData.value.market_price_yuan = tool.centToYuan(goodsData.market_price)
    formData.value.real_price = goodsData.real_price
    formData.value.real_price_yuan = tool.centToYuan(goodsData.real_price)
    formData.value.status = goodsData.status
    formData.value.sort = goodsData.sort
    if (goodsData.feat_items !== '') {
      formData.value.feat_items = JSON.parse(goodsData.feat_items) as FormShopGoodsEditFeatItem[]
    }
  }
  handleFeatItems()
})
</script>
<style>
.overlay-options {
  display: inline-block;
  font-size: 20px;
}
</style>
<style lang="less" scoped>
.feat-items-list {
  width: 100%;

  .feat-item {
    display: flex;
    width: 100%;
    height: 60px;

    .feat-item-icon {
      flex: 0 0 auto;
      width: 80px;
      display: flex;
    }

    .feat-item-text {
      flex: 1 1 auto;
      width: 100%;
      padding: 0 10px;
      box-sizing: border-box;
    }

    .feat-item-tool {
      flex: 0 0 auto;
      width: 60px;
    }
  }
}
</style>
