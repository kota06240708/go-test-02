package test

import (
	"bytes"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http/httptest"
	"os"

	"github.com/api/router"
	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/jinzhu/gorm"
)

type TApiData struct {
	Type  string
	Url   string
	Param string
}

// APIを叩く
func SendApi(t *TApiData) ([]byte, error) {
	router := router.StartRouter()

	req := httptest.NewRequest(t.Type, t.Url, changeBaffer(t.Param))
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	body, err := ioutil.ReadAll(rec.Body)

	return body, err
}

func changeBaffer(str string) *bytes.Buffer {
	return bytes.NewBuffer([]byte(str))
}

// gormのDBを取得
func GetDB() (*gorm.DB, error) {
	DBHost := os.Getenv("MYSQL_HOST")
	DBPort := os.Getenv("MYSQL_PORT")
	DBNameTest := os.Getenv("MYSQL_DATABASE_TEST")
	DBUserTest := os.Getenv("MYSQL_ROOT_USER")
	DBPassTest := os.Getenv("MYSQL_ROOT_PASSWORD")

	// dbとの接続データを格納
	dbConnection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", DBUserTest, DBPassTest, DBHost, DBPort, DBNameTest)

	// dbと接続
	db, err := gorm.Open("mysql", dbConnection)

	return db, err
}

// DB構築
func InitDB(pass string) error {
	var (
		db       *sql.DB
		fixtures *testfixtures.Loader
	)

	DBHost := os.Getenv("MYSQL_HOST")
	DBNameTest := os.Getenv("MYSQL_DATABASE_TEST")
	DBUserTest := os.Getenv("MYSQL_ROOT_USER")
	DBPassTest := os.Getenv("MYSQL_ROOT_PASSWORD")

	// dbと接続
	db, errDB := sql.Open("mysql", DBUserTest+":"+DBPassTest+"@tcp("+DBHost+":3306)/"+DBNameTest+"?parseTime=true")
	if errDB != nil {
		log.Fatal(errDB)
	}

	var errFixtures error

	p, _ := os.Getwd()

	fmt.Println(p)

	fixtures, errFixtures = testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("mysql"),
		testfixtures.Directory(pass+"fixtures"),
	)
	if errFixtures != nil {
		log.Fatal(errFixtures)
	}

	err := fixtures.Load()

	return err
}
