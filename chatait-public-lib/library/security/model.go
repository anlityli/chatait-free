// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package security

import "github.com/anlityli/chatait-free/chatait-public-lib/app/model/entity"

type ParseUserTokenRe struct {
	User      *entity.User
	IsRefresh bool
	DeviceId  string
}

type ParseAdminTokenRe struct {
	Admin     *entity.Admin
	IsRefresh bool
	DeviceId  string
}

type EmailCodeRe struct {
	Code           string
	IntervalSecond int64
	ExpireIn       int64
}

type ValidateEmailCodeRe struct {
	IsRight   bool
	IsExpired bool
}
