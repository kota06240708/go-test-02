package util

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var variableNameDB = "DB"

// DBを取得
func DB(c *gin.Context) *gorm.DB {
	return c.MustGet(variableNameDB).(*gorm.DB)
}

// DBを格納
func SetDB(c *gin.Context, db *gorm.DB) {
	c.Set(variableNameDB, db)
}
