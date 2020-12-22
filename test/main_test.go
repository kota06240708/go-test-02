package test

import (
	"bytes"
	"database/sql"
	"io/ioutil"
	"log"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/api/router"
	"github.com/go-testfixtures/testfixtures/v3"
)

var (
	db       *sql.DB
	fixtures *testfixtures.Loader
)

func TestMain(m *testing.M) {
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

	fixtures, errFixtures = testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("mysql"),
		testfixtures.Directory("./fixtures"),
	)
	if errFixtures != nil {
		log.Fatal(errFixtures)
	}

	os.Exit(m.Run())
}

// DB構築
func PrepareTestDatabase() {
	if err := fixtures.Load(); err != nil {
		log.Fatal(err)
	}
}

type TApiData struct {
	Type  string
	Url   string
	Param string
	Body  interface{}
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
