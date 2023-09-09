<!--
  - Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<template>
  <img :src="realUrl" alt="0" />
</template>

<script lang="ts" setup>
import { computed } from 'vue'
import avatar0 from '@/assets/avatar/0.png'
import avatar1 from '@/assets/avatar/1.png'
import avatar2 from '@/assets/avatar/2.png'
import avatar3 from '@/assets/avatar/3.png'
import avatar4 from '@/assets/avatar/4.png'
import avatar5 from '@/assets/avatar/5.png'
import avatar6 from '@/assets/avatar/6.png'
import avatar7 from '@/assets/avatar/7.png'
import avatar8 from '@/assets/avatar/8.png'
import avatar9 from '@/assets/avatar/9.png'
import avatar10 from '@/assets/avatar/10.png'
import avatar11 from '@/assets/avatar/11.png'
import avatar12 from '@/assets/avatar/12.png'
import avatar13 from '@/assets/avatar/13.png'
import avatar14 from '@/assets/avatar/14.png'
import avatar15 from '@/assets/avatar/15.png'
import avatar16 from '@/assets/avatar/16.png'
import avatar17 from '@/assets/avatar/17.png'
import avatar18 from '@/assets/avatar/18.png'
import avatar19 from '@/assets/avatar/19.png'
import avatar20 from '@/assets/avatar/20.png'

const props = defineProps({
  url: {
    type: String,
    default: '',
  },
})

const realUrl = computed(() => {
  switch (props.url) {
    case '':
      return avatar0
    case '/avatar/1.png':
      return avatar1
    case '/avatar/2.png':
      return avatar2
    case '/avatar/3.png':
      return avatar3
    case '/avatar/4.png':
      return avatar4
    case '/avatar/5.png':
      return avatar5
    case '/avatar/6.png':
      return avatar6
    case '/avatar/7.png':
      return avatar7
    case '/avatar/8.png':
      return avatar8
    case '/avatar/9.png':
      return avatar9
    case '/avatar/10.png':
      return avatar10
    case '/avatar/11.png':
      return avatar11
    case '/avatar/12.png':
      return avatar12
    case '/avatar/13.png':
      return avatar13
    case '/avatar/14.png':
      return avatar14
    case '/avatar/15.png':
      return avatar15
    case '/avatar/16.png':
      return avatar16
    case '/avatar/17.png':
      return avatar17
    case '/avatar/18.png':
      return avatar18
    case '/avatar/19.png':
      return avatar19
    case '/avatar/20.png':
      return avatar20
    default:
      return props.url
  }
})
</script>
