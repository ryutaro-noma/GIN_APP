package persistence

import (
	"gin_app/domain/model"
	"gin_app/domain/repository"

	"github.com/jinzhu/gorm"
)

type userPersistence struct {
	Conn *gorm.DB
}

func NewUserPersistence(conn *gorm.DB) repository.UserRepository {
	return &userPersistence{Conn: conn}
}

// userIDによってユーザ情報を取得する
func (up *userPersistence) GetByUserID(userID string) (User *model.User, err error) {

	db := up.Conn

	User = &model.User{}

	// Preloadでusersに紐づくuser_informationsのデータを取得したい。
	if err := db.Preload("UserInformation").Find(&User, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}

	/*if err := db.Where("user_id = ?", userID).First(&User).Error; err != nil {
		return nil, err
	}*/

	return User, nil
}
