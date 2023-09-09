// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package file

import (
	"github.com/anlityli/chatait-free/chatait-public-lib/library/helper"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/snowflake"
	"github.com/gogf/gf/crypto/gsha1"
	"github.com/gogf/gf/encoding/gurl"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"io"
	"os"
	"os/exec"
)

func RemoteFileSave(url string, proxy ...string) (re *RemoteFileSaveResult, err error) {
	urlParse, err := gurl.ParseURL(url, 32)
	if err != nil {
		return nil, err
	}
	fileExt := gfile.Ext(urlParse["path"])
	oriFileName := gfile.Name(urlParse["path"])
	// 首先储存临时文件
	httpClient := ghttp.NewClient()
	if len(proxy) > 0 && proxy[0] != "" {
		httpClient.SetProxy(proxy[0])
	}
	response, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Close()
	defer response.Body.Close()
	id := snowflake.GenerateID()
	tmpFileName := gconv.String(id) + fileExt
	tmpDir := helper.FormatDirStr(g.Config().GetString("commonConf.fileTmpPath"))
	if !gfile.IsDir(tmpDir) {
		if err := gfile.Mkdir(tmpDir); err != nil {
			return nil, err
		}
	}
	tmpPath := tmpDir + tmpFileName
	file, err := os.Create(tmpPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return nil, err
	}
	// 对临时文件进行处理移动到目标目录
	sha1, err := gsha1.EncryptFile(tmpPath)
	if err != nil {
		return nil, err
	}
	fileSize := gfile.Size(tmpPath)
	// 先查看本地目录路径是否存在，不存在则创建
	savePath := helper.FormatDirStr(g.Config().GetString("commonConf.fileSavePath"))
	if !gfile.IsDir(savePath) {
		if err := gfile.Mkdir(savePath); err != nil {
			return nil, err
		}
	}
	relativePath := ""
	pathName1 := gstr.SubStr(sha1, 0, 3)
	pathName2 := gstr.SubStr(sha1, 3, 3)
	savePath += pathName1 + "/"
	relativePath += pathName1 + "/"
	if !gfile.IsDir(savePath) {
		if err := gfile.Mkdir(savePath); err != nil {
			return nil, err
		}
	}
	savePath += pathName2 + "/"
	relativePath += pathName2 + "/"
	if !gfile.IsDir(savePath) {
		if err := gfile.Mkdir(savePath); err != nil {
			return nil, err
		}
	}
	fileName := sha1 + fileExt
	savePath += fileName
	relativePath += fileName
	if !gfile.IsFile(savePath) {
		// 把临时文件转移到最终目录
		if err := gfile.Move(tmpPath, savePath); err != nil {
			// 防止跨文件系统移动文件报错，采用系统原生命令
			var cmd *exec.Cmd
			cmd = exec.Command("mv", tmpPath, savePath)
			_, err = cmd.Output()
			if err != nil {
				return nil, err
			}
		}
	} else {
		// 删除临时文件
		_ = gfile.Remove(tmpPath)
	}

	re = &RemoteFileSaveResult{
		SavePath:     savePath,
		RelativePath: relativePath,
		FileName:     fileName,
		FileSize:     fileSize,
		OriUrl:       url,
		OriFileName:  oriFileName,
	}
	return re, nil
}
