// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package notice

const (
	NoError       ErrorCode = 0
	OtherError    ErrorCode = 400
	NotAuth       ErrorCode = 401
	NoPrint       ErrorCode = 402
	AuthForbidden ErrorCode = 403
	NotFind       ErrorCode = 404
	ShowDialog    ErrorCode = 405
)
