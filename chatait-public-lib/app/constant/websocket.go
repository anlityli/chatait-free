// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package constant

const (
	WSMsgRequestTypeExample = "request_example"
)

const (
	WSMsgResponseTypeError  = "response_error"
	WSMsgResponseTypeIgnore = "response_ignore"

	WSMsgResponseTypeMidjourneyInsertQueue = "response_mj_insert_queue"
	WSMsgResponseTypeMidjourneyCreate      = "response_mj_create"
	WSMsgResponseTypeMidjourneyEnd         = "response_mj_end"
	WSMsgResponseTypeMidjourneyError       = "response_mj_error"
	WSMsgResponseTypeMidjourneyProgress    = "response_mj_progress"
)
