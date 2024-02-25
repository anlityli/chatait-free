/*
 * Copyright 2024 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

/**
 * 钱包字段
 */
export enum WalletType {
  balance = 'balance', // 余额
  gpt3 = 'gpt3',
  gpt4 = 'gpt4',
  midjourney = 'midjourney',
}

export enum MidjourneyCreateModel {
  fast = 'fast',
  relax = 'relax',
  turbo = 'turbo',
}

export enum BaiduFeature {
  translate = 'translate',
  censor = 'censor',
}
