/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

import { defineStore } from 'pinia'
import { store } from '@/store'
import http from '@/utils/network/http'
import { ResponseConversationTopic } from '@/utils/model/response/conversation'
import { AppConversationTopic } from '@/utils/model/app/topic'

const getTopicList = async (lastId: string): Promise<ResponseConversationTopic[]> => {
  return (await http.get(
    `conversation/topic-list-by-last-id?last_id=${lastId}&limit=100`,
  )) as ResponseConversationTopic[]
}

export const useTopicStore = defineStore('topic', {
  state: () => ({
    topicList: <AppConversationTopic[]>[],
    lastId: '0',
  }),
  getters: {},
  actions: {
    async getList() {
      const listData = (await getTopicList(this.lastId)) as ResponseConversationTopic[]
      if (listData.length > 0) {
        for (let i = 0; i < listData.length; i++) {
          this.topicList.push({
            id: listData[i].id,
            title: listData[i].title,
            type: listData[i].type,
            tool_show_status: false,
            edit_status: false,
          })
        }
        this.lastId = listData[listData.length - 1].id
      }
    },
    unshift(topicItem: AppConversationTopic) {
      this.topicList.unshift(topicItem)
    },
    del(index: number) {
      this.topicList.splice(index, 1)
    },
    edit(index: number, title: string) {
      this.topicList[index].title = title
    },
    init() {
      this.topicList = []
    },
  },
})

export function getTopicStore() {
  return useTopicStore(store)
}
