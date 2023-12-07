package persistence

import (
	"gin_app/domain/model"
	"gin_app/domain/repository"

	"github.com/jinzhu/gorm"
)

type userInformationPersistence struct {
	Conn *gorm.DB
}

func NewUserInformationPersistence(conn *gorm.DB) repository.UserInformationRepository {
	return &userInformationPersistence{Conn: conn}
}

// userIDによってユーザ情報を取得する
func (up *userInformationPersistence) GetByUserInformation(userID string) (UserInfo *model.UserInformation, err error) {

	db := up.Conn

	UserInfo = &model.UserInformation{}

	if err := db.Where("id = ?", userID).First(&UserInfo).Error; err != nil {
		return nil, err
	}

	return UserInfo, nil
}
