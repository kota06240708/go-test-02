package usecase

import (
	"github.com/gin-gonic/gin"

	"github.com/api/domain/model"
	"github.com/api/domain/repository"
)

type TodoUseCase interface {
	TodoGetAll(*gin.Context) ([]*model.Todo, error)
}

type todoUseCase struct {
	todoRepository repository.TodoRepository //TodoRepositoryインターフェースを満たす必要がある
}

//ここでドメイン層のインターフェースとユースケース層のインターフェースをつなげている。
func NewTodoUseCase(tr repository.TodoRepository) TodoUseCase {
	return &todoUseCase{
		todoRepository: tr,
	}
}

//Todoデータを全件取得するためのユースケース
func (tu todoUseCase) TodoGetAll(c *gin.Context) (todos []*model.Todo, err error) {
	// Persistenceを呼出
	todos, err = tu.todoRepository.GetAll(c)
	if err != nil {
		return nil, err
	}
	return todos, nil
}
