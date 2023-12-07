// routes/route.go
package routes

import (
	"gin_app/config"
	"gin_app/handlers"
	persistence "gin_app/infrastructure"
	"gin_app/usecase"

	"github.com/julienschmidt/httprouter"
)

// TODO:出力内容ごとにパスを分けたい
func AppRoutes() *httprouter.Router {

	userPersistence := persistence.NewUserPersistence(config.Connect())
	userUseCase := usecase.NewUserUseCase(userPersistence)
	userHandler := handlers.NewUserHandler(userUseCase)

	userInformationPersistence := persistence.NewUserInformationPersistence(config.Connect())
	userInformationUseCase := usecase.NewUserInformationUseCase(userInformationPersistence)
	userInformationHandler := handlers.NewUserInformationHandler(userInformationUseCase)

	// ユーザ情報取得
	router := httprouter.New()
	router.GET("/get-user", userHandler.GetUser)

	// ユーザ詳細情報取得
	router.GET("/get-user-information", userInformationHandler.GetUserInformation)

	return router
}
