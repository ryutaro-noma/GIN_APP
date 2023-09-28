// handlers/newrecipehandler.go
package handlers

import (
	"gin_app/domain"
	"gin_app/domain/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

// hellow worldを返答するだけ
func HellowWorld(c *gin.Context) {
	c.JSON(http.StatusOK, "hellow world")
}

// ユーザー情報の取得(現状、直書きのデータを出力するだけ)
func GetUserInfomation(c *gin.Context) {
	var user domain.User
	user.UserID = xid.New().String()
	user.Name = "田中　太郎"
	user.Email = "xxx@gmail.com"
	c.JSON(http.StatusOK, user)
}

// ユーザー情報の更新(未実装)　TODO：追々Dockerと接続してデータ更新できるようにしたい。
func NewUserHandler(c *gin.Context) {
	var user repository.UserRepository
	// リクエストデータを取り出す
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
