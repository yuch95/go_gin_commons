package resource

import (
	"tools.com/libs/libs/auth"
	"tools.com/libs/libs/database"
)

type Resource struct {
	DbDao    database.DbDao
	UserInfo auth.UserInfo
}
