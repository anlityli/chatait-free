// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package helper

import (
	"bytes"
	"encoding/hex"
	"errors"
	"github.com/gogf/gf/os/gfile"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
)

type FileTypeInfo struct {
	Ext      string
	MimeType string
}

var fileTypeMap sync.Map

func init() {
	fileTypeMap.Store("ffd8ffe000104a464946", &FileTypeInfo{Ext: "jpg", MimeType: "image/jpeg"})                //JPEG (jpg)
	fileTypeMap.Store("ffd8ff", &FileTypeInfo{Ext: "jpg", MimeType: "image/jpeg"})                              //JPEG (jpg)
	fileTypeMap.Store("89504e470d0a1a0a0000", &FileTypeInfo{Ext: "png", MimeType: "image/png"})                 //PNG (png)
	fileTypeMap.Store("89504e47", &FileTypeInfo{Ext: "png", MimeType: "image/png"})                             //PNG (png)
	fileTypeMap.Store("47494638396126026f01", &FileTypeInfo{Ext: "gif", MimeType: "image/gif"})                 //GIF (gif)
	fileTypeMap.Store("47494638", &FileTypeInfo{Ext: "gif", MimeType: "image/gif"})                             //GIF (gif)
	fileTypeMap.Store("49492a00227105008037", &FileTypeInfo{Ext: "tif", MimeType: "image/tif"})                 //TIFF (tif)
	fileTypeMap.Store("49492a00", &FileTypeInfo{Ext: "tif", MimeType: "image/tif"})                             //TIFF (tif)
	fileTypeMap.Store("424d228c010000000000", &FileTypeInfo{Ext: "bmp", MimeType: "image/bmp"})                 //16色位图(bmp)
	fileTypeMap.Store("424d8240090000000000", &FileTypeInfo{Ext: "bmp", MimeType: "image/bmp"})                 //24位位图(bmp)
	fileTypeMap.Store("424d8e1b030000000000", &FileTypeInfo{Ext: "bmp", MimeType: "image/bmp"})                 //256色位图(bmp)
	fileTypeMap.Store("424d", &FileTypeInfo{Ext: "bmp", MimeType: "image/bmp"})                                 //bmp
	fileTypeMap.Store("41433130313500000000", &FileTypeInfo{Ext: "dwg", MimeType: "application/x-autocad"})     //CAD (dwg)
	fileTypeMap.Store("41433130", &FileTypeInfo{Ext: "dwg", MimeType: "application/x-autocad"})                 //CAD (dwg)
	fileTypeMap.Store("3c21444f435459504520", &FileTypeInfo{Ext: "html", MimeType: "text/html"})                //HTML (html)   3c68746d6c3e0  3c68746d6c3e0
	fileTypeMap.Store("3c68746d6c3e0", &FileTypeInfo{Ext: "html", MimeType: "text/html"})                       //HTML (html)   3c68746d6c3e0  3c68746d6c3e0
	fileTypeMap.Store("3c21646f637479706520", &FileTypeInfo{Ext: "html", MimeType: "text/html"})                //HTM (htm)
	fileTypeMap.Store("68746d6c3e", &FileTypeInfo{Ext: "html", MimeType: "text/html"})                          //HTM (htm)
	fileTypeMap.Store("48544d4c207b0d0a0942", &FileTypeInfo{Ext: "css", MimeType: "text/css"})                  //css
	fileTypeMap.Store("696b2e71623d696b2e71", &FileTypeInfo{Ext: "js", MimeType: "application/x-javascript"})   //js
	fileTypeMap.Store("7b5c727466315c616e73", &FileTypeInfo{Ext: "rtf", MimeType: "application/rtf"})           //Rich Text Format (rtf)
	fileTypeMap.Store("38425053000100000000", &FileTypeInfo{Ext: "psd", MimeType: "image/vnd.adobe.photoshop"}) //Photoshop (psd)
	fileTypeMap.Store("46726f6d3a203d3f6762", &FileTypeInfo{Ext: "eml", MimeType: "message/rfc822"})            //Email [Outlook Express 6] (eml)
	fileTypeMap.Store("d0cf11e0a1b11ae10000", &FileTypeInfo{Ext: "doc", MimeType: "application/vnd.ms-word"})   //MS Excel 注意：word、msi 和 excel的文件头一样
	fileTypeMap.Store("d0cf11e0a1b11ae10000", &FileTypeInfo{Ext: "vsd", MimeType: "application/vnd.visio"})     //Visio 绘图
	fileTypeMap.Store("5374616E64617264204A", &FileTypeInfo{Ext: "mdb", MimeType: "application/x-msaccess"})    //MS Access (mdb)
	fileTypeMap.Store("252150532D41646F6265", &FileTypeInfo{Ext: "ps", MimeType: "application/postscript"})
	fileTypeMap.Store("255044462d312e350d0a", &FileTypeInfo{Ext: "pdf", MimeType: "application/pdf"})       //Adobe Acrobat (pdf)
	fileTypeMap.Store("255044462d312e", &FileTypeInfo{Ext: "pdf", MimeType: "application/pdf"})             //Adobe Acrobat (pdf)
	fileTypeMap.Store("2e524d46000000120001", &FileTypeInfo{Ext: "rmvb", MimeType: "audio/x-pn-realaudio"}) //rmvb/rm相同
	fileTypeMap.Store("464c5601050000000900", &FileTypeInfo{Ext: "flv", MimeType: "video/x-flv"})           //flv与f4v相同
	fileTypeMap.Store("00000020667479706", &FileTypeInfo{Ext: "mp4", MimeType: "video/mp4"})
	fileTypeMap.Store("0000001c667479706", &FileTypeInfo{Ext: "mp4", MimeType: "video/mp4"})
	fileTypeMap.Store("49443303000000002176", &FileTypeInfo{Ext: "mp3", MimeType: "audio/mpeg"})
	fileTypeMap.Store("000001ba210001000180", &FileTypeInfo{Ext: "mpg", MimeType: "video/mpeg"})     //
	fileTypeMap.Store("3026b2758e66cf11a6d9", &FileTypeInfo{Ext: "wmv", MimeType: "video/x-ms-wmv"}) //wmv与asf相同
	fileTypeMap.Store("52494646e27807005741", &FileTypeInfo{Ext: "wav", MimeType: "audio/wav"})      //Wave (wav)
	fileTypeMap.Store("52494646d07d60074156", &FileTypeInfo{Ext: "avi", MimeType: "video/x-msvideo"})
	fileTypeMap.Store("4d546864000000060001", &FileTypeInfo{Ext: "mid", MimeType: "audio/midi"}) //MIDI (mid)
	fileTypeMap.Store("504b0304140000000800", &FileTypeInfo{Ext: "zip", MimeType: "application/zip"})
	fileTypeMap.Store("504b0304", &FileTypeInfo{Ext: "zip", MimeType: "application/zip"})
	fileTypeMap.Store("526172211a0700cf9073", &FileTypeInfo{Ext: "rar", MimeType: "application/x-rar"})
	fileTypeMap.Store("52617221", &FileTypeInfo{Ext: "rar", MimeType: "application/x-rar"})
	fileTypeMap.Store("235468697320636f6e66", &FileTypeInfo{Ext: "ini", MimeType: ""})
	fileTypeMap.Store("504b03040a0000000000", &FileTypeInfo{Ext: "jar", MimeType: "application/java-archive"})
	fileTypeMap.Store("4d5a9000030000000400", &FileTypeInfo{Ext: "exe", MimeType: "application/octet-stream"})                                                 //可执行文件
	fileTypeMap.Store("3c25402070616765206c", &FileTypeInfo{Ext: "jsp", MimeType: ""})                                                                         //jsp文件
	fileTypeMap.Store("4d616e69666573742d56", &FileTypeInfo{Ext: "mf", MimeType: ""})                                                                          //MF文件
	fileTypeMap.Store("3c3f786d6c2076657273", &FileTypeInfo{Ext: "xml", MimeType: "text/xml"})                                                                 //xml文件
	fileTypeMap.Store("494e5345525420494e54", &FileTypeInfo{Ext: "sql", MimeType: "text/x-sql"})                                                               //xml文件
	fileTypeMap.Store("7061636b616765207765", &FileTypeInfo{Ext: "java", MimeType: "text/x-java-source"})                                                      //java文件
	fileTypeMap.Store("406563686f206f66660d", &FileTypeInfo{Ext: "bat", MimeType: ""})                                                                         //bat文件
	fileTypeMap.Store("1f8b0800000000000000", &FileTypeInfo{Ext: "gz", MimeType: "application/x-gzip"})                                                        //gz文件
	fileTypeMap.Store("6c6f67346a2e726f6f74", &FileTypeInfo{Ext: "properties", MimeType: ""})                                                                  //bat文件
	fileTypeMap.Store("cafebabe0000002e0041", &FileTypeInfo{Ext: "class", MimeType: "application/x-java-applet"})                                              //bat文件
	fileTypeMap.Store("49545346030000006000", &FileTypeInfo{Ext: "chm", MimeType: "application/octet-stream"})                                                 //bat文件
	fileTypeMap.Store("04000000010000001300", &FileTypeInfo{Ext: "mxp", MimeType: ""})                                                                         //bat文件
	fileTypeMap.Store("504b0304140006000800", &FileTypeInfo{Ext: "docx", MimeType: "application/vnd.openxmlformats-officedocument.wordprocessingml.document"}) //docx文件
	fileTypeMap.Store("d0cf11e0a1b11ae10000", &FileTypeInfo{Ext: "wps", MimeType: "application/vnd.ms-works"})                                                 //WPS文字wps、表格et、演示dps都是一样的
	fileTypeMap.Store("6431303a637265617465", &FileTypeInfo{Ext: "torrent", MimeType: ""})
	fileTypeMap.Store("6D6F6F76", &FileTypeInfo{Ext: "mov", MimeType: "video/quicktime"})      //Quicktime (mov)
	fileTypeMap.Store("FF575043", &FileTypeInfo{Ext: "wpd", MimeType: ""})                     //WordPerfect (wpd)
	fileTypeMap.Store("CFAD12FEC5FD746F", &FileTypeInfo{Ext: "dbx", MimeType: ""})             //Outlook Express (dbx)
	fileTypeMap.Store("2142444E", &FileTypeInfo{Ext: "pst", MimeType: ""})                     //Outlook (pst)
	fileTypeMap.Store("AC9EBD8F", &FileTypeInfo{Ext: "qdf", MimeType: ""})                     //Quicken (qdf)
	fileTypeMap.Store("E3828596", &FileTypeInfo{Ext: "pwl", MimeType: ""})                     //Windows Password (pwl)
	fileTypeMap.Store("2E7261FD", &FileTypeInfo{Ext: "ram", MimeType: "audio/x-pn-realaudio"}) //Real Audio (ram)
}

// 获取前面结果字节的二进制
func bytesToHexString(src []byte) string {
	res := bytes.Buffer{}
	if src == nil || len(src) <= 0 {
		return ""
	}
	temp := make([]byte, 0)
	for _, v := range src {
		sub := v & 0xFF
		hv := hex.EncodeToString(append(temp, sub))
		if len(hv) < 2 {
			res.WriteString(strconv.FormatInt(int64(0), 10))
		}
		res.WriteString(hv)
	}
	return res.String()
}

// GetFileType 用文件前面几个字节来判断
// fSrc: 文件字节流（就用前面几个字节）
func GetFileType(fSrc []byte) *FileTypeInfo {
	var fileType *FileTypeInfo
	fileCode := bytesToHexString(fSrc)

	fileTypeMap.Range(func(key, value interface{}) bool {
		k := key.(string)
		v := value.(*FileTypeInfo)
		if strings.HasPrefix(fileCode, strings.ToLower(k)) ||
			strings.HasPrefix(k, strings.ToLower(fileCode)) {
			fileType = v
			return false
		}
		return true
	})
	return fileType
}

// GetFileExtName 获取文件的扩展名
func GetFileExtName(filePath string) (re *FileTypeInfo, err error) {
	fHandle, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	fSrc, err := ioutil.ReadAll(fHandle)
	if err != nil {
		return nil, err
	}
	re = GetFileType(fSrc[:10])
	if re == nil {
		return nil, errors.New("文件不能识别")
	}
	return re, nil
}

// RemoveFile 删除文件
func RemoveFile(filePath string) (err error) {
	if gfile.IsFile(filePath) {
		return gfile.Remove(filePath)
	}
	return errors.New("不是文件")
}

// FormatDirStr 格式化目录 把不带/结尾的目录带上/
func FormatDirStr(dirPath string) (re string) {
	if dirPath[len(dirPath)-1:] != "/" {
		return dirPath + "/"
	}
	return dirPath
}
