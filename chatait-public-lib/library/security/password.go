// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package security

import (
	"github.com/gogf/gf/encoding/gbinary"
	"github.com/gogf/gf/util/gconv"
	"golang.org/x/crypto/bcrypt"
)

// GeneratePassword 生成密码
func GeneratePassword(password string) (re string, err error) {
	passwordByte := gbinary.EncodeString(password)
	passwordHash, err := bcrypt.GenerateFromPassword(passwordByte, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return gconv.String(passwordHash), err
}

// ValidatePassword 校验密码
func ValidatePassword(password, passwordHash string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password)); err != nil {
		return false
	}
	return true
}
