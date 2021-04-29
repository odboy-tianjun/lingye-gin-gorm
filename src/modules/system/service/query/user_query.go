package query

import "lingye-gin/src/base"

type UserQuery struct {
	base.SimplePage
	Username string `json:"username"`
}
