package util

import (
	"github.com/api/domain/model"

	"github.com/gin-gonic/gin"
)

var variableNameUser = "currentUser"

// ログインユーザーを取得
func CurrentUser(c *gin.Context) *model.User {
	return c.MustGet(variableNameUser).(*model.User)
}

// ログインユーザーを格納
func SetCurrentUser(c *gin.Context, user *model.User) {
	c.Set(variableNameUser, user)
}
