package user_api

import "QAComm/models"

type UserApi struct {
}
type Query struct {
	models.PageInfo
	UserId int `form:"user_id"`
}
