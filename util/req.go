package util

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v9"
)

type TError struct {
	Key string `json:"key"`
	Tag string `json:"tag"`
	Err string `json:"err"`
}

// requestを取得
func BindParam(c *gin.Context, data interface{}) error {
	// reqのjsonデータを取得
	if err := c.ShouldBindBodyWith(&data, binding.JSON); err != nil {
		return err
	}

	return nil
}

// requestを取得
func GetRequestValidate(c *gin.Context, data interface{}) (error, []*TError) {
	validate := validator.New()

	// reqのjsonデータを取得
	if errBind := c.ShouldBindBodyWith(&data, binding.JSON); errBind != nil {
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
