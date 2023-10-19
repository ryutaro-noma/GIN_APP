// handlers/newrecipehandler.go
package handlers

import (
	"encoding/json"
	"gin_app/usecase"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserHandler interface {
	GetUserInfomation(http.ResponseWriter, *http.Request, httprouter.Params)
}

type userHandler struct {
	userUsecase usecase.UserUseCase
}

func NewUserHandler(uu usecase.UserUseCase) UserHandler {
	return &userHandler{
		userUsecase: uu,
	}
}

// ユーザー情報の取得(現状、直書きのデータを出力するだけ)
func (uh userHandler) GetUserInfomation(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {

	id := r.FormValue("userID")
	// usecaseの呼び出し
	user, err := uh.userUsecase.GetByUserID(id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// クライアントにレスポンスを返却
	if err = json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
