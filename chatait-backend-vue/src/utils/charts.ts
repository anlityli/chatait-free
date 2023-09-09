/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

import dayjs from 'dayjs'

/**
 * 获取表头数据
 *
 * @export
 * @param {string[]} dateTime
 * @param {number} divideNum
 * @returns {string[]}
 */
export function getDateArray(dateTime: string[] = [], divideNum = 10): string[] {
  const timeArray: string[] = []
  if (dateTime.length > 0) {
    for (let i = 0; i < divideNum; i++) {
      const dateAbsTime: number = (new Date(dateTime[1]).getTime() - new Date(dateTime[0]).getTime()) / divideNum
      const enhandTime: number = new Date(dateTime[0]).getTime() + dateAbsTime * i
      timeArray.push(dayjs(enhandTime).format('YYYY-MM-DD'))
    }
  }

  return timeArray
}

/**
 * 获取随机数
 *
 * @param {number} [num=100]
 * @returns
 *
 * @memberOf DashboardBase
 */
export function getRandomArray(num = 100): number {
  let resultNum = Number((Math.random() * num).toFixed(0))

  if (resultNum <= 1) {
    resultNum = 1
  }

  return resultNum
}
