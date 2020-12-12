package util

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

// requestを取得
func GetRequest(c *gin.Context, data interface{}) error {
	validate := validator.New()

	// reqのjsonデータを取得
	if errBind := c.ShouldBindJSON(&data); errBind != nil {
		return errBind
	}

	// validate
	if errValidate := validate.Struct(data); errValidate != nil {
		return errValidate
	}

	return nil
}
