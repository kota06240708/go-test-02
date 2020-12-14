package rest

import (
	"net/http"

	"github.com/api/usecase"
	"github.com/api/util"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	GetUserAll(*gin.Context)
	AddUser(*gin.Context)
}

// usercaseのintefaceと紐ずける
type userHandler struct {
	userUseCase usecase.UserUseCase
}

// NewTodoUseCase : Todo データに関する Handler を生成
func NewUserkHandler(tu usecase.UserUseCase) UserHandler {
	return &userHandler{
		userUseCase: tu,
	}
}

// ユーザー一覧を取得
func (uh userHandler) GetUserAll(c *gin.Context) {

	// DBデータを取得
	DB := util.DB(c)

	// DBからデータを取得
	users, err := uh.userUseCase.GetUserAll(DB)

	// エラーかどうかチェック
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//クライアントにレスポンスを返却
	c.JSON(http.StatusOK, &users)
}

// ユーザーを追加
func (uh userHandler) AddUser(c *gin.Context) {
	//request：TodoAPIのパラメータ
	type TRequset struct {
		Name     string `json:"name" validate:"required"`
		Age      int    `json:"age" validate:"required"`
		Icon     string `json:"icon" validate:"required"`
		Password string `json:"password" validate:"required"`
		Email    string `json:"email" validate:"required"`
	}

	var req TRequset

	// DBデータを取得
	DB := util.DB(c)

	// ユーザーデータを取得
	if err := util.GetRequest(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// パスワードを暗号化
	password, errPass := util.PasswordEncrypt(req.Password)

	if errPass != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errPass.Error()})
		return
	}

	// DBにデータを追加
	if errDB := uh.userUseCase.AddUser(DB, req.Name, req.Age, req.Icon, password, req.Email); errDB != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errDB.Error()})
		return
	}

	//クライアントにレスポンスを返却
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
