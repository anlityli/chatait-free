// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package security

import (
	"database/sql"
	"errors"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/dao"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/model/entity"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/helper"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/snowflake"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/xtime"
	"github.com/dgrijalva/jwt-go"
	"github.com/gogf/gf/crypto/gsha1"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/encoding/gbase64"
	"github.com/gogf/gf/encoding/gbinary"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtimer"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/grand"
	"time"
)

// Key 加密密钥
const Key string = "zAr92gmhb9fujgsYgk9Wf6trkfxneuaNG"

const (
	ScenarioSignup       = "signup"
	ScenarioFindPassword = "findPassword"
)

// GenerateUserAccessToken 生成Token
func GenerateUserAccessToken(user *entity.User) (tokenStr string, exp int64, expIn int64, err error) {
	expTime, _ := time.ParseDuration(g.Config().GetString("frontendConf.accessTokenExp"))
	return generateUserToken(user, expTime, false)
}

func GenerateAdminAccessToken(admin *entity.Admin) (tokenStr string, exp int64, expIn int64, err error) {
	expTime, _ := time.ParseDuration(g.Config().GetString("backendConf.accessTokenExp"))
	return generateAdminToken(admin, expTime, false)
}

func GenerateUserRefreshToken(user *entity.User) (tokenStr string, exp int64, expIn int64, err error) {
	expTime, _ := time.ParseDuration(g.Config().GetString("frontendConf.refreshTokenExp"))
	return generateUserToken(user, expTime, true)
}

func GenerateAdminRefreshToken(admin *entity.Admin) (tokenStr string, exp int64, expIn int64, err error) {
	expTime, _ := time.ParseDuration(g.Config().GetString("backendConf.refreshTokenExp"))
	return generateAdminToken(admin, expTime, true)
}

// ValidateParamsSign 校验sign是否正确
func ValidateParamsSign(dataMap map[string]interface{}, privateKey string) bool {
	if _, ok := dataMap["sign"]; !ok {
		return false
	}
	sign := gconv.String(dataMap["sign"])
	validateSign := GenerateParamsSign(dataMap, privateKey)
	return sign == validateSign
}

// GenerateParamsSign 生成验签sign
func GenerateParamsSign(dataMap map[string]interface{}, privateKey string) (sign string) {
	// 先把参数按照中的sign删掉
	delete(dataMap, "sign")
	// map当中加入私钥
	dataMap["private_key"] = privateKey
	// 把data按照字典排序并转换成url字符串形式
	urlData := helper.MapKeySortToUrlParams(dataMap, false)
	//glog.Line().Debug(urlData)
	urlData = gbase64.EncodeString(urlData)
	// sha1加密urlData
	sign = gsha1.Encrypt(urlData)
	return
}

// ParseUserToken 解析Token
func ParseUserToken(tokenStr string) (re *ParseUserTokenRe, err error) {
	re = &ParseUserTokenRe{}
	re.User = &entity.User{}
	token, err := jwt.Parse(tokenStr, secret())
	if err != nil {
		return
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = errors.New("cannot convert claim to mapclaim")
		return
	}
	//验证token，如果token被修改过则为false
	if !token.Valid {
		err = errors.New("token is invalid")
		return
	}
	// 登录token 在强制重新登录时间之前，则要求重新登录
	forcedReLoginTime := g.Config().GetInt64("frontendConf.forcedReLoginTime")
	if gconv.Int64(claim["iat"]) < forcedReLoginTime {
		err = errors.New("token is forced invalid")
		return
	}
	re.User.Id = gconv.Int64(claim["id"])
	re.IsRefresh = gconv.Bool(claim["refresh"])
	return
}

// ParseAdminToken 解析管理员Token
func ParseAdminToken(tokenStr string) (re *ParseAdminTokenRe, err error) {
	re = &ParseAdminTokenRe{}
	re.Admin = &entity.Admin{}
	token, err := jwt.Parse(tokenStr, secret())
	if err != nil {
		return
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = errors.New("cannot convert claim to mapclaim")
		return
	}
	//验证token，如果token被修改过则为false
	if !token.Valid {
		err = errors.New("token is invalid")
		return
	}
	// 登录token 在强制重新登录时间之前，则要求重新登录
	forcedReLoginTime := g.Config().GetInt64("backendConf.forcedReLoginTime")
	if gconv.Int64(claim["iat"]) < forcedReLoginTime {
		err = errors.New("token is forced invalid")
		return
	}
	re.Admin.Id = gconv.Int64(claim["id"])
	re.Admin.AdminName = gconv.String(claim["adminName"])
	re.Admin.UserId = gconv.Int64(claim["userId"])
	re.Admin.RoleId = gconv.Int64(claim["roleId"])
	re.IsRefresh = gconv.Bool(claim["refresh"])
	return
}

// generateUserToken 生成会员Token
func generateUserToken(user *entity.User, expTime time.Duration, isRefresh bool) (tokenStr string, exp int64, expIn int64, err error) {
	claim := jwt.MapClaims{
		"id":      gconv.String(user.Id),
		"exp":     xtime.GetNow().Add(expTime).Unix(),
		"iat":     xtime.GetNow().Unix(),
		"refresh": isRefresh,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenStr, err = token.SignedString(getSecurityKeyByte())
	exp = gconv.Int64(claim["exp"])
	expIn = gconv.Int64(expTime.Seconds())
	return
}

// generateAdminToken 生成adminToken
func generateAdminToken(admin *entity.Admin, expTime time.Duration, isRefresh bool) (tokenStr string, exp int64, expIn int64, err error) {
	claim := jwt.MapClaims{
		"id":        gconv.String(admin.Id),
		"adminName": admin.AdminName,
		"userId":    gconv.String(admin.UserId),
		"roleId":    gconv.String(admin.RoleId),
		"exp":       xtime.GetNow().Add(expTime).Unix(),
		"iat":       xtime.GetNow().Unix(),
		"refresh":   isRefresh,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenStr, err = token.SignedString(getSecurityKeyByte())
	exp = gconv.Int64(claim["exp"])
	expIn = gconv.Int64(expTime.Seconds())
	return
}

// secret jwt加密密钥回调方法
func secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return getSecurityKeyByte(), nil
	}
}

// GenerateEmailCode 生成4位验证码
// 传入场景不同的场景不同的验证码互不干扰
func GenerateEmailCode(scenario string, email string, ip string) (re *EmailCodeRe, err error) {
	nowTime := xtime.GetNowTime()
	intervalSecond := gconv.Int64(g.Config().GetString("commonConf.emailCodeIntervalSecond"))
	ipData := &entity.EmailCode{}
	err = dao.EmailCode.Where("ip=?", ip).Scan(ipData)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if ipData.Id > 0 && nowTime-gconv.Int64(ipData.UpdatedAt) <= intervalSecond {
		return nil, errors.New("间隔时间太短，请稍后再发送")
	}
	emailData := &entity.EmailCode{}
	err = dao.EmailCode.Where("email=?", email).Scan(emailData)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if emailData.Id > 0 && nowTime-gconv.Int64(ipData.UpdatedAt) <= intervalSecond {
		return nil, errors.New("间隔时间太短，请稍后再发送")
	}
	code := grand.Digits(4)
	updateData := g.Map{
		"email":    email,
		"ip":       ip,
		"scenario": scenario,
		"code":     code,
	}
	if emailData.Id <= 0 {
		id := snowflake.GenerateID()
		emailData.Id = id
		updateData["id"] = id
		updateData["created_at"] = nowTime
		updateData["updated_at"] = nowTime
		if _, err = dao.EmailCode.Data(updateData).Insert(); err != nil {
			return nil, err
		}
	} else {
		updateData["updated_at"] = nowTime
		if _, err = dao.EmailCode.Data(updateData).Where("id=?", emailData.Id).Update(); err != nil {
			return nil, err
		}
	}
	// 600秒后删除该记录
	emailCodeExpireIn := gconv.Int64(g.Config().GetString("commonConf.emailCodeExpireIn"))
	gtimer.AddOnce(time.Second*time.Duration(emailCodeExpireIn), func() {
		_, _ = dao.EmailCode.Where("id=?", emailData.Id).Delete()
	})
	re = &EmailCodeRe{
		Code:           code,
		IntervalSecond: intervalSecond,
		ExpireIn:       emailCodeExpireIn,
	}
	return re, nil
}

// ValidateEmailCode 校验验证码
func ValidateEmailCode(scenario string, email string, code string) (re *ValidateEmailCodeRe, err error) {
	re = &ValidateEmailCodeRe{}
	emailData := &entity.EmailCode{}
	err = dao.EmailCode.Where("email=? AND scenario=?", email, scenario).Scan(emailData)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if emailData.Id <= 0 {
		re.IsExpired = true
		return re, nil
	}
	// 如果校验次数大于3次直接返回失败，删除当前记录，防止刷码
	if emailData.ValidateTimes >= 3 {
		_, _ = dao.EmailCode.Where("id=?", emailData.Id).Delete(emailData)
		return nil, errors.New("检验次数过多，请重新获取")
	}
	if _, err = dao.EmailCode.Data(g.Map{
		"validate_times": gdb.Raw("validate_times+1"),
	}).Where("id=?", emailData.Id).Update(emailData); err != nil {
		return nil, err
	}
	re.IsRight = emailData.Code == code
	return re, nil
}

// ValidateEmailCodeIsRight 校验是否正确 不管什么错，直接返回错
func ValidateEmailCodeIsRight(scenario string, email string, code string) (err error) {
	validateRe, err := ValidateEmailCode(scenario, email, code)
	if err != nil {
		return err
	}
	if !validateRe.IsRight {
		return errors.New("验证码不正确")
	}
	if validateRe.IsExpired {
		return errors.New("验证码已失效")
	}
	return nil
}

// getSecurityKeyByte 获取加密密钥
func getSecurityKeyByte() []byte {
	return gbinary.EncodeString(Key)
}
