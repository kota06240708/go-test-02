package test

import (
	"encoding/json"
	"log"
	"net/http"
	"testing"

	"github.com/api/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestExampleSuccess(t *testing.T) {
	users := []model.User{}
	param := &TApiData{
		Type:  "GET",
		Url:   "/api/v1/users",
		Param: "",
	}

	// apiを送信
	body, _ := SendApi(param)

	if err := json.Unmarshal(body, &users); err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, 200)
}
