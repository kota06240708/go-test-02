package repository

import (
	"github.com/api/domain/model"

	"github.com/gin-gonic/gin"
)

type TodoRepository interface {
	GetAll(c *gin.Context) ([]*model.Todo, error)
}
