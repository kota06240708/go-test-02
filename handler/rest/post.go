package rest

import (
	"net/http"
	"strconv"

	"github.com/api/domain/model"

	"github.com/api/usecase"
	"github.com/api/util"

	"github.com/gin-gonic/gin"
)

type PostHandler interface {
	GetPostAll(*gin.Context)
	GetCurrentPosts(*gin.Context)
	AddPost(*gin.Context)
	UpdatePost(*gin.Context)
	DeletePost(*gin.Context)
	GetUserPosts(*gin.Context)
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

// @description 投稿を全て取得
// @version 1.0
// @Tags post
// @Summary 投稿を全て取得
// @accept application/x-json-stream
// @Security ApiKeyAuth
// @in header
// @name Authorization
// @Success 200 {object} model.PostRes
// @router /api/v1/posts/ [get]
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

// @description ログインユーザーの投稿
// @version 1.0
// @Tags post
// @Summary ログインユーザーの投稿
// @accept application/x-json-stream
// @Security ApiKeyAuth
// @Success 200 {object} model.PostRes
// @router /api/v1/self/posts [GET]
func (ph postHandler) GetCurrentPosts(c *gin.Context) {
	user := util.CurrentUser(c)

	// DBデータを取得
	DB := util.DB(c)

	// DBからデータを取得
	post, err := ph.postUseCase.GetUserPosts(DB, user.ID)

	// エラーかどうかチェック
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//クライアントにレスポンスを返却
	c.JSON(http.StatusOK, &post)
}

type postPostReq struct {
	Text string `json:"text" validate:"required"`
}

// @description 投稿を追加
// @version 2.0
// @Tags post
// @Summary 投稿を追加
// @accept application/x-json-stream
// @param request body postPostReq false "リクエスト"
// @Security ApiKeyAuth
// @Success 200 {object} gin.H {"status": "success"}
// @router /api/v1/self/post [post]
func (ph postHandler) AddPost(c *gin.Context) {

	var req postPostReq

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

type TUpdatePostReq struct {
	Text string `json:"text" validate:"required"`
}

// @description 指定した投稿を更新
// @version 1.0
// @Tags post
// @Summary 指定した投稿を更新
// @accept application/x-json-stream
// @param id path int true "投稿ID"
// @param request body TUpdatePostReq false "リクエスト"
// @Security ApiKeyAuth
// @Success 200 {object} gin.H {"status": "success"}
// @router /api/v1/self/post/:id [PATCH]
func (ph postHandler) UpdatePost(c *gin.Context) {

	var req TUpdatePostReq

	postId, errUint := strconv.ParseUint(c.Param("id"), 10, 32)

	var post model.Post

	// パースがうまくいったか確認
	if errUint != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errUint.Error()})
		return
	}

	// DBデータを取得
	DB := util.DB(c)

	// reqのデータvalidate
	if err, errorMessages := util.GetRequestValidate(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "messages": errorMessages})
		return
	}

	// reqのデータをbind
	if err := util.BindParam(c, &post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// DBから指定したデータを取得
	errPost := ph.postUseCase.UpdatePost(DB, &post, uint(postId))

	// エラーかどうかチェック
	if errPost != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errPost.Error()})
		return
	}

	//クライアントにレスポンスを返却
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

// @description 指定した投稿を削除
// @version 1.0
// @Tags post
// @Summary 指定した投稿を削除
// @accept application/x-json-stream
// @param id path int true "投稿ID"
// @Security ApiKeyAuth
// @Success 200 {object} gin.H {"status": "success"}
// @router /api/v1/self/post/:id [DELETE]
func (ph postHandler) DeletePost(c *gin.Context) {
	postId, errUint := strconv.ParseUint(c.Param("id"), 10, 32)

	// パースがうまくいったか確認
	if errUint != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errUint.Error()})
		return
	}

	// DBデータを取得
	DB := util.DB(c)

	// DBから指定したデータを削除
	err := ph.postUseCase.DeletePost(DB, uint(postId))

	// エラーかどうかチェック
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//クライアントにレスポンスを返却
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

// ShowBottle godoc
// @Summary 指定したユーザーの投稿を取得
// @Tags post
// @description 指定したユーザーの投稿を取得
// @version 1.0
// @accept application/x-json-stream
// @param id path int true "ユーザーID"
// @Security ApiKeyAuth
// @Success 200 {object} []model.PostRes
// @router /api/v1/user/posts/:id [GET]
func (ph postHandler) GetUserPosts(c *gin.Context) {
	postId, errUint := strconv.ParseUint(c.Param("id"), 10, 32)

	// パースがうまくいったか確認
	if errUint != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errUint.Error()})
		return
	}

	// DBデータを取得
	DB := util.DB(c)

	// DBからデータを取得
	posts, err := ph.postUseCase.GetUserPosts(DB, uint(postId))

	// エラーかどうかチェック
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//クライアントにレスポンスを返却
	c.JSON(http.StatusOK, &posts)
}
