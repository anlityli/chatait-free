// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package response

type FileMidjourney struct {
	Id         string `json:"id"`
	UserId     string `json:"user_id"`
	Username   string `json:"username"`
	Nickname   string `json:"nickname"`
	QueueId    string `json:"queue_id"`
	FileName   string `json:"file_name"`
	Path       string `json:"path"`
	Prompt     string `json:"prompt"`
	MjFileName string `json:"mj_file_name"`
	MjUrl      string `json:"mj_url"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
	Size       int    `json:"size"`
	CreatedAt  int    `json:"created_at"`
}

type FileMidjourneyList []*FileMidjourney
