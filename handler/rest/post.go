package rest

import (
	"net/http"

	"github.com/api/domain/model"

	"github.com/api/usecase"
	"github.com/api/util"

	"github.com/gin-gonic/gin"
)

type PostHandler interface {
	GetPostAll(*gin.Context)
	GetCurrentPosts(*gin.Context)
	AddPost(*gin.Context)
	GetSelectPosts(*gin.Context)
}

// postcaseのintefaceと紐ずける
type postHandler struct {
	postUseCase usecase.PostUseCase
}

// NewTodoUseCase : Todo データに関する Handler を生成
func NewPostHandler(ph usecase.PostUseCase) PostHandler {
	return &postHandler{
		postUseCase: ph,
	}
}

// ユーザー一覧を取得
func (ph postHandler) GetPostAll(c *gin.Context) {

	// DBデータを取得
	DB := util.DB(c)

	// DBからデータを取得
	posts, err := ph.postUseCase.GetPosts(DB)

	// エラーかどうかチェック
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//クライアントにレスポンスを返却
	c.JSON(http.StatusOK, &posts)
}

// 現在の投稿を受け取る
func (ph postHandler) GetCurrentPosts(c *gin.Context) {
	user := util.CurrentUser(c)

	// DBデータを取得
	DB := util.DB(c)

	// DBからデータを取得
	post, err := ph.postUseCase.GetSelectPost(DB, user.ID)

	// エラーかどうかチェック
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//クライアントにレスポンスを返却
	c.JSON(http.StatusOK, &post)
}

// 投稿を追加
func (ph postHandler) AddPost(c *gin.Context) {
	type TReq struct {
		Text string `json:"text" validate:"required"`
	}

	var req TReq

	// reqのデータvalidate
	if err, errorMessages := util.GetRequestValidate(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "messages": errorMessages})
		return
	}

	var comment model.Post

	// reqのデータをbind
	if err := util.BindParam(c, &comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 投稿したユーザーデータ
	user := util.CurrentUser(c)

	// ユーザーIDを追加
	comment.UserId = user.ID

	// DBデータを取得
	DB := util.DB(c)

	// DBにデータを追加
	_, err := ph.postUseCase.AddPost(DB, &comment)

	// エラーかどうかチェック
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//クライアントにレスポンスを返却
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

// 指定したコメントを取得
func (ph postHandler) GetSelectPosts(c *gin.Context) {
	type TReq struct {
		UserId uint `json:"name" validate:"required"`
	}

	var req *TReq

	// reqのデータをbind
	if err, errorMessages := util.GetRequestValidate(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "messages": errorMessages})
		return
	}

	// DBデータを取得
	DB := util.DB(c)

	// DBからデータを取得
	posts, err := ph.postUseCase.GetSelectPost(DB, req.UserId)

	// エラーかどうかチェック
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//クライアントにレスポンスを返却
	c.JSON(http.StatusOK, &posts)
}