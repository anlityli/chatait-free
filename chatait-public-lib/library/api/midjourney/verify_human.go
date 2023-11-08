// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package midjourney

import (
	"errors"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/model/entity"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/file"
	"github.com/TannerKvarfordt/hfapigo"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/text/gstr"
)

type VerifyHuman struct {
	Config *entity.ConfigMidjourney
}

func NewVerifyHuman(config *entity.ConfigMidjourney) (v *VerifyHuman, err error) {
	if config.HuggingFaceToken == "" {
		return nil, errors.New("HuggingFaceToken is required")
	}
	v = &VerifyHuman{}
	v.Config = config
	return v, nil
}

func (v *VerifyHuman) Verify(imageUrl string, categories []string) (re string, err error) {
	glog.Line(true).Debug("verifyStart", imageUrl, categories)
	// 把远程图片保存到本地
	fileData, err := file.RemoteFileSave(imageUrl, v.Config.HuggingFaceToken)
	if err != nil {
		return "", err
	}
	imageCategories, err := hfapigo.SendImageClassificationRequest(hfapigo.RecommendedImageClassificationModel, fileData.SavePath)
	if err != nil {
		return "", err
	}
	for _, item := range imageCategories {
		label := item.Label
		for _, category := range categories {
			if gstr.Contains(label, category) {
				return category, nil
			}
		}
	}
	return "", nil
}
