<!--
  - Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import {
  CartIcon,
  DeleteIcon,
  EditIcon,
  LogoutIcon,
  RootListIcon,
  UserIcon,
  ImageIcon,
  ChatIcon,
} from 'tdesign-icons-vue-next'
import { useRoute, useRouter } from 'vue-router'
import { getTopicStore } from '@/store'
import http from '@/utils/network/http'
import storage from '@/utils/storage/storage'
import logoPath from '@/assets/image/logo.png'
import { TopicTypeMidjourney, TopicTypeOpenaiGPT3, TopicTypeOpenaiGPT4 } from '@/utils/constant/topic'

const props = defineProps({
  isDrawer: Boolean,
})
const topicStore = getTopicStore()
const route = useRoute()
const router = useRouter()
const tMenuEle = ref<any | null>(null)

const emit = defineEmits(['clickItem'])

const handleClickItem = (type: string, index: number, item: any) => {
  emit('clickItem', type, index, item)
}

const heightLightValue = computed((): string => {
  return route.params.topicId as string
})

const handleTopicDel = async (index: number, topicId: string) => {
  await http.post(`conversation/topic-del`, { topic_id: topicId })
  topicStore.del(index)
}

const handleTopicShowEdit = (index: number) => {
  topicStore.topicList[index].edit_status = !topicStore.topicList[index].edit_status
}

const handleTopicEdit = async (event: any, index: number, topicId: string, title: string) => {
  await http.post(`conversation/topic-edit`, { topic_id: topicId, title })
  topicStore.edit(index, title)
  event.target.blur()
  topicStore.topicList[index].edit_status = false
}

const handlePurchase = () => {
  router.push('/purchase/goods-list')
  handleClickItem('tool', 0, '/purchase/goods-list')
}

const handleUser = () => {
  router.push('/user/profile')
  handleClickItem('tool', 1, '/user/profile')
}

const handleLogout = () => {
  storage.clearToken()
  storage.clearUserInfo()
  router.push('/login')
}

onMounted(async () => {
  if (!props.isDrawer) {
    // 首先初始化菜单防止有之前遗留菜单数据
    topicStore.init()
    // 获取数据
    await topicStore.getList()
  }

  // 监听滚动条
  if (tMenuEle.value !== null) {
    // console.log(tMenuEle.value.$el.firstElementChild.children[1])
    const menuUlEle = tMenuEle.value.$el.firstElementChild.children[1]
    menuUlEle.addEventListener('scroll', async (event: any) => {
      if (event.target.scrollTop >= event.target.scrollHeight - event.target.clientHeight) {
        // 加载新的列表数据
        await topicStore.getList()
      }
    })
  }
})
</script>

<template>
  <div class="menu">
    <t-menu ref="tMenuEle" :value="heightLightValue">
      <template #logo>
        <img height="28" :src="logoPath" alt="ChatAIT" />
      </template>
      <t-menu-item value="0" :to="`/conversation/0`" @click="handleClickItem('conversation', 0, null)">
        <template #icon>
          <edit-icon />
        </template>
        <span>开始一个新的话题</span>
      </t-menu-item>
      <t-menu-item
        v-for="(item, index) in topicStore.topicList"
        :key="index"
        :class="`menu-list-${
          item.type === TopicTypeOpenaiGPT3 ? 'gpt3' : item.type === TopicTypeOpenaiGPT4 ? 'gpt4' : 'mj'
        }`"
        :value="item.id"
        :to="`/conversation/${item.id}`"
        @mouseover="item.tool_show_status = true"
        @mouseleave="item.tool_show_status = false"
        @click="handleClickItem('conversation', index, item)"
      >
        <template #icon>
          <chat-icon v-if="item.type === TopicTypeOpenaiGPT3 || item.type === TopicTypeOpenaiGPT4" />
          <image-icon v-else-if="item.type === TopicTypeMidjourney" />
          <root-list-icon v-else />
        </template>
        <template #content>
          <div class="menu-item">
            <div class="menu-title">
              <t-input
                v-if="item.edit_status"
                v-model="item.title"
                autofocus
                @enter="
                  (value: any, context: any) => {
                    handleTopicEdit(context.e, index, item.id, item.title)
                  }
                "
                @blur="
                  (value: any, context: any) => {
                    handleTopicEdit(context.e, index, item.id, item.title)
                  }
                "
              ></t-input>
              <span v-else>{{ item.title }}</span>
            </div>
            <div v-if="item.tool_show_status" class="menu-tool">
              <div class="menu-edit">
                <edit-icon @click="handleTopicShowEdit(index)"></edit-icon>
              </div>
              <div class="menu-del">
                <delete-icon @click="handleTopicDel(index, item.id)"></delete-icon>
              </div>
            </div>
          </div>
        </template>
      </t-menu-item>
      <template #operations>
        <t-list class="menu-operations">
          <t-list-item
            :class="route.path.indexOf('/purchase') !== -1 ? 'operations-active' : ''"
            @click="handlePurchase"
          >
            <template #content>
              <div class="menu-operations-item-content">
                <div class="menu-operations-item-icon">
                  <cart-icon></cart-icon>
                </div>
                <div class="menu-operations-item-label">购买计划</div>
              </div>
            </template>
          </t-list-item>
          <t-list-item :class="route.path.indexOf('/user') !== -1 ? 'operations-active' : ''" @click="handleUser">
            <template #content>
              <div class="menu-operations-item-content">
                <div class="menu-operations-item-icon">
                  <user-icon></user-icon>
                </div>
                <div class="menu-operations-item-label">个人信息</div>
              </div>
            </template>
          </t-list-item>
          <t-list-item @click="handleLogout">
            <template #content>
              <div class="menu-operations-item-content">
                <div class="menu-operations-item-icon">
                  <logout-icon></logout-icon>
                </div>
                <div class="menu-operations-item-label">退出登录</div>
              </div>
            </template>
          </t-list-item>
        </t-list>
      </template>
    </t-menu>
  </div>
</template>

<style lang="scss">
.menu {
  height: 100%;
  overflow: hidden;

  .t-menu__logo {
    height: 48px !important;
    box-sizing: border-box;
    padding-left: 26px;

    h1 {
      font-size: 22px;
    }
  }

  .t-menu__item > svg {
    flex: 0 0 auto !important;
  }

  .t-menu__item > .t-menu__content {
    flex: 1 1 auto !important;

    .menu-tool svg {
      width: 14px !important;
      height: 14px !important;
      overflow: hidden;
    }
  }

  .t-menu__item {
    position: relative;
  }

  .t-menu__item.menu-list-gpt3:after {
    position: absolute;
    content: 'G3';
    display: block;
    top: 0;
    left: 21px;
    font-size: 7px;
  }

  .t-menu__item.menu-list-gpt4:after {
    position: absolute;
    content: 'G4';
    display: block;
    top: 0;
    left: 21px;
    font-size: 7px;
  }

  .menu-item {
    display: flex;
    width: 100%;
    height: 100%;
    overflow: hidden;

    .menu-title {
      flex: 1 1 auto;
      width: 100%;
      height: 100%;
      padding-right: 10px;
      text-overflow: ellipsis;
      white-space: nowrap;
      box-sizing: border-box;
      overflow: hidden;
    }

    .menu-tool {
      display: flex;
      flex: 0 0 auto;
      width: 40px;
      height: 100%;
      overflow: hidden;

      .menu-edit {
        flex: 1 1 auto;
        width: 100%;
        height: 100%;
        box-sizing: border-box;
      }

      .menu-del {
        flex: 1 1 auto;
        width: 100%;
        height: 100%;
        box-sizing: border-box;
      }
    }
  }

  .t-menu__operations {
    padding: 8px !important;
  }

  .menu-operations {
    .t-list-item {
      color: var(--td-text-color-secondary);
      padding: 7px !important;
      padding-left: 16px !important;
      margin: 4px;
      box-sizing: border-box;
      cursor: pointer;
    }

    .menu-operations-item-content {
      display: flex;
      width: 100%;
      height: 100%;
      overflow: hidden;

      .menu-operations-item-icon {
        flex: 0 0 auto;
        width: 25px;
        height: 100%;
        overflow: hidden;

        svg {
          width: 18px !important;
          height: 18px !important;
        }
      }

      .menu-operations-item-label {
        flex: 1 1 auto;
        width: 100%;
        height: 100%;
        overflow: hidden;
      }
    }

    .operations-active {
      color: var(--td-brand-color);
      background-color: var(--td-brand-color-light);
      border-radius: 5px;
    }

    .t-list-item:hover {
      background-color: var(--td-bg-color-container-hover);
    }
  }
}
</style>
