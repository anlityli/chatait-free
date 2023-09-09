<!--
  - Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<template>
  <div v-if="showFrame">
    <template v-for="frame in getFramePages" :key="frame.path">
      <frame-content v-if="hasRenderFrame(frame.name)" v-show="showIframe(frame)" :frame-src="frame.meta.frameSrc" />
    </template>
  </div>
</template>
<script lang="ts">
import { computed, defineComponent, unref } from 'vue'

import FrameContent from '../components/FrameContent.vue'
import { useFrameKeepAlive } from './useFrameKeepAlive'

export default defineComponent({
  name: 'FrameLayout',
  components: { FrameContent },
  setup() {
    const { getFramePages, hasRenderFrame, showIframe } = useFrameKeepAlive()

    const showFrame = computed(() => unref(getFramePages).length > 0)

    return { getFramePages, hasRenderFrame, showIframe, showFrame }
  },
})
</script>
