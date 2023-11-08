// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package midjourney

import (
	"github.com/gogf/gf/util/gconv"
	"regexp"
	"strings"
)

// msgContentHandler 消息处理 把****中间的内容匹配出来
func msgContentHandler(messageContent string) string {
	rep := regexp.MustCompile(`\*\*(.*?)\*\*`)
	match := rep.FindStringSubmatch(messageContent)
	if len(match) > 1 {
		return match[1]
	} else {
		return ""
	}
}

// matchMsgContentProgress 处理消息中的百分比
func matchMsgContentProgress(messageContent string) int {
	rep := regexp.MustCompile(`\((\d+)%\)`)
	match := rep.FindStringSubmatch(messageContent)
	if len(match) > 1 {
		return gconv.Int(match[1])
	} else {
		return 0
	}
}

// matchMsgContentIndex 匹配到内容
func matchMsgContentIndex(messageContent string) int {
	// 使用正则表达式查找 Image # 后面的数字
	rep := regexp.MustCompile(`Image #(\d+)`)
	match := rep.FindStringSubmatch(messageContent)

	if len(match) > 1 {
		return gconv.Int(match[1])
	} else {
		return 0
	}
}

// getMessageHash 获取messageHash值
func getMessageHash(fileName string) string {
	splitStr := strings.Split(fileName, "_")
	lastPart := splitStr[len(splitStr)-1]
	fileParts := strings.Split(lastPart, ".")
	return fileParts[0]
}

// isUpscaleString 用于区分Upscale返回的内容中的Image
func isUpscaleString(str string) bool {
	// 定义判断 "Image" 的正则表达式
	re := regexp.MustCompile(`- Image #\d+`)
	return re.MatchString(str)
}

// isVariationsString 用于区分Variation返回的内容中的Variations
func isVariationsString(str string) bool {
	// 定义判断 "Variations" 的正则表达式
	re := regexp.MustCompile(`- Variations \(.*\) by`)
	return re.MatchString(str)
}

// isPanString 用于区分Pan返回的内容中的 Pan
func isPanString(str string) bool {
	re := regexp.MustCompile(`- Pan \w+ by`)
	return re.MatchString(str)
}

func isZoomOutString(str string) bool {
	re := regexp.MustCompile(`- Zoom Out by`)
	return re.MatchString(str)
}

func matchMsgContentPan(messageContent string) (pan string, index int) {
	// 使用正则表达式查找 Image # 后面的数字
	rep := regexp.MustCompile(`- Pan (\w+)`)
	match := rep.FindStringSubmatch(messageContent)

	pan = ""
	if len(match) > 1 {
		pan = match[1]
	}
	switch pan {
	case "Left":
		return pan, 1
	case "Right":
		return pan, 2
	case "Up":
		return pan, 3
	case "Down":
		return pan, 4
	}
	return pan, 0
}
