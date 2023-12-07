// handlers/newrecipehandler.go
package handlers

import (
	"encoding/json"
	"gin_app/usecase"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserInformationHandler interface {
	GetUserInformation(http.ResponseWriter, *http.Request, httprouter.Params)
}

type userInformationHandler struct {
	userInformationUsecase usecase.UserInformationUseCase
}

func NewUserInformationHandler(ui usecase.UserInformationUseCase) UserInformationHandler {
	return &userInformationHandler{
		userInformationUsecase: ui,
	}
}

// ユーザー詳細情報の取得
func (uh userInformationHandler) GetUserInformation(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {

	id := r.FormValue("Id")
	// usecaseの呼び出し
	user, err := uh.userInformationUsecase.GetByUserInformation(id)
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
