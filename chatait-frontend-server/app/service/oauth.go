// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/anlityli/chatait-free/chatait-frontend-server/app/model/request"
	"github.com/anlityli/chatait-free/chatait-frontend-server/app/model/response"
	"github.com/anlityli/chatait-free/chatait-frontend-server/library/auth"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/constant"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/dao"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/libservice"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/model/entity"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/helper"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/mail"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/security"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/snowflake"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/web"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/xtime"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/grand"
)

var Oauth = &oauthService{}

type oauthService struct {
}

func (s *oauthService) SignupSendCode(r *ghttp.Request) (re *response.OauthSignUpSendCode, err error) {
	requestModel := &request.OauthSignupSendCode{}
	if err := r.Parse(requestModel); err != nil {
		return nil, err
	}
	if err := s.validateEmail(requestModel.Username); err != nil {
		return nil, err
	}
	ip := web.GetClientIP(r)
	// 生成验证码
	codeRe, err := security.GenerateEmailCode(security.ScenarioSignup, requestModel.Username, ip)
	if err != nil {
		return nil, err
	}
	// 发送短信验证码
	// 从配置文件中获取是否真的发送短信验证码如果不真的发送，则验证码通过接口返回
	if g.Config().GetBool("frontendConf.isRealSendEmailCode") {
		err = mail.SendMail(&mail.SendMailParams{
			ReceiverEmails: []string{requestModel.Username},
			MsgSubject:     fmt.Sprintf("ChatAIT 注册验证码 %s", xtime.GetNow().Format("Y-m-d H:i:s")),
			MsgBody:        fmt.Sprintf(`<h3>ChatAIT 注册验证码</h3><p>您好！</p><p>您正在注册ChatAIT，验证码为：<span style="color:#ff0000;">%s</span>，请妥善保管，不要泄露给其他人。%d分钟内有效，请尽快使用。</p>`, codeRe.Code, codeRe.IntervalSecond/60),
		})
		if err != nil {
			glog.Line().Println("验证码发送失败", requestModel.Username, err)
			return nil, errors.New("验证码发送失败")
		}
		codeRe.Code = ""
	} else {
		glog.Line().Debug(codeRe.Code)
	}
	re = &response.OauthSignUpSendCode{}
	re.Email = requestModel.Username
	re.Code = codeRe.Code
	re.IntervalSecond = gconv.Int(codeRe.IntervalSecond)
	re.ExpireIn = gconv.Int(codeRe.ExpireIn)
	return
}

// SignupValidate 校验短信验证码
func (s *oauthService) SignupValidate(r *ghttp.Request) (re *response.OauthSignUpValidateCode, err error) {
	requestModel := &request.OauthSignUpValidateCode{}
	if err := r.Parse(requestModel); err != nil {
		return nil, err
	}
	if err := s.validateEmail(requestModel.Username); err != nil {
		return nil, err
	}
	validateRe, err := security.ValidateEmailCode(security.ScenarioSignup, requestModel.Username, requestModel.Code)
	if err != nil {
		return nil, errors.New("验证码校验失败")
	}
	re = &response.OauthSignUpValidateCode{}
	re.Email = requestModel.Username
	if validateRe.IsRight {
		re.IsRight = 1
	}
	if validateRe.IsExpired {
		re.CodeExpired = 1
	}
	return re, nil
}

// SignupFinish 注册完成
func (s *oauthService) SignupFinish(r *ghttp.Request) (re *response.OauthUserToken, err error) {
	requestModel := &request.OauthSignupFinish{}
	if err = r.Parse(requestModel); err != nil {
		return nil, err
	}
	err = s.validateEmail(requestModel.Username)
	if err != nil {
		return nil, err
	}

	emailCodeEnable, err := helper.GetConfig("emailCodeEnable")
	if err != nil {
		return nil, err
	}
	if emailCodeEnable == "1" {
		if requestModel.Code == "" {
			return nil, errors.New("验证码必填")
		}
		if err := security.ValidateEmailCodeIsRight(security.ScenarioSignup, requestModel.Username, requestModel.Code); err != nil {
			return nil, err
		}
	}

	passwordHash, err := security.GeneratePassword(requestModel.Password)
	if err != nil {
		return nil, err
	}
	var userModel *entity.User
	if err := g.DB().Transaction(context.TODO(), func(ctx context.Context, tx *gdb.TX) (err error) {
		userModel, err = s.AddUser(ctx, tx, &AddUserParam{
			Username: requestModel.Username,
			Password: passwordHash,
			Nickname: requestModel.Nickname,
			Avatar:   "",
		})
		if err != nil {
			return err
		}
		// 更新最后登陆时间
		if _, err = dao.User.Ctx(ctx).TX(tx).Data(g.Map{
			"last_login_at": xtime.GetNowTime(),
		}).Where("id=?", userModel.Id).Update(); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	// 生成token
	return s.generateToken(r, userModel)
}

// Login 登录
func (s *oauthService) Login(r *ghttp.Request) (re *response.OauthUserToken, err error) {
	requestModel := &request.OauthLogin{}
	if err := r.Parse(requestModel); err != nil {
		return nil, err
	}
	userModel := &entity.User{}
	if err := dao.User.Where("username=?", requestModel.Username).Scan(userModel); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("用户名或密码错误")
		}
		return nil, err
	}
	if userModel.IsBan == 1 {
		return nil, errors.New("用户名或密码错误")
	}
	if !security.ValidatePassword(requestModel.Password, userModel.Password) {
		return nil, errors.New("用户名或密码错误")
	}
	// 更新最后登陆时间
	if _, err = dao.User.Data(g.Map{
		"last_login_at": xtime.GetNowTime(),
	}).Where("id=?", userModel.Id).Update(); err != nil {
		return nil, err
	}
	return s.generateToken(r, userModel)
}

// RefreshToken 刷新access_token
func (s *oauthService) RefreshToken(r *ghttp.Request) (re *response.OauthUserToken, err error) {
	userModel, err := auth.ValidateRefresh(r)
	if err != nil {
		return nil, err
	}
	// 生成 accessToken
	tokenStr, _, expIn, err := security.GenerateUserAccessToken(userModel)
	if err != nil {
		return nil, errors.New("token生成失败")
	}
	// 生成 refreshToken
	refreshToken, _, refreshExpIn, err := security.GenerateUserRefreshToken(userModel)
	if err != nil {
		return nil, errors.New("refresh_token生成失败")
	}
	re = &response.OauthUserToken{}
	re.AccessToken = tokenStr
	re.AccessTokenExpireIn = gconv.Int(expIn)
	re.RefreshToken = refreshToken
	re.RefreshTokenExpireIn = gconv.Int(refreshExpIn)
	return re, nil
}

// AddUserParam 添加会员参数
type AddUserParam struct {
	Username string
	Password string
	Nickname string
	Avatar   string
}

// AddUser 添加会员
func (s *oauthService) AddUser(ctx context.Context, tx *gdb.TX, addUserParam *AddUserParam) (re *entity.User, err error) {
	nowTime := xtime.GetNowTime()
	// 获取一个会员ID
	userID := snowflake.GenerateID()
	// 插入用户数据
	insertData := g.Map{
		"id":         userID,
		"username":   addUserParam.Username,
		"password":   addUserParam.Password,
		"created_at": nowTime,
	}
	if _, err = dao.User.Ctx(ctx).TX(tx).Data(insertData).Insert(); err != nil {
		return nil, err
	}
	nickname := addUserParam.Nickname
	if nickname == "" {
		nickname = addUserParam.Username
	}
	avatar := addUserParam.Avatar
	if avatar == "" {
		// 随机分配头像
		randHeadNum := grand.N(1, 20)
		if randHeadNum <= 0 {
			return nil, errors.New("头像保存失败")
		}
		avatar = "/avatar/" + gconv.String(randHeadNum) + ".png"
	}

	insertUserData := g.Map{
		"user_id":  userID,
		"nickname": nickname,
		"avatar":   avatar,
	}
	if _, err = dao.UserInfo.Ctx(ctx).TX(tx).Data(insertUserData).Insert(); err != nil {
		return nil, err
	}
	// 增加钱包数据
	if _, err = dao.Wallet.TX(tx).Ctx(ctx).Data(g.Map{
		"user_id": userID,
	}).Insert(); err != nil {
		return nil, err
	}
	newUserAddBalance, err := helper.GetConfig("newUserAddBalance")
	newUserAddGpt3, err := helper.GetConfig("newUserAddGpt3")
	if err != nil {
		return nil, err
	}
	if gconv.Int(newUserAddBalance) > 0 {
		// 赠送提问次数
		err = libservice.Wallet.ChangeWalletBalance(ctx, tx, &libservice.ChangeWalletParam{
			UserId:     userID,
			WalletType: constant.WalletTypeBalance,
			Amount:     gconv.Int(newUserAddBalance),
			Remark:     "注册免费赠送",
			TargetType: constant.WalletChangeTargetTypeAddUser,
			TargetID:   userID,
		})
		if err != nil {
			return nil, err
		}
	}
	if gconv.Int(newUserAddGpt3) > 0 {
		// 赠送提问次数
		err = libservice.Wallet.ChangeWalletBalance(ctx, tx, &libservice.ChangeWalletParam{
			UserId:     userID,
			WalletType: constant.WalletTypeGpt3,
			Amount:     gconv.Int(newUserAddGpt3),
			Remark:     "注册免费赠送",
			TargetType: constant.WalletChangeTargetTypeAddUser,
			TargetID:   userID,
		})
		if err != nil {
			return nil, err
		}
	}

	re = &entity.User{}
	err = gconv.Scan(insertData, re)
	if err != nil {
		return nil, err
	}
	return re, nil
}

// FindPasswordSendCode 找回密码发送验证啊吗
func (s *oauthService) FindPasswordSendCode(r *ghttp.Request) (re *response.OauthSignUpSendCode, err error) {
	requestModel := &request.OauthSignupSendCode{}
	if err := r.Parse(requestModel); err != nil {
		return nil, err
	}
	if err := s.validateEmailFind(requestModel.Username); err != nil {
		return nil, err
	}
	ip := web.GetClientIP(r)
	// 生成验证码
	codeRe, err := security.GenerateEmailCode(security.ScenarioFindPassword, requestModel.Username, ip)
	if err != nil {
		return nil, err
	}
	// 发送短信验证码
	// 从配置文件中获取是否真的发送短信验证码如果不真的发送，则验证码通过接口返回
	if g.Config().GetBool("frontendConf.isRealSendEmailCode") {
		err = mail.SendMail(&mail.SendMailParams{
			ReceiverEmails: []string{requestModel.Username},
			MsgSubject:     fmt.Sprintf("ChatAIT 找回密码验证码 %s", xtime.GetNow().Format("Y-m-d H:i:s")),
			MsgBody:        fmt.Sprintf(`<h3>ChatAIT 找回密码验证码</h3><p>您好！</p><p>您正在找回密码，验证码为：<span style="color:#ff0000;">%s</span>，请妥善保管，不要泄露给其他人。%d分钟内有效，请尽快使用。如果不是您本人操作，请忽略本条消息。</p>`, codeRe.Code, codeRe.IntervalSecond/60),
		})
		if err != nil {
			glog.Line().Println("验证码发送失败", requestModel.Username, err)
			return nil, errors.New("验证码发送失败")
		}
		codeRe.Code = ""
	} else {
		glog.Line().Debug(codeRe.Code)
	}
	re = &response.OauthSignUpSendCode{}
	re.Email = requestModel.Username
	re.Code = codeRe.Code
	re.IntervalSecond = gconv.Int(codeRe.IntervalSecond)
	re.ExpireIn = gconv.Int(codeRe.ExpireIn)
	return
}

// FindPasswordValidate 校验短信验证码
func (s *oauthService) FindPasswordValidate(r *ghttp.Request) (re *response.OauthSignUpValidateCode, err error) {
	requestModel := &request.OauthSignUpValidateCode{}
	if err := r.Parse(requestModel); err != nil {
		return nil, err
	}
	if err := s.validateEmailFind(requestModel.Username); err != nil {
		return nil, err
	}
	validateRe, err := security.ValidateEmailCode(security.ScenarioFindPassword, requestModel.Username, requestModel.Code)
	if err != nil {
		return nil, errors.New("验证码校验失败")
	}
	re = &response.OauthSignUpValidateCode{}
	re.Email = requestModel.Username
	if validateRe.IsRight {
		re.IsRight = 1
	}
	if validateRe.IsExpired {
		re.CodeExpired = 1
	}
	return re, nil
}

// FindPasswordFinish 找回完成
func (s *oauthService) FindPasswordFinish(r *ghttp.Request) (re *response.OauthUserToken, err error) {
	requestModel := &request.OauthFindPasswordFinish{}
	if err = r.Parse(requestModel); err != nil {
		return nil, err
	}
	err = s.validateEmailFind(requestModel.Username)
	if err != nil {
		return nil, err
	}

	if err := security.ValidateEmailCodeIsRight(security.ScenarioFindPassword, requestModel.Username, requestModel.Code); err != nil {
		return nil, err
	}

	passwordHash, err := security.GeneratePassword(requestModel.Password)
	if err != nil {
		return nil, err
	}
	userModel := &entity.User{}
	err = dao.User.Where("username=?", requestModel.Username).Scan(userModel)
	if err != nil {
		return nil, err
	}
	if err := g.DB().Transaction(context.TODO(), func(ctx context.Context, tx *gdb.TX) (err error) {
		if _, err = dao.User.Ctx(ctx).TX(tx).Data(g.Map{
			"password": passwordHash,
		}).Where("username=?", requestModel.Username).Update(); err != nil {
			return err
		}
		if err != nil {
			return err
		}
		// 更新最后登陆时间
		if _, err = dao.User.Ctx(ctx).TX(tx).Data(g.Map{
			"last_login_at": xtime.GetNowTime(),
		}).Where("id=?", userModel.Id).Update(); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	// 生成token
	return s.generateToken(r, userModel)
}

// generateToken 生成token
func (s *oauthService) generateToken(r *ghttp.Request, userModel *entity.User) (re *response.OauthUserToken, err error) {
	var tokenStr string
	var expIn int64
	var refreshToken string
	var refreshExpIn int64
	// 生成 accessToken
	tokenStr, _, expIn, err = security.GenerateUserAccessToken(userModel)
	if err != nil {
		return nil, errors.New("token生成失败")
	}
	// 生成 refreshToken
	refreshToken, _, refreshExpIn, err = security.GenerateUserRefreshToken(userModel)
	if err != nil {
		return nil, errors.New("refresh_token生成失败")
	}
	re = &response.OauthUserToken{}
	re.AccessToken = tokenStr
	re.AccessTokenExpireIn = gconv.Int(expIn)
	re.RefreshToken = refreshToken
	re.RefreshTokenExpireIn = gconv.Int(refreshExpIn)
	return re, nil
}

func (s *oauthService) validateEmail(email string) (err error) {
	// 校验用户是否存在
	data, err := dao.User.Where(dao.User.Columns.Username+"=?", email).One()
	if err != nil {
		return err
	}
	if !data.IsEmpty() {
		return errors.New("该邮箱已经被注册，不能重复注册")
	}
	return err
}

func (s *oauthService) validateEmailFind(email string) (err error) {
	// 校验用户是否存在
	data, err := dao.User.Where(dao.User.Columns.Username+"=?", email).One()
	if err != nil {
		return err
	}
	if data.IsEmpty() {
		return errors.New("该账号不存在")
	}
	return err
}
