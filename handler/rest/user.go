package rest

import (
	"net/http"
	"strconv"

	_ "github.com/api/domain/model"
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

// @description ユーザー一覧を取得
// @version 1.0
// @Tags user
// @Summary ユーザー一覧を取得
// @accept application/x-json-stream
// @Security ApiKeyAuth
// @in header
// @name Authorization
// @Success 200 {object} []model.User
// @router /api/v1/users [GET]
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

// @description idのユーザーを削除
// @version 1.0
// @Tags user
// @Summary idのユーザーを削除
// @accept application/x-json-stream
// @Security ApiKeyAuth
// @in header
// @param id path int true "ユーザーID"
// @name Authorization
// @Success 200 {object} gin.H {"status": "success"}
// @router /api/v1/user/:id [DELETE]
func (uh userHandler) DeleteUser(c *gin.Context) {
	userId, errUint := strconv.Atoi(c.Param("id"))

	// パースがうまくいったか確認
	if errUint != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errUint.Error()})
		return
	}

	// DBデータを取得
	DB := util.DB(c)

	// DBにあるデータを削除
	err := uh.userUseCase.DeleteUser(DB, userId)

	// エラーかどうかチェック
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//クライアントにレスポンスを返却
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

// @description ログインユーザー情報
// @version 1.0
// @Tags user
// @Summary ログインユーザー情報
// @accept application/x-json-stream
// @Security ApiKeyAuth
// @in header
// @name Authorization
// @Success 200 {object} model.User
// @router /api/v1/self/user [GET]
func (uh userHandler) GetCurrentUser(c *gin.Context) {
	user := util.CurrentUser(c)

	//クライアントにレスポンスを返却
	c.JSON(http.StatusOK, user)
}

type TUserAddRequset struct {
	Name     string `json:"name" validate:"required"`
	Age      int    `json:"age" validate:"required"`
	Icon     string `json:"icon" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required"`
}

// @description ユーザーを追加
// @version 1.0
// @Tags user
// @Summary ユーザーを追加
// @accept application/x-json-stream
// @param request body TUserAddRequset false "リクエスト"
// @Security ApiKeyAuth
// @in header
// @name Authorization
// @Success 200 {object} gin.H {"status": "success"}
// @router /api/v1/signup [POST]
func (uh userHandler) AddUser(c *gin.Context) {

	var req TUserAddRequset

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

type TUpdateUserRequset struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Icon     string `json:"icon"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// @description ユーザー情報を更新
// @version 1.0
// @Tags user
// @Summary ユーザー情報を更新
// @accept application/x-json-stream
// @param request body TUpdateUserRequset false "リクエスト"
// @Security ApiKeyAuth
// @in header
// @name Authorization
// @Success 200 {object} gin.H {"status": "success"}
// @router /api/v1/self/user [PATCH]
func (uh userHandler) UpdateUser(c *gin.Context) {

	var param TUpdateUserRequset

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
