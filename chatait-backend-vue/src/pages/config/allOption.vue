<!--
  - Copyright 2024 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<template>
  <div class="main-content-container">
    <t-form class="edit-form" :label-width="220" :data="formData" :colon="true">
      <div v-for="(item, index) in optionList" :key="index">
        <t-divider>{{ item.group_name }}</t-divider>
        <t-form-item
          v-for="(itemC, indexC) in item.group_items"
          :key="indexC"
          :label="itemC.title"
          :name="itemC.config_name"
        >
          <template v-if="itemC.input_type === OptionTypeSelect">
            <t-select v-model="itemC.value" class="edit-input" @change="handleSubmit(itemC.config_name, itemC.value)">
              <t-option
                v-for="(optionItem, optionIndex) in itemC.options"
                :key="optionIndex"
                :label="optionItem.label"
                :value="optionItem.value"
              />
            </t-select>
          </template>
          <template v-else-if="itemC.input_type === OptionTypeCheckbox">
            <t-select
              v-model="itemC.value"
              class="edit-input"
              multiple
              @change="handleSubmit(itemC.config_name, itemC.value)"
            >
              <t-option label="全选" :check-all="true" />
              <t-option
                v-for="(optionItem, optionIndex) in itemC.options"
                :key="optionIndex"
                :label="optionItem.label"
                :value="optionItem.value"
              />
            </t-select>
          </template>
          <template v-else-if="itemC.input_type === OptionTypeTextarea">
            <t-input v-model="itemC.value" class="edit-input" @change="handleSubmit(itemC.config_name, itemC.value)">
              <template #suffix>
                {{ itemC.unit !== '' ? itemC.unit : '' }}
              </template>
            </t-input>
          </template>
          <template v-else>
            <t-input v-model="itemC.value" class="edit-input" @change="handleSubmit(itemC.config_name, itemC.value)">
              <template #suffix>
                {{ itemC.unit !== '' ? itemC.unit : '' }}
              </template>
            </t-input>
          </template>
        </t-form-item>
      </div>
    </t-form>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import http from '@/utils/network/http'
import { OptionListItem } from '@/pages/config/model/option'
import { ResponseConfigAllOption } from '@/utils/model/response/config'
import { OptionTypeCheckbox, OptionTypeSelect, OptionTypeTextarea } from '@/pages/config/model/constant'
import tool from '@/utils/tool/tool'

const router = useRouter()
const route = useRoute()

const formData = ref({})

const optionList = ref<OptionListItem[]>([])

const handleSubmit = async (configName: string, value: string) => {
  if (
    configName === `newUserAddBalance` ||
    configName === 'newUserAddGpt3' ||
    configName === 'newUserAddGpt4' ||
    configName === 'newUserAddMidjourney' ||
    configName === 'gpt3UseBalance' ||
    configName === 'gpt4UseBalance' ||
    configName === 'midjourneyUseBalance'
  ) {
    value = tool.yuanToCent(value).toString()
  } else if (configName === 'allowTopicType') {
    value = JSON.stringify(value)
  }
  const requestData = {
    config_name: configName,
    value,
  }
  await http.post(`config/option-edit`, requestData)
}

onMounted(async () => {
  const response = (await http.get(`config/all-option`)) as ResponseConfigAllOption[]
  let tempGroup = ''
  for (let i = 0; i < response.length; i++) {
    const configName = response[i].config_name
    if (
      configName === `newUserAddBalance` ||
      configName === 'newUserAddGpt3' ||
      configName === 'newUserAddGpt4' ||
      configName === 'newUserAddMidjourney' ||
      configName === 'gpt3UseBalance' ||
      configName === 'gpt4UseBalance' ||
      configName === 'midjourneyUseBalance'
    ) {
      response[i].value = tool.centToYuan(Number(response[i].value))
    } else if (configName === 'allowTopicType') {
      response[i].value = JSON.parse(response[i].value)
    }
    if (tempGroup === '' || tempGroup !== response[i].type) {
      tempGroup = response[i].type
      optionList.value.push({
        group_name: response[i].type,
        group_items: [response[i]],
      })
    } else {
      const lastIndex = optionList.value.length - 1
      optionList.value[lastIndex].group_items.push(response[i])
    }
  }
})
</script>
