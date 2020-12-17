package util

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

type TError struct {
	Key string `json:"key"`
	Tag string `json:"tag"`
	Err string `json:"err"`
}

// requestを取得
func GetRequest(c *gin.Context, data interface{}) (error, []*TError) {
	validate := validator.New()

	// reqのjsonデータを取得
	if errBind := c.ShouldBindJSON(&data); errBind != nil {
		return errBind, nil
	}

	// validate
	if errValidate := validate.Struct(data); errValidate != nil {
		var errorMessages []*TError //バリデーションでNGとなった独自エラーメッセージを格納

		for _, err := range errValidate.(validator.ValidationErrors) {

			errorMessage := &TError{
				Key: err.Field(),
				Tag: err.Tag(),
				Err: "error message for " + err.Field(),
			}

			errorMessages = append(errorMessages, errorMessage)
		}

		return errValidate, errorMessages
	}

	return nil, nil
}
