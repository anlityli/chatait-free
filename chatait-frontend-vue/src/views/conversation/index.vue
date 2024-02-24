<!--
  - Copyright 2024 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<script setup lang="ts">
import { computed, h, onBeforeUnmount, onMounted, reactive, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { marked } from 'marked'
import hljs from 'highlight.js'
import { EnterIcon, IconFont, LoadingIcon } from 'tdesign-icons-vue-next'
import { MessagePlugin } from 'tdesign-vue-next'
import http from '@/utils/network/http'
import { ResponsePage } from '@/utils/model/response/page'
import {
  ResponseConversationMidjourneySpeak,
  ResponseConversationSpeak,
  ResponseConversationSpeakItem,
  ResponseConversationStreamUuid,
  ResponseConversationTopic,
} from '@/utils/model/response/conversation'
import 'github-markdown-css/github-markdown-light.css'
// import 'highlight.js/styles/default.css'
import 'highlight.js/styles/atom-one-light.css'
// import 'highlight.js/styles/dark.css'
import { getTopicStore, getAppStore } from '@/store'
import storage from '@/utils/storage/storage'
import iconPath from '@/assets/image/icon_back.png'
import { TopicTypeMidjourney, TopicTypeOpenaiGPT3, TopicTypeOpenaiGPT4 } from '@/utils/constant/topic'
import eventBus from '@/utils/eventBus/eventBus'
import { WSMsg, WSConversationMidjourneyListenerEvent } from '@/utils/model/response/websocket'
import {
  WSMsgResponseTypeMidjourneyCreate,
  WSMsgResponseTypeMidjourneyEnd,
  WSMsgResponseTypeMidjourneyError,
  WSMsgResponseTypeMidjourneyInsertQueue,
  WSMsgResponseTypeMidjourneyProgress,
} from '@/utils/constant/websocket'
import {
  ActionTypeReRoll,
  ActionTypeUpscale,
  ActionTypeVariate,
  ActionTypePan,
  ActionTypeVary,
  ActionTypeZoomOut,
  ApplicationTypeMJ,
} from '@/utils/constant/conversation'
import InputParams from '@/views/conversation/inputParams.vue'
import AvatarImage from '@/components/avatarImage/AvatarImage.vue'
import {
  MidjourneyCustomForm,
  MidjourneySpeakForm,
  MidjourneySpeakFormKey,
  SpeakForm,
} from '@/views/conversation/script/model'

const appStore = getAppStore()
if (appStore.theme === 'dark') {
  // import('highlight.js/styles/dark.css').then()
  // import('github-markdown-css/github-markdown-dark.css').then()
}
const topicStore = getTopicStore()

const router = useRouter()
const route = useRoute()
const midjourneyInputEle = ref<any>(null)
const pageLoading = ref(true)

const topicTypes = ref([
  { label: '文字聊天 GPT3.5', value: TopicTypeOpenaiGPT3 },
  { label: '文字聊天 GPT4', value: TopicTypeOpenaiGPT4 },
  {
    label: '绘画聊天 Midjourney',
    value: TopicTypeMidjourney,
  },
])
const currentTopicType = ref(TopicTypeOpenaiGPT3)
const currentTopicId = ref('0')
const speakList = reactive<ResponseConversationSpeakItem[]>([])
const speakListPage = ref(1)
const speakForm = ref<SpeakForm>({
  topic_type: currentTopicType.value,
  stream_uuid: '',
  topic_id: '0',
  content: '',
})
const midjourneySpeakForm = ref<MidjourneySpeakForm>({
  topic_id: '0',
  content: '',
  application_type: ApplicationTypeMJ,
  no: '',
  images: '',
  seed: '',
  ar: '',
  chaos: '',
  quality: '',
  model: '',
  style: '',
  stylize: '',
  tile: '',
  iw: '',
})
const midjourneyCustomForm = ref<MidjourneyCustomForm>({
  refer_conversation_id: '',
  action_type: ActionTypeUpscale,
  index: 1,
  custom_id: '',
})
const submitBtnLoading = ref(false)
const conversationListEle = ref<any>(null)
const inputContent = ref('')

const userInfoData = storage.getUserInfo()

const gptExplainData = [
  {
    icon: 'precise-monitor',
    title: '几个例子',
    items: [
      '用简单的易懂的话解释什么是量子力学？',
      '马上要过10岁生日了，请给一些有创意的点子。',
      '我怎么样用javascript去发一个http请求？',
    ],
  },
  {
    icon: 'chart-bubble',
    title: '什么能力',
    items: ['记住之前的内容，联系上下文对话', '允许用户编辑和删除话题', '用你的标准训练智能模型，给出合适的回答'],
  },
  {
    icon: 'error-circle',
    title: '一些限制',
    items: [
      '可能会给出不够准确的结果，需要您自行甄别',
      '不允许提交涉及政治和违法的提问，法律责任自负',
      '不要提交隐私信息，可以用其他文字代替',
    ],
  },
]

const mjExplainData = [
  {
    icon: 'precise-monitor',
    title: '几个例子',
    items: [
      'App图标设计: squared with round edges mobile app logo design, flat vector app icon of an open box, minimalistic, white background',
      '漂亮的笔记本便签纸: sticker design, lined paper for writing, cute, vector',
      '用图片去作为提示词: http://xxx.com/xx.jpg, a bird',
    ],
  },
  {
    icon: 'error-circle',
    title: '一些限制',
    items: [
      '请用英语作为提示词，每个提示词用英文逗号隔开',
      '不允许提交涉及政治和违法的提问，法律责任自负',
      '不要提交隐私信息，可以用其他文字代替',
    ],
  },
]

const explainData = computed(() => {
  if (currentTopicType.value === TopicTypeOpenaiGPT3 || currentTopicType.value === TopicTypeOpenaiGPT4) {
    return gptExplainData
  }
  if (currentTopicType.value === TopicTypeMidjourney) {
    return mjExplainData
  }
  return []
})

const handleInitSpeakForm = () => {
  speakForm.value.stream_uuid = ''
  speakForm.value.content = ''
}

const handleInitMidjourneySpeakForm = () => {
  midjourneySpeakForm.value.content = ''
  midjourneySpeakForm.value.no = ''
  midjourneySpeakForm.value.images = ''
  midjourneySpeakForm.value.seed = ''
  midjourneySpeakForm.value.ar = ''
  midjourneySpeakForm.value.chaos = ''
  midjourneySpeakForm.value.quality = ''
  midjourneySpeakForm.value.model = ''
  midjourneySpeakForm.value.stylize = ''
  midjourneySpeakForm.value.tile = ''
  midjourneySpeakForm.value.iw = ''
}

const handleInitMidjourneyCustomForm = () => {
  midjourneyCustomForm.value.refer_conversation_id = ''
  midjourneyCustomForm.value.action_type = ActionTypeUpscale
  midjourneyCustomForm.value.index = 1
  midjourneyCustomForm.value.custom_id = ''
}

const handleGetStreamUuid = async () => {
  const uuidResponse = (await http.get('conversation/stream-uuid')) as ResponseConversationStreamUuid
  speakForm.value.stream_uuid = uuidResponse.uuid
}

const handleListenListHeight = (oriListHeight: number) => {
  let newListHeight = 0
  let loopTime = 0
  const listTime = setInterval(() => {
    if (loopTime > 100) {
      clearTimeout(listTime)
      return
    }
    if (conversationListEle.value.$el.firstElementChild.clientHeight !== oriListHeight) {
      newListHeight = conversationListEle.value.$el.firstElementChild.clientHeight
      clearTimeout(listTime)
      console.log('拿数据后', newListHeight)
      conversationListEle.value.$el.scrollTop = newListHeight - oriListHeight
    }
    loopTime++
  }, 10)
}

const handleGetTopicDetail = async () => {
  if (currentTopicId.value !== '' && currentTopicId.value !== '0') {
    const topicDataResponse = (await http.get(`conversation/topic-detail`, {
      topic_id: currentTopicId.value,
    })) as ResponseConversationTopic
    currentTopicType.value = topicDataResponse.type
  }
}

const handleGetSpeakListData = async () => {
  //  拿到列表高度
  const oriListHeight = conversationListEle.value.$el.firstElementChild.clientHeight
  console.log('拿数据前', oriListHeight)
  const speakListResponse = (await http.get(`conversation/speak-list`, {
    topic_id: currentTopicId.value,
    page: speakListPage.value,
  })) as ResponsePage
  if (speakListResponse.page_count > 0) {
    console.log('有数据')
    for (const item of speakListResponse.list_data) {
      speakList.unshift(item as ResponseConversationSpeakItem)
    }
    handleListenListHeight(oriListHeight)
    setTimeout(() => {
      hljs.highlightAll()
    }, 100)
  } else if (speakListPage.value > 1) {
    speakListPage.value -= 1
  }
}

const handleScroll = (event: any) => {
  if (event.scrollTop <= 0) {
    console.log('到顶了')
    speakListPage.value += 1
    handleGetSpeakListData()
  } else if (event.scrollBottom <= 0) {
    console.log('到底了')
  }
}

const handleScrollToBottom = () => {
  setTimeout(() => {
    if (conversationListEle.value !== null) {
      conversationListEle.value.$el.scrollTop = conversationListEle.value.$el.firstElementChild.clientHeight
    }
  }, 100)
}

const handleGptSpeak = async () => {
  // 获取新的uuid
  await handleGetStreamUuid()
  // 首先建立流连接
  const eventSource = new EventSource(
    `${http.requestApiHost()}/conversation/es/speak-stream?stream_uuid=${speakForm.value.stream_uuid}`,
  )
  let newTopicId = currentTopicId.value
  let newTopicTitle = ''
  let newTopicType = currentTopicType.value
  eventSource.addEventListener('open', async () => {
    console.log('Connected to stream')
    try {
      // 连接成功后，提交问题
      const conversationSpeakResponse = (await http.post(
        'conversation/speak',
        speakForm.value,
      )) as ResponseConversationSpeak
      console.log(conversationSpeakResponse)
      if (currentTopicId.value === '0') {
        newTopicId = conversationSpeakResponse.topic_id
        newTopicTitle = conversationSpeakResponse.title
        newTopicType = conversationSpeakResponse.topic_type
        console.log(newTopicId)
      }
    } catch (e) {
      console.log(e)
      eventSource.close()
      speakList.splice(-2)
      submitBtnLoading.value = false
    }
    handleInitSpeakForm()
  })
  eventSource.addEventListener('error', (event: any) => {
    console.error('Error connecting to stream:', event)
    if (event.readyState === EventSource.CLOSED) {
      console.log('event was closed')
    }
  })
  eventSource.addEventListener('close', (event: any) => {
    console.log('close:', event)
    eventSource.close()
    submitBtnLoading.value = false
    if (event.data === 'close') {
      setTimeout(() => {
        // 高亮代码
        hljs.highlightAll()
        // 如果是开启的新话题，则需要把左侧话题加入新话题
        console.log(currentTopicId.value)
        console.log(newTopicId)
        if (currentTopicId.value === '0') {
          topicStore.unshift({
            id: newTopicId,
            title: newTopicTitle,
            type: newTopicType,
            tool_show_status: false,
            edit_status: false,
          })
          // 跳转到新的topicId
          router.push(`/conversation/${newTopicId}`)
        }
      }, 100)
    }
  })
  eventSource.addEventListener('message', async (event) => {
    console.log('Received event:', event.data)
    const dataObj = JSON.parse(event.data)
    if (dataObj.error === 0) {
      speakList[speakList.length - 1].content += dataObj.message.content
      handleScrollToBottom()
    } else if (dataObj.error === 400) {
      speakList.splice(-2)
      await MessagePlugin.error(dataObj.message)
    }
  })
}

const handleMidjourneySpeak = async () => {
  let newTopicId = currentTopicId.value
  // 提交问题
  try {
    const conversationSpeakResponse = (await http.post(
      'conversation/midjourney-speak',
      midjourneySpeakForm.value,
    )) as ResponseConversationMidjourneySpeak
    console.log(conversationSpeakResponse)
    if (currentTopicId.value === '0') {
      newTopicId = conversationSpeakResponse.topic_id
      console.log(newTopicId)
    }
    speakList[speakList.length - 1].id = conversationSpeakResponse.answer_id
    speakList[speakList.length - 2].id = conversationSpeakResponse.question_id
    speakList[speakList.length - 2].content = conversationSpeakResponse.question_content
  } catch (e) {
    console.log(e)
    speakList.splice(-2)
  }
  handleInitMidjourneySpeakForm()
  // 提交完就允许再次输入，允许会员可以并发任务
  submitBtnLoading.value = false
}

const handleMidjourneyCustom = async () => {
  // 提交问题
  try {
    const conversationSpeakResponse = (await http.post(
      'conversation/midjourney-custom',
      midjourneyCustomForm.value,
    )) as ResponseConversationMidjourneySpeak
    console.log(conversationSpeakResponse)
    speakList[speakList.length - 1].id = conversationSpeakResponse.answer_id
    speakList[speakList.length - 2].id = conversationSpeakResponse.question_id
    speakList[speakList.length - 2].content = conversationSpeakResponse.question_content
  } catch (e) {
    console.log(e)
    speakList.splice(-2)
  }
  handleInitMidjourneyCustomForm()
  // 提交完就允许再次输入，允许会员可以并发任务
  submitBtnLoading.value = false
}

const handleOnMidjourneyInsertQueue = (speakListIndex: number, data: WSConversationMidjourneyListenerEvent) => {
  console.log('进入队列', data)
  speakList[speakListIndex].mj_data.progress_text = '开始奋力出图，请耐心等待'
}

const handleOnMidjourneyCreate = (speakListIndex: number, data: WSConversationMidjourneyListenerEvent) => {
  console.log('图片开始创建', data)
  speakList[speakListIndex].mj_data.progress_text = '正在构建图片模型，请耐心等待'
}

const handleOnMidjourneyEnd = (speakListIndex: number, data: WSConversationMidjourneyListenerEvent) => {
  console.log('图片创建完成', data)
  speakList[speakListIndex].content = data.content
  speakList[speakListIndex].mj_data.img_url = data.img_url
  speakList[speakListIndex].mj_data.thumbnail_img_url = data.thumbnail_img_url
  speakList[speakListIndex].mj_data.action_type = data.action_type
  speakList[speakListIndex].mj_data.components = data.components
  if (data.referenced_components.length > 0) {
    for (let i = 0; i < speakList.length; i++) {
      if (speakList[i].id === data.referenced_conversation_id) {
        speakList[i].mj_data.components = data.referenced_components
      }
    }
  }
  // 如果是开启的新话题，则需要把左侧话题加入新话题
  console.log(currentTopicId.value)
  if (currentTopicId.value === '0') {
    topicStore.unshift({
      id: data.topic_id,
      title: data.topic_title,
      type: data.topic_type,
      tool_show_status: false,
      edit_status: false,
    })
    // 跳转到新的topicId
    router.push(`/conversation/${data.topic_id}`)
  }
  submitBtnLoading.value = false
}

const handleOnMidjourneyError = (speakListIndex: number, data: WSConversationMidjourneyListenerEvent) => {
  console.log('图片创建报错', data)
  speakList[speakListIndex].mj_data.error = data.error
  submitBtnLoading.value = false
  // 如果是开启的新话题，则需要把左侧话题加入新话题
  console.log(currentTopicId.value)
  if (currentTopicId.value === '0') {
    topicStore.unshift({
      id: data.topic_id,
      title: data.topic_title,
      type: data.topic_type,
      tool_show_status: false,
      edit_status: false,
    })
    // 跳转到新的topicId
    router.push(`/conversation/${data.topic_id}`)
  }
  submitBtnLoading.value = false
}

const handleOnMidjourneyProgress = (speakListIndex: number, data: WSConversationMidjourneyListenerEvent) => {
  console.log('图片创建进度', data)
  speakList[speakListIndex].mj_data.progress = data.progress
}

const handleSpeakOnWsMessage = (data: any) => {
  const dataObj = JSON.parse(data) as WSMsg
  const eventType = dataObj.type
  if (dataObj.data === undefined || dataObj.data === null) {
    return
  }
  const eventData = dataObj.data as WSConversationMidjourneyListenerEvent
  let speakListIndex = -1
  for (let i = 0; i < speakList.length; i++) {
    if (eventData.conversation_id === speakList[i].id) {
      speakListIndex = i
      break
    }
  }
  if (speakListIndex === -1) {
    return
  }
  switch (eventType) {
    case WSMsgResponseTypeMidjourneyInsertQueue:
      handleOnMidjourneyInsertQueue(speakListIndex, eventData)
      break
    case WSMsgResponseTypeMidjourneyCreate:
      handleOnMidjourneyCreate(speakListIndex, eventData)
      break
    case WSMsgResponseTypeMidjourneyEnd:
      handleOnMidjourneyEnd(speakListIndex, eventData)
      break
    case WSMsgResponseTypeMidjourneyError:
      handleOnMidjourneyError(speakListIndex, eventData)
      break
    case WSMsgResponseTypeMidjourneyProgress:
      handleOnMidjourneyProgress(speakListIndex, eventData)
      break
    default:
  }
}

const handleSubmitSpeak = async (formName: string) => {
  if (submitBtnLoading.value) {
    return
  }
  submitBtnLoading.value = true
  const requestContent = inputContent.value
  if (formName === 'speakForm') {
    speakForm.value.topic_type = currentTopicType.value
    speakForm.value.content = inputContent.value
  } else if (formName === 'midjourneySpeakForm') {
    midjourneySpeakForm.value.content = inputContent.value
  }
  inputContent.value = ''
  // 把提交的内容渲染出来
  speakList.push({
    id: '',
    topic_id: currentTopicId.value,
    role: 'user',
    content: requestContent,
    mj_data: {
      action_type: 0,
      img_url: '',
      thumbnail_img_url: '',
      progress: 0,
      components: [],
      error: '',
    },
    created_at: 0,
  })
  handleScrollToBottom()
  // 渲染出回答的位置
  const assistantAnswer: ResponseConversationSpeakItem = {
    id: '',
    topic_id: currentTopicId.value,
    role: 'assistant',
    content: '',
    mj_data: {
      action_type: 0,
      img_url: '',
      thumbnail_img_url: '',
      progress: 0,
      components: [],
      error: '',
      progress_text: '正在排队，请稍后',
    },
    created_at: 0,
  }
  speakList.push(assistantAnswer)
  handleScrollToBottom()
  if (formName === 'speakForm') {
    await handleGptSpeak()
  } else if (formName === 'midjourneySpeakForm') {
    await handleMidjourneySpeak()
  } else if (formName === 'midjourneyCustomForm') {
    await handleMidjourneyCustom()
  }
}

const handleMidjourneyInputSubmit = async (midjourneyInputData: any) => {
  for (const key in midjourneyInputData) {
    if (key === 'topic_id') {
      continue
    }
    const formKey = key as MidjourneySpeakFormKey
    midjourneySpeakForm.value[formKey] = midjourneyInputData[key] as never
  }
  await handleSubmitSpeak('midjourneySpeakForm')
}

const handleMatchCustomActionType = (customId: string): number => {
  let actionType = 0
  if (customId.indexOf('upsample') !== -1) {
    actionType = ActionTypeUpscale
  } else if (customId.indexOf('high_variation') !== -1) {
    actionType = ActionTypeVary
  } else if (customId.indexOf('low_variation') !== -1) {
    actionType = ActionTypeVary
  } else if (customId.indexOf('variation') !== -1) {
    actionType = ActionTypeVariate
  } else if (customId.indexOf('reroll') !== -1) {
    actionType = ActionTypeReRoll
  } else if (customId.indexOf('Outpaint') !== -1) {
    actionType = ActionTypeZoomOut
  } else if (customId.indexOf('pan_') !== -1) {
    actionType = ActionTypePan
  }
  return actionType
}

const handleClickCustom = async (referConversationId: string, index: number, customId: string) => {
  midjourneyCustomForm.value.action_type = handleMatchCustomActionType(customId)
  midjourneyCustomForm.value.refer_conversation_id = referConversationId
  midjourneyCustomForm.value.index = index
  midjourneyCustomForm.value.custom_id = customId
  console.log(midjourneyCustomForm.value)
  await handleSubmitSpeak('midjourneyCustomForm')
}

onMounted(async () => {
  currentTopicId.value = route.params.topicId as string
  speakForm.value.topic_id = route.params.topicId as string
  midjourneySpeakForm.value.topic_id = route.params.topicId as string
  await handleGetTopicDetail()
  await handleGetSpeakListData()
  eventBus.on('conversationMj', handleSpeakOnWsMessage)
  pageLoading.value = false
})

onBeforeUnmount(() => {
  eventBus.off('conversationMj', handleSpeakOnWsMessage)
})
</script>

<template>
  <div class="main-space">
    <div class="conversation-space">
      <t-list
        v-show="!pageLoading && speakList.length > 0"
        ref="conversationListEle"
        class="conversation-list"
        @scroll="handleScroll"
      >
        <t-list-item
          v-for="(item, index) in speakList"
          :key="index"
          :style="item.role === 'assistant' ? 'background-color: #ffffff !important;' : ''"
        >
          <div class="conversation-item">
            <div class="conversation-avatar">
              <t-avatar v-if="item.role === 'user'" size="large">
                <avatar-image :url="userInfoData?.avatar"></avatar-image>
              </t-avatar>
              <t-avatar v-else-if="item.role === 'assistant'" size="large" :image="iconPath"></t-avatar>
              <t-avatar v-else size="large"> {{ item.role }}</t-avatar>
            </div>
            <div class="conversation-content">
              <div
                v-if="currentTopicType === TopicTypeMidjourney && item.role === 'assistant'"
                class="conversation-content-midjourney"
              >
                <t-skeleton
                  :loading="item.mj_data.img_url === ''"
                  animation="gradient"
                  :row-col="[
                    [1].map(() => ({
                      width: '100%',
                      height: '100%',
                      content:
                        item.mj_data.error !== ''
                          ? h('span', { style: { color: 'var(--td-error-color-5)' } }, item.mj_data.error)
                          : item.mj_data.progress === 0
                          ? h('span', { style: { color: 'var(--td-warning-color-5)' } }, item.mj_data.progress_text)
                          : '进入绘画阶段，进度: ' + item.mj_data.progress + '%',
                    })),
                  ]"
                >
                </t-skeleton>
                <t-image-viewer v-if="item.mj_data.img_url !== ''" :images="[item.mj_data.img_url]">
                  <template #trigger="{ open }">
                    <t-image
                      class="conversation-content-midjourney-image"
                      :src="item.mj_data.thumbnail_img_url"
                      shape="round"
                      fit="contain"
                      position="left"
                      @click="open"
                    />
                  </template>
                </t-image-viewer>
                <div class="conversation-content-midjourney-tool">
                  <div class="conversation-content-midjourney-button-group">
                    <div
                      v-for="(rowItem, rowIndex) in item.mj_data.components"
                      :key="rowIndex"
                      class="conversation-content-midjourney-button-row"
                    >
                      <template v-if="rowItem.type === 1">
                        <template v-for="(btnItem, btnIndex) in rowItem.components">
                          <t-button
                            v-if="
                              btnItem.label !== 'Web' &&
                              btnItem.custom_id.indexOf('BOOKMARK') === -1 &&
                              btnItem.custom_id.indexOf('CustomZoom') === -1 &&
                              btnItem.custom_id.indexOf('Inpaint') === -1
                            "
                            :key="btnIndex"
                            :theme="btnItem.style === 1 ? 'success' : 'default'"
                            class="conversation-content-midjourney-button"
                            :disabled="submitBtnLoading"
                            @click="handleClickCustom(item.id, btnIndex + 1, btnItem.custom_id)"
                            >{{ btnItem.emoji != null ? btnItem.emoji.name : '' }} {{ btnItem.label }}
                          </t-button>
                        </template>
                      </template>
                    </div>
                  </div>
                </div>
              </div>
              <div
                v-if="
                  currentTopicType === TopicTypeOpenaiGPT3 ||
                  currentTopicType === TopicTypeOpenaiGPT4 ||
                  item.role === 'user'
                "
                class="markdown-body"
                v-html="marked.parse(item.content)"
              ></div>
            </div>
          </div>
        </t-list-item>
      </t-list>
      <div v-show="!pageLoading && speakList.length <= 0" class="conversation-explain">
        <div class="conversation-explain-space">
          <div class="conversation-explain-title">
            <span>Chat AIT</span>
          </div>
          <div class="conversation-explain-topic">
            <t-select v-model="currentTopicType" style="width: 200px">
              <t-option
                v-for="(item, index) in topicTypes"
                :key="index"
                :value="item.value"
                :label="item.label"
              ></t-option>
            </t-select>
          </div>
          <div class="conversation-explain-content">
            <div v-for="(item, index) in explainData" :key="index" class="conversation-explain-info">
              <div class="conversation-explain-icon">
                <icon-font size="30px" :name="item.icon"></icon-font>
              </div>
              <div class="conversation-explain-items-title">
                <span>{{ item.title }}</span>
              </div>
              <div class="conversation-explain-items">
                <div v-for="(cItem, cIndex) in item.items" :key="cIndex" class="conversation-explain-item">
                  <span>{{ cItem }}</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div v-show="pageLoading" class="conversation-page-loading">
        <loading-icon size="20px"></loading-icon>
      </div>
    </div>

    <div class="input-space">
      <t-input
        v-if="currentTopicType === TopicTypeOpenaiGPT3 || currentTopicType === TopicTypeOpenaiGPT4"
        v-model="inputContent"
        class="content-input"
        size="large"
        :disabled="submitBtnLoading"
        @enter="handleSubmitSpeak('speakForm')"
      >
        <template #suffix-icon>
          <template v-if="submitBtnLoading">
            <loading-icon></loading-icon>
          </template>
          <template v-else>
            <enter-icon @click="handleSubmitSpeak('speakForm')"></enter-icon>
          </template>
        </template>
      </t-input>
      <input-params
        v-else-if="currentTopicType === TopicTypeMidjourney"
        ref="midjourneyInputEle"
        v-model="inputContent"
        :disabled="submitBtnLoading"
        @enter="handleMidjourneyInputSubmit"
        @submit="handleMidjourneyInputSubmit"
      ></input-params>
    </div>
  </div>
</template>

<style lang="scss">
.markdown-body {
  background: none !important;
}

.main-space {
  .t-list-item {
    background: none !important;
    border-bottom: 1px solid #e8e8fa;
  }

  .input-space .content-input {
    .t-input {
      border-style: none;
      box-shadow: rgba(99, 99, 99, 0.2) 0 2px 8px 0;
    }
  }

  .t-skeleton {
    width: 35%;
    aspect-ratio: 1;

    .t-skeleton__row {
      width: 100%;
      height: 100%;
    }
  }
}
</style>

<style lang="scss" scoped>
.main-space {
  display: flex;
  flex-direction: column;
  width: 100%;
  height: 100%;
  background-color: #f6f6fd;
  overflow: hidden;
  position: relative;

  .conversation-space {
    flex: 1 1 auto;
    width: 100%;
    height: 100%;
    overflow: hidden;

    .conversation-explain {
      width: 100%;
      height: 100%;
      overflow-y: auto;
      color: rgba(0, 0, 0, 0) !important;
      transition: color 0.5s !important;
      transition-delay: 1s !important;

      .conversation-explain-space {
        width: 70%;
        color: var(--td-text-color-secondary);
        padding-bottom: 100px;

        .conversation-explain-title {
          display: flex;
          width: 100%;
          height: 100px;
          justify-content: center;
          font-size: 35px;
          line-height: 80px;
          color: var(--td-text-color-primary);
        }

        .conversation-explain-topic {
          display: flex;
          width: 100%;
          justify-content: center;
          align-items: center;
        }

        .conversation-explain-content {
          display: flex;
          width: 100%;
          overflow: hidden;

          .conversation-explain-info {
            flex: 1 1 auto;
            width: 100%;
            box-sizing: border-box;
            padding: 5px;
            overflow: hidden;

            .conversation-explain-icon {
              display: flex;
              width: 100%;
              height: 50px;
              justify-content: center;
              align-items: center;
            }

            .conversation-explain-items-title {
              display: flex;
              width: 100%;
              height: 30px;
              justify-content: center;
              align-items: center;
              font-size: 16px;
            }

            .conversation-explain-items {
              width: 100%;
              box-sizing: border-box;
              overflow: hidden;

              .conversation-explain-item {
                background-color: #ffffff;
                line-height: 20px;
                font-size: 14px;
                border-radius: 5px;
                box-sizing: border-box;
                padding: 10px;
                margin: 15px;
                overflow: hidden;
              }
            }
          }
        }

        @media screen and (max-width: 768px) {
          .conversation-explain-content {
            display: block;
          }
        }
      }

      @media screen and (max-width: 768px) {
        .conversation-explain-space {
          width: 100%;
        }
      }
    }

    .conversation-explain:hover {
      color: rgba(0, 0, 0, 0.2) !important;
      transition-delay: 0s !important;
    }

    .conversation-explain::-webkit-scrollbar {
      width: 6px;
      height: 4px;
    }

    .conversation-explain::-webkit-scrollbar-thumb {
      border-radius: 3px;
      box-shadow: inset 0 0 0 10px;
    }

    @media screen and (min-width: 768px) {
      .conversation-explain {
        display: flex;
        justify-content: center;
        align-items: center;
      }
    }

    .conversation-page-loading {
      display: flex;
      width: 100%;
      height: 100%;
      justify-content: center;
      align-items: center;
      color: var(--td-text-color-secondary);
    }

    .conversation-list {
      width: 100%;
      height: 100%;
      box-sizing: border-box;
      padding-bottom: 100px;
      overflow-y: auto;

      .conversation-item {
        display: flex;
        width: 100%;
        box-sizing: border-box;
        overflow: hidden;

        .conversation-avatar {
          display: flex;
          align-items: center;
          justify-content: end;
          flex: 0 0 auto;
          width: 110px;
          height: 100%;
          overflow: hidden;
        }

        @media screen and (max-width: 768px) {
          .conversation-avatar {
            width: 50px;
          }
        }

        .conversation-content {
          flex: 1 1 auto;
          width: 100%;
          box-sizing: border-box;
          padding: 10px 30px;
          overflow: hidden;

          .conversation-content-midjourney-image {
            width: 35%;
            aspect-ratio: 1;
            background: none;
            box-sizing: border-box;
            overflow: hidden;
          }

          @media screen and (max-width: 768px) {
            .conversation-content-midjourney-image {
              width: 100%;
            }
          }

          .conversation-content-midjourney-tool {
            width: 100%;
            display: flex;
            padding: 5px 0;
            box-sizing: border-box;
            overflow: hidden;

            .conversation-content-midjourney-button-group {
              flex: 0 0 auto;
              width: 100%;

              .conversation-content-midjourney-button-row {
                width: 100%;
                padding: 5px 0;
                box-sizing: border-box;

                .conversation-content-midjourney-button {
                  margin-right: 10px;
                }
              }
            }
          }
        }
      }
    }

    .conversation-list {
      position: relative;
      display: block;
      color: rgba(0, 0, 0, 0) !important;
      transition: color 0.5s !important;
      transition-delay: 1s !important;
    }

    .conversation-list:hover {
      color: rgba(0, 0, 0, 0.2) !important;
      transition-delay: 0s !important;
    }

    .conversation-list::-webkit-scrollbar {
      width: 6px;
      height: 4px;
    }

    .conversation-list::-webkit-scrollbar-thumb {
      border-radius: 3px;
      box-shadow: inset 0 0 0 10px;
    }
  }

  .input-space {
    position: absolute;
    bottom: 0;
    width: 100%;
    height: 100px;
    padding: 40px 60px 20px 85px;
    box-sizing: border-box;
    //overflow: hidden;
    background-image: linear-gradient(to bottom, transparent, #f6f6fd 45%);
    z-index: 10;
  }

  @media screen and (max-width: 768px) {
    .input-space {
      padding: 40px 20px 20px 20px;
    }
  }
}
</style>
