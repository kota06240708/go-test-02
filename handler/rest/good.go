package rest

import (
	"net/http"
	"strconv"

	"github.com/api/util"

	"github.com/api/usecase"
	"github.com/gin-gonic/gin"
)

type GoodHandler interface {
	SetGood(*gin.Context)
}

type goodHandler struct {
	goodUseCase usecase.GoodUseCase
}

func NewGoodHandler(gu usecase.GoodUseCase) GoodHandler {
	return &goodHandler{
		goodUseCase: gu,
	}
}

// @description いいねを更新
// @version 1.0
// @Tags good
// @Summary いいねを更新
// @accept application/x-json-stream
// @param id path int true "投稿ID"
// @Security ApiKeyAuth
// @in header
// @name Authorization
// @Success 200 {object} gin.H {"status": "success"}
// @router /api/v1/good/:id [POST]
func (gh goodHandler) SetGood(c *gin.Context) {
	type TReq struct {
		IsGood *bool `json:"isGood" validate:"required"`
	}

	var req TReq

	// ユーザー情報を取得
	user := util.CurrentUser(c)

	// パラメータを取得
	postId, errParam := strconv.Atoi(c.Param("id"))
	if errParam != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errParam.Error()})
		return
	}

	// reqのデータvalidate
	if err, errorReq := util.GetRequestValidate(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "messages": errorReq})
		return
	}

	// DBデータを取得
	DB := util.DB(c)

	// いいねの処理を追加
	if err := gh.goodUseCase.UpdateGood(DB, user.ID, postId, req.IsGood); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//クライアントにレスポンスを返却
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
