// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package column

// AllListColumn 所有列表的字段(用于分配字段权限，每增加一个列表，都要在这里写一个)
func AllListColumn() (re []interface{}) {
	re = []interface{}{
		&Admin{},
		&Config{},
		&Conversation{},
		&File{},
		&Finance{},
		&Shop{},
		&User{},
	}
	return
}
