package rest

import (
	"net/http"

	"github.com/api/usecase"
	"github.com/api/util"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	GetUserAll(*gin.Context)
	GetCurrentUser(*gin.Context)
	AddUser(*gin.Context)
	UpdateUser(*gin.Context)
	DeleteUser(*gin.Context)
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

// idのユーザーを削除
func (uh userHandler) DeleteUser(c *gin.Context) {
	//request：TodoAPIのパラメータ
	type TRequset struct {
		UserId int `json:"user_id" validate:"required"`
	}

	// Reqを受け取る
	var param TRequset

	// DBデータを取得
	DB := util.DB(c)

	// validate
	if err, messages := util.GetRequestValidate(c, &param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "messages": messages})
		return
	}

	// DBにあるデータを削除
	err := uh.userUseCase.DeleteUser(DB, param.UserId)

	// エラーかどうかチェック
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//クライアントにレスポンスを返却
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

// 現在のユーザー情報を返す
func (uh userHandler) GetCurrentUser(c *gin.Context) {
	user := util.CurrentUser(c)

	//クライアントにレスポンスを返却
	c.JSON(http.StatusOK, user)
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
	if err, errorMessages := util.GetRequestValidate(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "messages": errorMessages})
		return
	}

	// パスワードを暗号化
	password, errPass := util.PasswordEncrypt(req.Password)

	if errPass != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errPass.Error()})
		return
	}

	// awsにアイコンをアップデート
	url, errch := util.UploadToS3(req.Icon)

	select {
	case err := <-errch:
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	case url := <-url:

		// DBにデータを追加
		if errDB := uh.userUseCase.AddUser(DB, req.Name, req.Age, url, password, req.Email); errDB != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errDB.Error()})
			return
		}

		//クライアントにレスポンスを返却
		c.JSON(http.StatusOK, gin.H{"status": "success"})
	}
}

// ユーザー一覧を取得
func (uh userHandler) UpdateUser(c *gin.Context) {
	type TRequset struct {
		Name     string `json:"name"`
		Age      int    `json:"age"`
		Icon     string `json:"icon"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	var param TRequset

	// DBデータを取得
	DB := util.DB(c)
	currentUser := util.CurrentUser(c)

	// validate
	if err, _ := util.GetRequestValidate(c, &param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// reqとparamをマージ
	if err := util.BindParam(c, &currentUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// アイコンがないそのまま再代入
	if param.Icon == "" {
		// DBのデータを更新
		if errDB := uh.userUseCase.UpdateUser(DB, currentUser); errDB != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errDB.Error()})
			return
		}

		//クライアントにレスポンスを返却
		c.JSON(http.StatusOK, &currentUser)
		return
	}

	// アイコンがある場合はawsにアイコンをアップデートしてからDBを更新
	url, errIcon := util.UploadToS3(param.Icon)

	select {
	case err := <-errIcon:
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	case url := <-url:
		currentUser.Icon = url

		// DBのデータを更新
		if errDB := uh.userUseCase.UpdateUser(DB, currentUser); errDB != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errDB.Error()})
			return
		}

		//クライアントにレスポンスを返却
		c.JSON(http.StatusOK, &currentUser)
	}
}
