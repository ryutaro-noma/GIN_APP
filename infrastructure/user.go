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

// userIDによってユーザ情報を取得する TODO:構造体を返すように修正。いったん単純に返答するように
func (up *userPersistence) GetByUserID(userID string) (User *model.User, err error) {

	db := up.Conn

	User = &model.User{}

	if err := db.Table("User").Where("UserId = ?", userID).First(&User).Error; err != nil {
		return nil, err
	}

	return User, nil
}
