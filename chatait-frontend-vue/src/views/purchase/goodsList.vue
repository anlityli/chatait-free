<!--
  - Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { CheckIcon } from 'tdesign-icons-vue-next'
import { useRoute, useRouter } from 'vue-router'
import http from '@/utils/network/http.js'
import { ResponseShopGoodsItem } from '@/utils/model/response/shop.js'
import tool from '../../utils/tool/tool'
import { ShopGoodsBuyTypeLevel } from '@/utils/constant/shop'

const route = useRoute()
const router = useRouter()

const goodsList = ref<ResponseShopGoodsItem[]>([])

const handleGetData = async () => {
  goodsList.value = (await http.get('shop/goods-list')) as ResponseShopGoodsItem[]
}

const handleToConfirm = async (goodsId: string) => {
  await router.push(`/purchase/confirm-order/${goodsId}`)
}

onMounted(() => {
  handleGetData()
})
</script>
<template>
  <div class="purchase-goods-list-wrap">
    <t-row :gutter="[20, 20]">
      <t-col v-for="(item, index) in goodsList" :key="index" :md="4" :sm="6" :xs="12">
        <t-card :title="item.title">
          <template #actions>
            <span
              >￥{{ tool.centToYuan(item.real_price) }}
              <template v-if="item.buy_type === ShopGoodsBuyTypeLevel">{{
                item.active_expire_type === 1
                  ? '/天'
                  : item.active_expire_type === 2
                  ? '/月'
                  : item.active_expire_type === 3
                  ? '/年'
                  : ''
              }}</template>
            </span>
          </template>
          <div class="goods-space">
            <div v-for="(featItem, featIndex) in item.feat_items_slice" :key="featIndex" class="goods-introduce-list">
              <div class="goods-introduce-icon">
                <check-icon></check-icon>
              </div>
              <div class="goods-introduce-text">
                <span>{{ featItem.text }}</span>
              </div>
            </div>
          </div>
          <div class="goods-tool">
            <t-divider></t-divider>
            <div class="goods-introduce-list">
              <t-button block :disabled="item.real_price <= 0" @click="handleToConfirm(item.id)">立即购买</t-button>
            </div>
          </div>
        </t-card>
      </t-col>
    </t-row>
  </div>
</template>

<style lang="scss">
.purchase-goods-list-wrap {
  .goods-space {
    width: 100%;
    height: 120px;

    .goods-introduce-list {
      display: flex;
      width: 100%;

      .goods-introduce-icon {
        flex: 0 0 auto;
        width: 30px;
        height: 30px;
        text-align: center;

        svg {
          width: 20px;
          height: 20px;
        }
      }

      .goods-introduce-text {
        flex: 1 1 auto;
        width: 100%;
      }
    }
  }
}
</style>
