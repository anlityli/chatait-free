<!--
  - Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<template>
  <t-row :gutter="[16, 16]">
    <t-col v-for="(item, index) in dashboardPanel" :key="item.title" :xs="6" :xl="3">
      <t-card
        :title="item.title"
        :bordered="false"
        :class="{ 'dashboard-item': true, 'dashboard-item--main-color': index === currentCard }"
        @mouseover="handleCardBackground(index)"
        @click="handleToPage(index)"
      >
        <div class="dashboard-item-top">
          <span :style="{ fontSize: `${resizeTime * 28}px` }">{{ item.number }}</span>
        </div>
        <div class="dashboard-item-left">
          <span v-if="index === 0" :style="{ marginTop: `-24px` }">
            <chart-icon />
          </span>
          <span v-else-if="index === 1" :style="{ marginTop: `-24px` }">
            <view-module-icon />
          </span>
          <span v-else-if="index === 2" :style="{ marginTop: `-24px` }">
            <usergroup-icon />
          </span>
          <span v-else :style="{ marginTop: '-24px' }">
            <chat-icon />
          </span>
        </div>
        <template #footer>
          <div class="dashboard-item-bottom">
            <div class="dashboard-item-block">
              上月
              <trend
                class="dashboard-item-trend"
                :type="item.number >= item.lastNumber ? 'up' : 'down'"
                :is-reverse-color="index === currentCard"
                :describe="item.lastNumber"
              />
            </div>
            <t-icon v-show="index === 1 || index === 2" name="chevron-right" />
          </div>
        </template>
      </t-card>
    </t-col>
  </t-row>
</template>

<script setup lang="ts">
import { BarChart, LineChart } from 'echarts/charts'
import * as echarts from 'echarts/core'
import { CanvasRenderer } from 'echarts/renderers'
import { ChartIcon, ChatIcon, UsergroupIcon, ViewModuleIcon } from 'tdesign-icons-vue-next'
import { onMounted, ref } from 'vue'

// 导入样式
import { useRouter } from 'vue-router'
import Trend from '@/components/trend/index.vue'

import http from '@/utils/network/http'
import tool from '@/utils/tool/tool'
import { DashboardPanel } from '@/pages/dashboard/model/model'

const router = useRouter()

echarts.use([LineChart, BarChart, CanvasRenderer])

const resizeTime = ref(1)

// dashboardCards
const dashboardPanel = ref<DashboardPanel[]>([
  {
    title: '本月收入',
    number: '¥ 0',
    lastNumber: '¥ 0',
    leftType: 'echarts-line',
  },
  {
    title: '本月订单',
    number: 0,
    lastNumber: 0,
    leftType: 'echarts-bar',
  },
  {
    title: '本月新增用户',
    number: 0,
    lastNumber: 0,
    leftType: 'icon-usergroup',
  },
  {
    title: '本月提问次数',
    number: 0,
    lastNumber: 0,
    leftType: 'icon-file-paste',
  },
])

const currentCard = ref(0)
const handleCardBackground = (index: number) => {
  currentCard.value = index
}

const handleToPage = (index: number) => {
  switch (index) {
    case 0:
      break
    case 1:
      router.push('/shop/order-list')
      break
    case 2:
      router.push('/user/list')
      break
    case 3:
      break
    default:
  }
}

onMounted(async () => {
  // 获取统计数据
  const [userStatisticData, amountStatisticData, orderStatisticData, conversationStatisticData] = await Promise.all([
    http.get('dashboard/user-statistic', { type: 'monthly' }),
    http.get('dashboard/amount-statistic', { type: 'monthly' }),
    http.get('dashboard/order-statistic', { type: 'monthly' }),
    http.get('dashboard/conversation-statistic', { type: 'monthly' }),
  ])
  dashboardPanel.value[0].number = tool.centToYuan(amountStatisticData.this_count)
  dashboardPanel.value[0].lastNumber = tool.centToYuan(amountStatisticData.last_count)
  dashboardPanel.value[1].number = orderStatisticData.this_count
  dashboardPanel.value[1].lastNumber = orderStatisticData.last_count
  dashboardPanel.value[2].number = userStatisticData.this_count
  dashboardPanel.value[2].lastNumber = userStatisticData.last_count
  dashboardPanel.value[3].number = conversationStatisticData.this_count
  dashboardPanel.value[3].lastNumber = conversationStatisticData.last_count
})
</script>

<script lang="ts">
export default {
  name: 'DashboardBase',
}
</script>

<style lang="less" scoped>
.dashboard-item {
  padding: var(--td-comp-paddingTB-xl) var(--td-comp-paddingLR-xxl);

  :deep(.t-card__header) {
    padding: 0;
  }

  :deep(.t-card__footer) {
    padding: 0;
  }

  :deep(.t-card__title) {
    font: var(--td-font-body-medium);
    color: var(--td-text-color-secondary);
  }

  :deep(.t-card__body) {
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    flex: 1;
    position: relative;
    padding: 0;
    margin-top: var(--td-comp-margin-s);
    margin-bottom: var(--td-comp-margin-xxl);
  }

  &:hover {
    cursor: pointer;
  }

  &-top {
    display: flex;
    flex-direction: row;
    align-items: flex-start;

    > span {
      display: inline-block;
      color: var(--td-text-color-primary);
      font-size: var(--td-font-size-headline-medium);
      line-height: var(--td-line-height-headline-medium);
    }
  }

  &-bottom {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;

    > .t-icon {
      cursor: pointer;
      font-size: var(--td-comp-size-xxxs);
    }
  }

  &-block {
    display: flex;
    align-items: center;
    justify-content: center;
    color: var(--td-text-color-placeholder);
  }

  &-trend {
    margin-left: var(--td-comp-margin-s);
  }

  &-left {
    position: absolute;
    top: 0;
    right: 0;

    > span {
      display: inline-flex;
      align-items: center;
      justify-content: center;
      width: var(--td-comp-size-xxxl);
      height: var(--td-comp-size-xxxl);
      background: var(--td-brand-color-light);
      border-radius: 50%;

      .t-icon {
        font-size: 24px;
        color: var(--td-brand-color);
      }
    }
  }

  // 针对第一个卡片需要反色处理
  &--main-color {
    background: var(--td-brand-color);
    color: var(--td-text-color-primary);

    :deep(.t-card__title),
    .dashboard-item-top span,
    .dashboard-item-bottom {
      color: var(--td-text-color-anti);
    }

    .dashboard-item-block {
      color: var(--td-text-color-anti);
      opacity: 0.6;
    }

    .dashboard-item-bottom {
      color: var(--td-text-color-anti);
    }
  }
}
</style>
