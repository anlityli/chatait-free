/*
 * Copyright 2024 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

import { MidjourneySpeakFormKey } from '@/views/conversation/script/model'

export const mjTools = [
  {
    key: <MidjourneySpeakFormKey>'model',
    label: '模型',
    description: '不同模型版本',
    default: 0,
    params: [
      { label: '默认', value: '' },
      { label: 'MJ4', value: '--v 4' },
      { label: 'MJ5', value: '--v 5' },
      { label: 'MJ5.1', value: '--v 5.1' },
      { label: 'MJ5.2', value: '--v 5.2' },
      { label: 'MJ6', value: '--v 6' },
      { label: 'Niji5', value: '--niji 5' },
      { label: 'Niji6', value: '--niji 6' },
    ],
  },
  {
    key: <MidjourneySpeakFormKey>'style',
    label: '风格',
    description: '不同模型的艺术表现形式',
    default: 0,
    params: [
      { label: '默认', value: '' },
      { label: 'MJ raw', value: '--style raw' },
      { label: 'Niji original', value: '--style original' },
      { label: 'Niji cute', value: '--style cute' },
      { label: 'Niji expressive', value: '--style expressive' },
      { label: 'Niji scenic', value: '--style scenic' },
    ],
  },
  {
    key: <MidjourneySpeakFormKey>'iw',
    label: '图片权重',
    description: '该参数会不同程度的保留原图片的样式，数值越高越依从原图片',
    default: 2,
    params: [
      { label: '0.5', value: '--iw .5' },
      { label: '0.75', value: '--iw .75' },
      { label: '1', value: '' },
      { label: '1.25', value: '--iw 1.25' },
      { label: '1.5', value: '--iw 1.5' },
      { label: '1.75', value: '--iw 1.75' },
      { label: '2', value: '--iw 2' },
    ],
  },
  {
    key: <MidjourneySpeakFormKey>'ar',
    label: '比例',
    description: '该参数调整生成图片的比例',
    default: 0,
    params: [
      { label: '默认', value: '' },
      { label: '1:1', value: '--ar 1:1' },
      { label: '4:5', value: '--ar 4:5' },
      { label: '5:4', value: '--ar 5:4' },
      { label: '2:3', value: '--ar 2:3' },
      { label: '3:2', value: '--ar 3:2' },
      { label: '4:7', value: '--ar 4:7' },
      { label: '7:4', value: '--ar 7:4' },
      { label: '9:16', value: '--ar 9:16' },
      { label: '16:9', value: '--ar 16:9' },
    ],
  },
  {
    key: <MidjourneySpeakFormKey>'chaos',
    label: '复杂度',
    description: '复杂度越高，产生的图片可能更会出现意想不到的结果',
    default: 0,
    params: [
      { label: '无', value: '' },
      { label: '低', value: '--c 10' },
      { label: '中', value: '--c 25' },
      { label: '高', value: '--c 50' },
      { label: '超高', value: '--c 80' },
    ],
  },
  {
    key: <MidjourneySpeakFormKey>'quality',
    label: '图片质量',
    description:
      '该参数改变生成图像所花费的时间。更高质量的设置需要更长的时间来处理和产生更多的细节。更高的值也意味着每个作业使用更多的GPU分钟。质量设置不影响分辨率',
    default: 2,
    params: [
      { label: '低清', value: '--q .25' },
      { label: '标清', value: '--q .5' },
      { label: '高清', value: '' },
    ],
  },
  {
    key: <MidjourneySpeakFormKey>'stylize',
    label: '艺术化度',
    description: '默认值100 不同的艺术化度，对色彩和整体构图会有影响',
    default: 2,
    params: [
      { label: '0', value: '--s 0' },
      { label: '50', value: '--s 50' },
      { label: '100', value: '' },
      { label: '250', value: '--s 250' },
      { label: '500', value: '--s 500' },
      { label: '750', value: '--s 750' },
    ],
  },
  {
    key: <MidjourneySpeakFormKey>'tile',
    label: '瓷砖化',
    description: '该参数产生瓷砖、墙纸化等效果',
    default: 0,
    params: [
      { label: '不开启', value: '' },
      { label: '开启', value: '--tile' },
    ],
  },
]
