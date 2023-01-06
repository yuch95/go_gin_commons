package resource

import (
	"tools.com/libs/libs/auth"
	"tools.com/libs/libs/database"
)

type Resource struct {
	DbDao    database.DBDao
	UserInfo auth.UserInfo
}
