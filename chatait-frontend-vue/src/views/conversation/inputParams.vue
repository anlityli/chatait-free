<!--
  - Copyright 2024 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<template>
  <div class="input-params-wrap">
    <t-input
      v-model="form.content"
      class="content-input"
      size="large"
      :disabled="disabled"
      @change="handleChange"
      @enter="handleEnter"
    >
      <template #suffix-icon>
        <div class="btn-group">
          <div class="btn-group-item">
            <filter2-icon @click="showTools = !showTools"></filter2-icon>
          </div>
          <div class="btn-group-item">
            <template v-if="disabled">
              <loading-icon></loading-icon>
            </template>
            <template v-else>
              <enter-icon @click="handleSubmit"></enter-icon>
            </template>
          </div>
        </div>
      </template>
    </t-input>
    <div v-show="showTools" class="input-params-tools">
      <t-form :data="form">
        <t-form-item class="input-params-row" label="应用机器人" name="application_type">
          <t-radio-group v-model="form.application_type" :default-value="ApplicationTypeMJ" size="small">
            <t-radio-button :value="ApplicationTypeMJ">Midjourney</t-radio-button>
            <t-radio-button :value="ApplicationTypeNJ">Niji</t-radio-button>
          </t-radio-group>
        </t-form-item>
        <t-form-item class="input-params-row" label="反向提示" name="no">
          <t-textarea v-model="form.no" :autosize="{ maxRows: 3 }"></t-textarea>
        </t-form-item>
        <t-form-item class="input-params-row" label="图片提示" name="images">
          <t-input v-model="form.images" placeholder="请输入图片网址"></t-input>
        </t-form-item>
        <template v-for="(item, index) in mjTools" :key="index">
          <t-form-item
            v-if="item.key !== 'iw' && item.key !== 'model' && item.key !== 'style'"
            class="input-params-row"
            :label="item.label"
            :name="item.key"
          >
            <t-radio-group v-model="form[item.key]" :default-value="item.params[item.default].value" size="small">
              <t-radio-button v-for="(btnItem, btnIndex) in item.params" :key="btnIndex" :value="btnItem.value">
                {{ btnItem.label }}
              </t-radio-button>
            </t-radio-group>
          </t-form-item>
          <t-form-item
            v-else-if="item.key === 'iw' && form.images !== ''"
            class="input-params-row"
            :label="item.label"
            :name="item.key"
          >
            <t-radio-group v-model="form[item.key]" :default-value="item.params[item.default].value" size="small">
              <t-radio-button v-for="(btnItem, btnIndex) in item.params" :key="btnIndex" :value="btnItem.value">
                {{ btnItem.label }}
              </t-radio-button>
            </t-radio-group>
          </t-form-item>
          <t-form-item v-else-if="item.key === 'model'" class="input-params-row" :label="item.label" :name="item.key">
            <t-radio-group v-model="form[item.key]" :default-value="item.params[item.default].value" size="small">
              <template v-for="(btnItem, btnIndex) in item.params" :key="btnIndex">
                <t-radio-button v-if="btnItem.value === ''" :value="btnItem.value">
                  {{ btnItem.label }}
                </t-radio-button>
                <t-radio-button
                  v-else-if="form.application_type === ApplicationTypeMJ && btnItem.label.indexOf('MJ') !== -1"
                  :value="btnItem.value"
                >
                  {{ btnItem.label }}
                </t-radio-button>
                <t-radio-button
                  v-else-if="form.application_type === ApplicationTypeNJ && btnItem.label.indexOf('Niji') !== -1"
                  :value="btnItem.value"
                >
                  {{ btnItem.label }}
                </t-radio-button>
              </template>
            </t-radio-group>
          </t-form-item>
          <t-form-item v-else-if="item.key === 'style'" class="input-params-row" :label="item.label" :name="item.key">
            <t-radio-group v-model="form[item.key]" :default-value="item.params[item.default].value" size="small">
              <template v-for="(btnItem, btnIndex) in item.params" :key="btnIndex">
                <t-radio-button v-if="btnItem.value === ''" :value="btnItem.value">
                  {{ btnItem.label }}
                </t-radio-button>
                <t-radio-button
                  v-else-if="form.application_type === ApplicationTypeMJ && btnItem.label.indexOf('MJ') !== -1"
                  :value="btnItem.value"
                >
                  {{ btnItem.label }}
                </t-radio-button>
                <t-radio-button
                  v-else-if="form.application_type === ApplicationTypeNJ && btnItem.label.indexOf('Niji') !== -1"
                  :value="btnItem.value"
                >
                  {{ btnItem.label }}
                </t-radio-button>
              </template>
            </t-radio-group>
          </t-form-item>
        </template>
        <t-form-item class="input-params-row" label="记住参数">
          <t-radio-group
            v-model="appStore.saveMjSpeakParams"
            :default-value="false"
            size="small"
            @change="handleChangeSaveMjParams"
          >
            <t-radio-button :value="true">开启</t-radio-button>
            <t-radio-button :value="false">关闭</t-radio-button>
          </t-radio-group>
        </t-form-item>
      </t-form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref, toRefs, watch } from 'vue'
import { EnterIcon, Filter2Icon, LoadingIcon } from 'tdesign-icons-vue-next'
import { mjTools } from '@/views/conversation/script/mjTools'
import { MidjourneySpeakForm } from '@/views/conversation/script/model'
import { ApplicationTypeMJ, ApplicationTypeNJ } from '@/utils/constant/conversation'
import { useAppStore } from '@/store'

const props = defineProps({
  modelValue: {
    type: String,
    default: '',
  },
  disabled: {
    type: Boolean,
    default: false,
  },
})
const { modelValue } = toRefs(props)
const emit = defineEmits(['update:modelValue', 'enter', 'submit', 'change'])
const showTools = ref(false)
const appStore = useAppStore()

const form = ref<MidjourneySpeakForm>({
  topic_id: '',
  content: modelValue?.value,
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

watch(
  () => modelValue.value,
  () => {
    form.value.content = modelValue?.value
  },
)

const handleChange = (event: any) => {
  emit('update:modelValue', form.value.content)
  emit('change', event)
}

const handleInitForm = () => {
  form.value = {
    topic_id: '',
    content: modelValue?.value,
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
  }
}

const handleEnter = (event: any) => {
  showTools.value = false
  emit('enter', form.value, event)
  // handleInitForm()
  // 只清除反向提示和图片提示，其他的参数不变
  form.value.content = modelValue?.value
  form.value.no = ''
  form.value.images = ''
  appStore.setMjSpeakParams(form.value)
}

const handleSubmit = () => {
  showTools.value = false
  emit('submit', form.value)
  // handleInitForm()
  // 只清除反向提示和图片提示，其他的参数不变
  form.value.content = modelValue?.value
  form.value.no = ''
  form.value.images = ''
  appStore.setMjSpeakParams(form.value)
}

const handleChangeSaveMjParams = () => {
  console.log(appStore.saveMjSpeakParams)
  appStore.setMjSpeakParams(form.value)
}

onMounted(async () => {
  if (appStore.saveMjSpeakParams) {
    form.value.application_type = appStore.mjSpeakParams.application_type
    form.value.ar = appStore.mjSpeakParams.ar
    form.value.chaos = appStore.mjSpeakParams.chaos
    form.value.quality = appStore.mjSpeakParams.quality
    form.value.model = appStore.mjSpeakParams.model
    form.value.style = appStore.mjSpeakParams.style
    form.value.stylize = appStore.mjSpeakParams.stylize
    form.value.tile = appStore.mjSpeakParams.tile
    form.value.iw = appStore.mjSpeakParams.iw
  }
})

defineExpose({
  handleInitForm,
})
</script>

<style lang="scss" scoped>
.input-params-wrap {
  position: relative;
  width: 100%;
  height: 100%;

  .btn-group {
    width: 60px;
    display: flex;

    .btn-group-item {
      flex: 1 1 auto;
      width: 100%;
      display: flex;
      justify-content: center;
      align-items: center;
    }
  }

  .t-list-item {
    border-bottom: none;
  }

  .input-params-tools {
    position: absolute;
    width: 100%;
    max-height: 460px;
    bottom: 40px;
    left: 0;
    z-index: 999;
    background-color: #fff;
    box-shadow: rgba(99, 99, 99, 0.2) 0 -2px 8px 0;
    padding: 20px;
    box-sizing: border-box;
    overflow-y: auto;

    .input-params-row {
      margin-bottom: 10px;
    }
  }
}
</style>
