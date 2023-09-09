// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package xtime

import (
	"github.com/beevik/ntp"
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
)

// GetNetworkTime 获取当前网络时间
func GetNetworkTime() (time.Time, error) {
	ntpTime, err := ntp.Time("pool.ntp.org")
	if err != nil {
		return time.Time{}, err
	}

	return ntpTime, nil
}

// GetNow 获取时间对象单例
func GetNow() *gtime.Time {
	diffTime, _ := time.ParseDuration(g.Config().GetString("commonConf.diffTime"))
	_ = gtime.SetTimeZone("Asia/Shanghai'")
	return gtime.Now().Add(diffTime)
}

// GetNowTime 当前时间戳
func GetNowTime() int64 {
	return GetNow().Timestamp()
}

// GetNowFormat 按格式获取当前时间
func GetNowFormat(format string) string {
	return GetNow().Format(format)
}

// GetTodayBegin 今天开始时间
func GetTodayBegin() int64 {
	timeStr := GetNow().Format("Y-m-d")
	//t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr+" 00:00:00", time.Local)
	t, _ := gtime.StrToTime(timeStr)
	return t.Timestamp()
}

// GetTodayEnd 今天结束时间
func GetTodayEnd() int64 {
	timeStr := GetNow().Format("Y-m-d")
	//t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr+" 00:00:00", time.Local)
	t, _ := gtime.StrToTime(timeStr + " 23:59:59")
	return t.Timestamp()
}

// GetYesterdayBegin 昨天开始时间戳
func GetYesterdayBegin() int64 {
	timeStr := GetNow().AddDate(0, 0, -1).Format("Y-m-d")
	t, _ := gtime.StrToTime(timeStr)
	return t.Timestamp()
}

// GetYesterdayEnd 昨天结束时间戳
func GetYesterdayEnd() int64 {
	timeStr := GetNow().AddDate(0, 0, -1).Format("Y-m-d")
	t, _ := gtime.StrToTime(timeStr + " 23:59:59")
	return t.Timestamp()
}

// GetWeekBegin 本周开始时间戳
func GetWeekBegin() int64 {
	timeStr := GetNow().StartOfWeek().AddDate(0, 0, 1).Format("Y-m-d")
	t, _ := gtime.StrToTime(timeStr)
	return t.Timestamp()
}

// GetWeekEnd 本周结束时间戳
func GetWeekEnd() int64 {
	timeStr := GetNow().EndOfWeek().AddDate(0, 0, 1).Format("Y-m-d")
	t, _ := gtime.StrToTime(timeStr + " 23:59:59")
	return t.Timestamp()
}

// GetLastWeekBegin 上周开始
func GetLastWeekBegin() int64 {
	timeStr := GetNow().StartOfWeek().AddDate(0, 0, -6).Format("Y-m-d")
	t, _ := gtime.StrToTime(timeStr)
	return t.Timestamp()
}

// GetLastWeekEnd 上周结束
func GetLastWeekEnd() int64 {
	timeStr := GetNow().EndOfWeek().AddDate(0, 0, -6).Format("Y-m-d")
	t, _ := gtime.StrToTime(timeStr + " 23:59:59")
	return t.Timestamp()
}

// GetMonthBegin 本月开始
func GetMonthBegin() int64 {
	timeStr := GetNow().StartOfMonth().Format("Y-m-d")
	t, _ := gtime.StrToTime(timeStr)
	return t.Timestamp()
}

// GetMonthEnd 本月结束
func GetMonthEnd() int64 {
	timeStr := GetNow().EndOfMonth().Format("Y-m-d")
	t, _ := gtime.StrToTime(timeStr + " 23:59:59")
	return t.Timestamp()
}

// GetLastMonthBegin 上月开始
func GetLastMonthBegin() int64 {
	timeStr := GetNow().StartOfMonth().AddDate(0, 0, -1).StartOfMonth().Format("Y-m-d")
	t, _ := gtime.StrToTime(timeStr)
	return t.Timestamp()
}

// GetLastMonthEnd 上月结束
func GetLastMonthEnd() int64 {
	timeStr := GetNow().StartOfMonth().AddDate(0, 0, -1).EndOfMonth().Format("Y-m-d")
	t, _ := gtime.StrToTime(timeStr + " 23:59:59")
	return t.Timestamp()
}

// GetYearMonthDay 指定时间的年月日
func GetYearMonthDay(t *gtime.Time, yearFormat string, monthFormat string, dayFormat string) (string, string, string) {
	nowYear := t.Format(yearFormat)
	nowMonth := t.Format(monthFormat)
	nowDay := t.Format(dayFormat)
	return nowYear, nowMonth, nowDay
}

// GetMonthStartEnd 获取指定时间所在月的开始 结束时间
func GetMonthStartEnd(t time.Time) (time.Time, time.Time) {
	monthStartDay := t.AddDate(0, 0, -t.Day()+1)
	monthStartTime := time.Date(monthStartDay.Year(), monthStartDay.Month(), monthStartDay.Day(), 0, 0, 0, 0, t.Location())
	monthEndDay := monthStartTime.AddDate(0, 1, -1)
	monthEndTime := time.Date(monthEndDay.Year(), monthEndDay.Month(), monthEndDay.Day(), 23, 59, 59, 0, t.Location())
	return monthStartTime, monthEndTime
}

// GetLastMonthStartEnd 指定时间的上一个月的开始、结束时间
func GetLastMonthStartEnd(t time.Time) (time.Time, time.Time) {
	monthStartDay := t.AddDate(0, -1, 0)
	monthStartTime := time.Date(monthStartDay.Year(), monthStartDay.Month(), 1, 0, 0, 0, 0, t.Location())
	monthEndDay := monthStartTime.AddDate(0, 1, -1)
	monthEndTime := time.Date(monthEndDay.Year(), monthEndDay.Month(), monthEndDay.Day(), 23, 59, 59, 0, t.Location())
	return monthStartTime, monthEndTime
}

// GetTimeStampByTimeStr GetTimeStampByDateStr 根据日期字符串获取时间戳
func GetTimeStampByTimeStr(timeStr string) int64 {
	//t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr+" 00:00:00", time.Local)
	t, _ := gtime.StrToTime(timeStr)
	return t.Timestamp()
}

// GetAllDateFromStartToEnd 从给定的开始时间和结束时间获取所有日期
func GetAllDateFromStartToEnd(startDate string, endDate string, format ...string) (re []string) {
	re = make([]string, 0)
	formatStr := "Y-m-d"
	if len(format) > 0 {
		formatStr = format[0]
	}
	startTime := gtime.NewFromStrFormat(startDate, formatStr).Timestamp()
	endTime := gtime.NewFromStrFormat(endDate, formatStr).Timestamp()
	for theTime := startTime; theTime <= endTime; theTime += 86400 {
		theTimeObj := gtime.NewFromTimeStamp(theTime)
		theData := theTimeObj.Format(formatStr)
		re = append(re, theData)
	}
	return re
}

// GetAssignDayBegin 指定日期当天的开始时间
func GetAssignDayBegin(assignDay string) *gtime.Time {
	t, _ := gtime.StrToTime(assignDay)
	return t
}

// GetAssignDayEnd 指定日期当天的结束时间
func GetAssignDayEnd(assignDay string) *gtime.Time {
	t, _ := gtime.StrToTime(assignDay)
	nt := gtime.NewFromTimeStamp(t.AddDate(0, 0, 1).Timestamp() - 1)
	return nt
}
