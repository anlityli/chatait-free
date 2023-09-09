<!--
  - Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import QArt from 'qartjs'
import { MessagePlugin } from 'tdesign-vue-next'
import { CheckCircleIcon } from 'tdesign-icons-vue-next'
import http from '@/utils/network/http'
import { ResponseShopOrder } from '@/utils/model/response/shop'
import tool from '@/utils/tool/tool'
import iconPath from '@/assets/image/icon_back.png'
import { ResponseConfigPayListItem } from '@/utils/model/response/config'

const route = useRoute()
const router = useRouter()

const configPayList = ref<ResponseConfigPayListItem[]>([])
const orderDetail = ref<ResponseShopOrder | null>(null)
const payment = ref<any>(null)
const timerText = ref('')

const payAmount = computed((): string => {
  if (payment.value !== null) {
    return tool.centToYuan(payment.value.pay_amount)
  }
  return '0'
})

const handleGetData = async () => {
  configPayList.value = (await http.get('config/pay-list')) as ResponseConfigPayListItem[]
  orderDetail.value = (await http.get('shop/order-detail', {
    order_id: route.params.orderId,
  })) as ResponseShopOrder
}

const handleServerRemainSecond = async (targetTime: number) => {
  const serverTime = await tool.serverDatetime(true)
  return targetTime - serverTime
}

const handlePay = async (configPayId: number, channel: string) => {
  payment.value = await http.post('shop/pay-order', {
    order_id: route.params.orderId,
    config_pay_id: configPayId,
    pay_channel: channel,
  })
  // 二维码
  const qartParams = {
    value: payment.value.pay_url,
    imagePath: iconPath,
    filter: 'threshold',
    size: 300,
  }
  const qart = new QArt(qartParams)
  qart.make(document.getElementById('qr'))
  // 倒计时
  let timeout = payment.value.timeout * 60
  timerText.value = tool.secondToMinute(timeout)
  const timer = setInterval(async () => {
    try {
      if (timeout % 5 === 0) {
        await handleGetData()
        if (orderDetail.value?.status !== 0) {
          clearInterval(timer)
          await router.push(`/purchase/order-detail/${orderDetail.value?.id}`)
        }
        if (payment.value !== null) {
          timeout = await handleServerRemainSecond(payment.value.due_expire_at)
          console.log('计算服务器倒计时')
        }
      }
      if (timeout <= 0) {
        timerText.value = ''
        payment.value = null
        clearInterval(timer)
        return
      }
      timeout--
      timerText.value = tool.secondToMinute(timeout)
    } catch (e) {
      console.log(e)
      await MessagePlugin.error('倒计时服务错误')
      payment.value = null
      clearInterval(timer)
    }
  }, 1000)
}

onMounted(async () => {
  await handleGetData()
})
</script>
<template>
  <div class="purchase-pay-order-wrap">
    <t-card>
      <div class="pay-order-submit">
        <div class="pay-order-amount">
          <template v-if="payment !== null">
            {{ payAmount }}
          </template>
          <template v-else-if="orderDetail !== null">
            {{ tool.centToYuan(orderDetail?.order_amount) }}
          </template>
        </div>
        <div class="pay-order-sn">SN: {{ orderDetail?.order_sn }}</div>
        <div v-if="payment === null && orderDetail?.status === 0" class="pay-order-pay-list">
          <div v-for="(item, index) in configPayList" :key="index" class="pay-order-type">
            <div v-for="(channelItem, channelIndex) in item.pay_channel" :key="channelIndex" class="pay-order-btn">
              <t-button
                block
                :theme="
                  channelItem.channel_name.indexOf('微信') !== -1
                    ? 'success'
                    : channelItem.channel_name.indexOf('支付宝') !== -1
                    ? 'primary'
                    : 'warning'
                "
                @click="handlePay(item.id, channelItem.channel)"
              >
                {{ channelItem.channel_name }}{{ configPayList.length > 1 ? `(${item.api_name})` : '' }}
              </t-button>
            </div>
          </div>
        </div>
        <div v-show="payment !== null && orderDetail?.status === 0" class="pay-order-timer">
          请在 <span style="color: #d54941">{{ timerText }}</span> 内支付
        </div>
        <div v-show="payment !== null && orderDetail?.status === 0" class="pay-order-qr">
          <div id="qr"></div>
        </div>
        <div v-show="orderDetail !== null && orderDetail?.status !== 0" class="pay-order-status">
          <div style="text-align: center">
            <check-circle-icon size="50" style="color: green"></check-circle-icon>
            <div>支付已完成</div>
          </div>
        </div>
      </div>
    </t-card>
  </div>
</template>

<style lang="scss">
.purchase-pay-order-wrap {
  .pay-order-amount {
    width: 100%;
    text-align: center;
    font-size: 26px;
    font-weight: bold;
    color: #d54941;
    padding: 20px 0 0 0;
    box-sizing: border-box;
  }

  .pay-order-sn {
    width: 100%;
    text-align: center;
    font-size: 12px;
    color: var(--td-text-color-secondary);
    padding: 0 0 20px 0;
    box-sizing: border-box;
  }

  .pay-order-type {
    width: 100%;
    display: flex;

    .pay-order-btn {
      flex: 1 1 auto;
      width: 100%;
      padding: 10px;
      box-sizing: border-box;
    }
  }

  .pay-order-timer {
    text-align: center;
  }

  .pay-order-qr {
    width: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .pay-order-status {
    width: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
  }
}
</style>
