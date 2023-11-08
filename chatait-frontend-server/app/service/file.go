// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package service

import (
	"github.com/anlityli/chatait-free/chatait-frontend-server/app/model/request"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/dao"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/model/entity"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/helper"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"io"
	"os"
)

var File = &fileService{}

type fileService struct {
}

func (s *fileService) MidjourneyImage(r *ghttp.Request) {
	requestModel := &request.FileMidjourneyImage{}
	if err := r.Parse(requestModel); err != nil {
		r.Response.Write("404")
		r.ExitAll()
		return
	}
	//userId := auth.GetUserId(r)
	fileData := &entity.FileMidjourney{}
	err := dao.FileMidjourney.Where("id", requestModel.Id).Scan(fileData)
	if err != nil {
		r.Response.Write("404")
		r.ExitAll()
		return
	}
	fileName := fileData.FileName
	//if fileData.UserId != userId {
	//	r.Response.Write("404")
	//	r.ExitAll()
	//	return
	//}
	// 获取配置是从本地保存的，则直接从本地路径获取，不是从本地保存的，从网络获取
	midjourneyShowRemoteImage, err := helper.GetConfig("midjourneyShowRemoteImage")
	if err != nil {
		r.Response.Write("配置不存在")
		r.ExitAll()
		return
	}
	if midjourneyShowRemoteImage == "1" {
		proxy := ""
		configMidjourneyData := &entity.ConfigMidjourney{}
		err = dao.ConfigMidjourney.Where("proxy<>''").Scan(configMidjourneyData)
		if err == nil {
			proxy = configMidjourneyData.Proxy
		}

		httpClient := ghttp.NewClient()
		if proxy != "" {
			httpClient.SetProxy(proxy)
		}
		response, err := httpClient.Get(fileData.MjUrl)
		if err != nil {
			r.Response.Write(err.Error())
			return
		}
		defer response.Close()
		defer response.Body.Close()
		if response.StatusCode != 200 {
			r.Response.Write(response.StatusCode)
			return
		}
		file := response.Body
		r.Response.Header().Set("Content-Type", "image/png")
		r.Response.Header().Set("Accept-Ranges", "bytes")
		r.Response.Header().Set("Expires", "0")
		r.Response.Header().Set("Cache-Control", "must-revalidate")
		r.Response.Header().Set("Pragma", "public")
		r.Response.Header().Set("Content-Length", gconv.String(response.ContentLength))
		r.Response.Header().Set("Content-Disposition", "attachment;filename="+fileName)

		buf := make([]byte, 1000)
		for {
			n, err := file.Read(buf)
			if err != nil && err != io.EOF {
				r.Response.Write(err.Error())
				break
			}
			if n == 0 {
				break
			}
			r.Response.Write(buf[:n])
		}
		r.ExitAll()
	} else {
		path := helper.FormatDirStr(g.Config().GetString("commonConf.fileSavePath")) + fileData.Path
		if requestModel.Thumbnail == 1 && fileData.Thumbnail != "" {
			path = helper.FormatDirStr(g.Config().GetString("commonConf.fileSavePath")) + fileData.Thumbnail
		}

		file, err := os.Open(path)
		if err != nil {
			r.Response.Write(err.Error())
			r.ExitAll()
			return
		}
		defer file.Close()
		fileStat, err := file.Stat()
		if err != nil {
			r.Response.Write(err.Error())
			r.ExitAll()
			return
		}
		fileSize := fileStat.Size()

		r.Response.Header().Set("Content-Type", "image/png")
		r.Response.Header().Set("Accept-Ranges", "bytes")
		r.Response.Header().Set("Expires", "0")
		r.Response.Header().Set("Cache-Control", "must-revalidate")
		r.Response.Header().Set("Pragma", "public")
		r.Response.Header().Set("Content-Length", gconv.String(fileSize))
		r.Response.Header().Set("Content-Disposition", "attachment;filename="+fileName)

		buf := make([]byte, 1000)
		for {
			n, err := file.Read(buf)
			if err != nil && err != io.EOF {
				r.Response.Write(err.Error())
				break
			}
			if n == 0 {
				break
			}
			r.Response.Write(buf[:n])
		}
		r.ExitAll()
	}
}
