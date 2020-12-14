package middleware

import (
	"fmt"
	"os"

	"github.com/api/util"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func SetDB(c *gin.Context) {
	DBHost := os.Getenv("MYSQL_HOST")
	DBPort := "3306"
	DBName := os.Getenv("MYSQL_DATABASE")
	DBUser := os.Getenv("MYSQL_USER")
	DBPass := os.Getenv("MYSQL_PASSWORD")

	// dbとの接続データを格納
	dbConnection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", DBUser, DBPass, DBHost, DBPort, DBName)

	// dbと接続
	db, err := gorm.Open("mysql", dbConnection)

	// エラーの場合そのまま終了
	if err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		panic(err.Error())
	}

	// 最後にdbを閉じる
	defer db.Close()

	// dbのデータを格納
	util.SetDB(c, db)

	c.Next()
}
