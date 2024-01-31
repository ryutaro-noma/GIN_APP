package inmemory

import (
	"gin_app/domain/model"
	"gin_app/domain/repository"
	"strconv"
)

var dataStore model.UserInformation

type userInformationPersistence struct {
}

func NewUserInformationPersistence() repository.UserInformationRepository {
	dataStore = model.UserInformation{
		Id:        9,
		LastName:  "firstName",
		FirstName: "lastName",
		Sex:       "man",
		Age:       10,
	}
	return &userInformationPersistence{}
}

// userIDによってユーザ情報を取得する
func (up *userInformationPersistence) GetByUserInformation(userID string) (UserInfo *model.UserInformation, err error) {
	if strconv.Itoa(dataStore.Id) == userID {
		return &dataStore, nil
	}
	return nil, nil
}
