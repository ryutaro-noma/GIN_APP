// routes/route.go
package routes

import (
	"gin_app/handlers"

	"github.com/gin-gonic/gin"
)

// TODO:出力内容ごとにパスを分けたい
func AppRoutes() *gin.Engine {
	router := gin.Default()

	// HellowWorldの出力
	router.GET("hellow-world", handlers.HellowWorld)

	// ユーザ情報取得
	router.GET("/get-user-information", handlers.GetUserInfomation)

	return router
}
