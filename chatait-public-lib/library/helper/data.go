// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package helper

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/dao"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/model/entity"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/api/baidu"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/api/baidu/censor"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/snowflake"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/xtime"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/grand"
	uuid "github.com/satori/go.uuid"
	"math"
	"math/big"
	"net/url"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gogf/gf/text/gstr"

	"github.com/gogf/gf/util/gconv"
	mathRand "math/rand"
)

// MergeStr 高效合并字符串
func MergeStr(strings []string) string {
	var buffer bytes.Buffer
	for _, str := range strings {
		buffer.WriteString(str)
	}
	re := buffer.String()
	return re
}

// StrInArr 字符串是否在数组中
func StrInArr(arr interface{}, str string) bool {
	re := false
	for _, v := range gconv.SliceStr(arr) {
		if v == str {
			re = true
			break
		}
	}
	return re
}

// IntInArr 整形是否在数组中
func IntInArr(arr interface{}, element int) bool {
	for _, v := range gconv.SliceInt(arr) {
		if v == element {
			return true
		}
	}
	return false
}

// TypeOf 数据类型
func TypeOf(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

// Decimal 小数点后几位的Float64
func Decimal(value float64, decimal int) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%."+gconv.String(decimal)+"f", value), 64)
	return value
}

// JSONToMap Json字符串转map,要求有索引的json,数组不可以
func JSONToMap(jsonData string) (re map[string]interface{}, err error) {
	err = json.Unmarshal([]byte(jsonData), &re)
	return re, err
}

// JSONToSlice json转数组
func JSONToSlice(jsonData string) (re []interface{}, err error) {
	err = json.Unmarshal([]byte(jsonData), &re)
	return re, err
}

// RemoveKeys 移除list中不需要的key
func RemoveKeys(list []map[string]interface{}, keys []string) []map[string]interface{} {
	for _, value := range list {
		for _, key := range keys {
			delete(value, key)
		}
	}
	return list
}

// GetMapValueByKey map是否存在字段,存在返回值,不存在返回空字符串
func GetMapValueByKey(data map[string]interface{}, key string) string {
	if _, ok := data[key]; ok {
		return gconv.String(data[key])
	}
	return ""
}

// MergeMaps 合并多个map
func MergeMaps(maps ...map[string]interface{}) (re map[string]interface{}) {
	re = make(map[string]interface{})
	for _, oneMap := range maps {
		for k, v := range oneMap {
			re[k] = v
		}
	}
	return
}

// Average 平均值
func Average(arr []float64) (avg float64) {
	sum := 0.0
	switch len(arr) {
	case 0:
		avg = 0
	default:
		for _, v := range arr {
			sum += v
		}
		avg = sum / float64(len(arr))
	}
	return
}

// ArrayToStr 数组转为符号分隔的字符串
func ArrayToStr(arr interface{}, delimiter string) (re string) {
	arrR := gconv.SliceStr(arr)
	for _, str := range arrR {
		re += delimiter + gconv.String(str)
	}
	re = gstr.SubStr(re, 1, len(re)-1)
	return
}

// StrToArray 字符串转数组
func StrToArray(str string, delimiter string) (re []string) {
	return strings.Split(str, delimiter)
}

// URLEncode url encode
func URLEncode(urlStr string) string {
	return url.QueryEscape(urlStr)
}

// URLDecode url decode
func URLDecode(urlStr string) string {
	urlDecodeStr, _ := url.QueryUnescape(urlStr)
	return urlDecodeStr
}

// SliceDataValueToString 列表数据内容转字符串
func SliceDataValueToString(data interface{}, fields ...interface{}) (listData []map[string]interface{}) {
	waitFields := make([]interface{}, 0)
	if len(fields) > 0 && reflect.TypeOf(fields[0]).Kind() == reflect.Slice {
		waitFields = fields[0].([]interface{})
	} else {
		waitFields = fields
	}
	listData = gconv.SliceMap(data)
	if len(listData) > 0 {
		for index, value := range listData {
			for i, v := range value {
				intV := gconv.Int(v)
				if intV >= math.MaxInt32 || StrInArr(waitFields, i) {
					listData[index][i] = gconv.String(v)
				}
			}
		}
	} else {
		listData = make([]map[string]interface{}, 0)
	}
	return
}

// MapDataValueToString map数据内容转字符串
func MapDataValueToString(data interface{}, fields ...interface{}) (reData map[string]interface{}) {
	waitFields := make([]interface{}, 0)
	if len(fields) > 0 && reflect.TypeOf(fields[0]).Kind() == reflect.Slice {
		waitFields = fields[0].([]interface{})
	} else {
		waitFields = fields
	}
	reData = gconv.Map(data)
	if len(reData) > 0 {
		for index, value := range reData {
			intV := gconv.Int(value)
			if intV >= math.MaxInt32 || StrInArr(waitFields, index) {
				reData[index] = gconv.String(value)
			}
		}
	}
	return
}

// GenerateRandomString 生成随机字符串
func GenerateRandomString(length int) string {
	var container string
	var str = "abcdefghijkmnpqrstuvwxyz1234567890"
	b := bytes.NewBufferString(str)
	tempLength := b.Len()
	bigInt := big.NewInt(int64(tempLength))
	for i := 0; i < length; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}

// FormatPrice 任意数值价格保留两位小数
func FormatPrice(price interface{}) string {
	priceFloat := gconv.Float64(price)
	return fmt.Sprintf("%.2f", priceFloat)
}

// CentToYuanClearZero 格式化金额去除后面的零
func CentToYuanClearZero(points interface{}) (re string) {
	formatPrice := CentToYuan(points)
	formatPriceArr := gstr.Explode(".", formatPrice)
	match, _ := regexp.MatchString("^[0-9]0$", formatPriceArr[1])
	if match {
		formatPriceArr[1] = gstr.SubStr(formatPriceArr[1], 0, len(formatPriceArr[1])-1)
		match1, _ := regexp.MatchString("^0$", formatPriceArr[1])
		if match1 {
			formatPriceArr[1] = ""
		}
	}
	if formatPriceArr[1] == "" {
		re = formatPriceArr[0]
	} else {
		re = formatPriceArr[0] + "." + formatPriceArr[1]
	}
	return re
}

// CentToYuan 整数分为单位的金额转为元
func CentToYuan(points interface{}) string {
	return FormatPrice(gconv.Float64(gconv.Float64(points) / 100))
}

// YuanToCent 100.12 转为 10012
func YuanToCent(yuan interface{}) int {
	re, err := strconv.ParseFloat(fmt.Sprintf("%0f", gconv.Float64(yuan)*100), 32)
	if err != nil {
		return 0
	}
	return gconv.Int(re)
}

// URLParamsToMap url参数转map
func URLParamsToMap(paramsStr string) (re map[string]interface{}, err error) {
	urlRe, err := url.ParseQuery(paramsStr)
	if err != nil {
		return nil, err
	}
	re = make(map[string]interface{})
	for k, v := range urlRe {
		re[k] = v[0]
	}
	return re, nil
}

// MapToURLParams map转url参数
func MapToURLParams(paramsMap map[string]interface{}) (re string) {
	for key, value := range paramsMap {
		re += key + "=" + gconv.String(value) + "&"
	}
	re = gstr.SubStr(re, 0, len(re)-1)
	return re
}

// SliceDataValueAllToString 数组中的map值全都转字符串
func SliceDataValueAllToString(data interface{}) (re interface{}) {
	if reflect.ValueOf(data).Kind() == reflect.Slice {
		listData := gconv.SliceMap(data)
		if len(listData) > 0 {
			for index, value := range listData {
				for i, v := range value {
					listData[index][i] = gconv.String(v)
				}
			}
		} else {
			listData = make([]map[string]interface{}, 0)
		}
		return listData
	} else if reflect.ValueOf(data).Kind() == reflect.Map {
		listData := gconv.Map(data)
		if len(listData) > 0 {
			for index, value := range listData {
				valueMap := gconv.Map(value)
				for i, v := range valueMap {
					valueMap[i] = gconv.String(v)
				}
				listData[index] = valueMap
			}
		} else {
			listData = make(map[string]interface{})
		}
		return listData
	}
	return nil
}

// MapKeySortToUrlParams 对map的key进行排序并生成url参数形式
func MapKeySortToUrlParams(dataMap map[string]interface{}, urlEncode bool) (re string) {
	//对map的key进行排序 首先我们将map的key存放在一个切片中
	var sslice []string
	for key, _ := range dataMap {
		sslice = append(sslice, key)
	}
	sort.Strings(sslice)
	for _, mapKey := range sslice {
		urlStr := gconv.String(dataMap[mapKey])
		if urlEncode {
			urlStr = url.QueryEscape(urlStr)
		}
		re += mapKey + "=" + urlStr + "&"
	}
	re = SubStr(re, 0, -1)
	return re
}

// ULen 在用SubStr时可以对长度进行判断
func ULen(source interface{}) int {
	str := source.(string)
	var r = []rune(str)
	return len(r)
}

// SubStr 支持中文截取, 支持end 负数截取
func SubStr(source interface{}, start int, end int) string {
	str := source.(string)
	var r = []rune(str)
	length := len(r)
	subLen := end - start

	for {
		if start < 0 {
			break
		}
		if start == 0 && subLen == length {
			break
		}
		if end > length {
			subLen = length - start
		}
		if end < 0 {
			subLen = length - start + end
		}
		var substring bytes.Buffer
		if end > 0 {
			subLen = subLen + 1
		}
		for i := start; i < subLen; i++ {
			substring.WriteString(string(r[i]))
		}
		str = substring.String()

		break
	}

	return str
}

// DelFromArr 从数组中删除元素
func DelFromArr(arr interface{}, index int) []interface{} {
	arrSlice := gconv.SliceAny(arr)
	return append(arrSlice[:index], arrSlice[index+1:]...)
}

func RandSliceValue(xs []string) string {
	return xs[grand.Intn(len(xs))]
}

// FormatFloatToString 浮点型数值转字符串保留小数
func FormatFloatToString(data interface{}, decimal ...int) string {
	decimalValue := 2
	if len(decimal) > 0 {
		decimalValue = decimal[0]
	}
	priceFloat := gconv.Float64(data)
	return fmt.Sprintf("%."+gconv.String(decimalValue)+"f", priceFloat)
}

// RandInt64 随机数
func RandInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}

	mathRand.Seed(time.Now().UnixNano())
	return mathRand.Int63n(max-min) + min
}

// RandNumUniqueInt 随机数
func RandNumUniqueInt(min, max int, num int) []int {
	var randSlice []int
	if min >= max || max == 0 || num == 0 {
		return nil
	}

	mathRand.Seed(time.Now().UnixNano())

	for len(randSlice) < num {
		//randomNumber := mathRand.Intn(max-min) + min
		randomNumber := grand.N(min, max)

		//查重
		exist := false
		for _, v := range randSlice {
			if v == randomNumber {
				exist = true
				break
			}
		}
		if !exist {
			randSlice = append(randSlice, randomNumber)
		}
	}

	return randSlice
}

// FormatIntToLength 格式化数字以前导0补齐
func FormatIntToLength(value int64, length int) string {
	return fmt.Sprintf("%0"+gconv.String(length)+"d", value)
}

const (
	// MapSortTypeKey 按照键排序
	MapSortTypeKey = 1
	// MapSortTypeValue 按照值排序
	MapSortTypeValue = 2
	// MapSortTypeField 按照二维map中的字段排序
	MapSortTypeField = 3
)

// MapSort map排序
func MapSort(mapData interface{}, sortType int, fieldName ...string) (re []interface{}) {
	mapDataMap := gconv.Map(mapData)
	if len(mapDataMap) == 0 {
		return nil
	}
	// 复制一个map
	dataMap := make(map[string]interface{})
	for k, v := range mapDataMap {
		dataMap[gconv.String(k)] = v
	}
	re = make([]interface{}, 0)
	if sortType == MapSortTypeKey {
		// 按照key排序
		var keys []int
		for k, _ := range dataMap {
			kInt := gconv.Int(k)
			keys = append(keys, kInt)
		}
		sort.Ints(keys)
		for _, k := range keys {
			re = append(re, dataMap[gconv.String(k)])
		}
	} else if sortType == MapSortTypeValue {
		// 按照value排序
		var values []int
		for _, v := range dataMap {
			values = append(values, gconv.Int(v))
		}
		sort.Ints(values)
		for _, v := range values {
			re = append(re, v)
		}
	} else if sortType == MapSortTypeField {
		// 按照 map 中的 二维map的 field 排序
		var values []int
		for _, v := range dataMap {
			vMap := gconv.Map(v)
			vInt := gconv.Int(vMap[fieldName[0]])
			values = append(values, vInt)
		}
		sort.Ints(values)
		for _, v := range values {
			for mk, mv := range dataMap {
				mvMap := gconv.Map(mv)
				if gconv.Int(mvMap[fieldName[0]]) == v {
					re = append(re, mv)
					delete(dataMap, mk)
				}
			}
		}
	}
	return re
}

// MakeRange 根据起始数，长度，步长获取数组
func MakeRange(start int, count int, step int) []int {
	s := make([]int, count)
	for i := range s {
		s[i] = start
		start += step
	}
	return s
}

// Round 四舍五入
func Round(x float64) int {
	return int(math.Floor(x + 0.5))
}

func ItemIsInSlice(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// GetRandStr 生成对应长度的随机码
func GetRandStr(n int) (randStr string) {
	chars := "ABCDEFGHIJKMNPQRSTUVWXYZabcdefghijkmnpqrstuvwxyz23456789"
	charsLen := len(chars)
	if n > 10 {
		n = 10
	}

	mathRand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		randIndex := mathRand.Intn(charsLen)
		randStr += chars[randIndex : randIndex+1]
	}
	return randStr
}

// DataToJsonStr 任意数据转换成为json格式的字符串 基本类型只转换类型 布尔值转换为1和0 其他转换为 json 字符串
func DataToJsonStr(value interface{}) (re string, err error) {
	if reflect.TypeOf(value).Kind() == reflect.String ||
		reflect.TypeOf(value).Kind() == reflect.Uint64 ||
		reflect.TypeOf(value).Kind() == reflect.Uint32 ||
		reflect.TypeOf(value).Kind() == reflect.Uint16 ||
		reflect.TypeOf(value).Kind() == reflect.Uint8 ||
		reflect.TypeOf(value).Kind() == reflect.Uint ||
		reflect.TypeOf(value).Kind() == reflect.Int64 ||
		reflect.TypeOf(value).Kind() == reflect.Int32 ||
		reflect.TypeOf(value).Kind() == reflect.Int16 ||
		reflect.TypeOf(value).Kind() == reflect.Int8 ||
		reflect.TypeOf(value).Kind() == reflect.Int ||
		reflect.TypeOf(value).Kind() == reflect.Float64 ||
		reflect.TypeOf(value).Kind() == reflect.Float32 {
		re = gconv.String(value)
	} else if reflect.TypeOf(value).Kind() == reflect.Bool {
		if gconv.Bool(value) == true {
			re = "1"
		} else {
			re = "0"
		}
	} else {
		vJSON, err := gjson.Encode(value)
		if err != nil {
			return "", err
		}
		re = gconv.String(vJSON)
	}
	return re, nil
}

// SortIntSlice Int数组正序排序
func SortIntSlice(sortSlice *[]int) {
	if len(*sortSlice) == 0 {
		return
	}
	temp := 0
	for i := 0; i < len(*sortSlice)-1; i++ {
		for j := 0; j < len(*sortSlice)-1-i; j++ {
			if (*sortSlice)[j] > (*sortSlice)[j+1] {
				temp = (*sortSlice)[j]
				(*sortSlice)[j] = (*sortSlice)[j+1]
				(*sortSlice)[j+1] = temp
			}
		}
	}
}

// GenerateUuid 生成uuid
func GenerateUuid() string {
	uuidObj := uuid.NewV4()
	return uuidObj.String()
}

type SensitiveWordsValidateParams struct {
	UserId       int64
	ValidateType int
	TopicType    int
	Content      string
}

// SensitiveWordsValidate 敏感词过滤
func SensitiveWordsValidate(params *SensitiveWordsValidateParams) (re bool, err error) {
	wordList := &[]*entity.ConfigSensitiveWord{}
	err = dao.ConfigSensitiveWord.Where("1=1").Scan(wordList)
	if err != nil {
		return false, err
	}
	re = true
	words := make([]string, 0)
	if len(*wordList) > 0 {
		for _, item := range *wordList {
			if gstr.Contains(params.Content, item.Content) {
				re = false
				words = append(words, item.Content)
			}
		}
	}
	if !re {
		sensitiveWordsValidateToData(params.UserId, params.ValidateType, params.TopicType, params.Content, words)
		return re, nil
	}
	// 百度敏感词审核
	censorRe, err := censor.Text(&baidu.CensorTextParams{
		Text: params.Content,
	})
	if err != nil {
		return false, err
	}
	if censorRe.ConclusionType != baidu.CensorTextConclusionTypePass {
		sensitiveWordsValidateToData(params.UserId, params.ValidateType, params.TopicType, params.Content, censorRe)
		return false, nil
	}
	return true, nil
}

func sensitiveWordsValidateToData(userId int64, validateType int, topicType int, content string, validateRe interface{}) {
	validateJson, err := gjson.Encode(validateRe)
	if err == nil {
		id := snowflake.GenerateID()
		_, _ = dao.UserSensitiveWord.Data(g.Map{
			"id":              id,
			"user_id":         userId,
			"type":            validateType,
			"topic_type":      topicType,
			"content":         content,
			"validate_result": gconv.String(validateJson),
			"created_at":      xtime.GetNowTime(),
		}).Insert()
	}
}
