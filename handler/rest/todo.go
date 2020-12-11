package rest

import (
	"net/http"
	"time"

	"github.com/api/usecase"
	"github.com/gin-gonic/gin"
)

type TodoHandler interface {
	Index(*gin.Context)
}

//この構造体は元々TodoUseCaseinterfaceと紐づいていて、Indexメソッドの宣言の際にこの構造体と新たに紐づけられる
type todoHandler struct {
	todoUseCase usecase.TodoUseCase
}

// NewTodoUseCase : Todo データに関する Handler を生成
func NewTodokHandler(tu usecase.TodoUseCase) TodoHandler {
	return &todoHandler{
		todoUseCase: tu,
	}
}

//Index: Get /todos -> todoデータ一覧取得
func (th todoHandler) Index(c *gin.Context) {
	//request：TodoAPIのパラメータ
	//type requset struct {
	//  Begin uint `query:begin`
	//  Limit uint `query:limit`
	//}

	type TodoField struct {
		Id        int64     `json:"id"`
		Title     string    `json:"title"`
		Author    string    `json:"author"`
		CreatedAt time.Time `json:"created_at"`
	}
	//response : Todo　API　のレスポンス
	type response struct {
		Todos []TodoField `json:"todos"`
	}

	//ユースケースの呼び出し
	todos, err := th.todoUseCase.TodoGetAll(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//取得したドメインモデルをresponseに変換
	res := new(response)
	for _, todo := range todos {
		var tf TodoField
		tf.Id = int64(todo.Id)
		tf.Title = todo.Title
		tf.Author = todo.Author
		tf.CreatedAt = todo.CreatedAt
		res.Todos = append(res.Todos, tf)
	}

	//クライアントにレスポンスを返却
	c.JSON(http.StatusOK, &res)
}
