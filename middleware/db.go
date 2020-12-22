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
	DBPort := os.Getenv("MYSQL_PORT")
	DBName := os.Getenv("MYSQL_DATABASE")
	DBNameTest := os.Getenv("MYSQL_DATABASE_TEST")
	DBUser := os.Getenv("MYSQL_USER")
	DBUserTest := os.Getenv("MYSQL_ROOT_USER")
	DBPass := os.Getenv("MYSQL_PASSWORD")
	DBPassTest := os.Getenv("MYSQL_ROOT_PASSWORD")

	MODE := os.Getenv("MODE")

	// dbとの接続データを格納
	var dbConnection string

	if MODE == "test" {
		dbConnection = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", DBUserTest, DBPassTest, DBHost, DBPort, DBNameTest)
	} else {
		dbConnection = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", DBUser, DBPass, DBHost, DBPort, DBName)
	}

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
