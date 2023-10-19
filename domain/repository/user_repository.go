package repository

import (
	"gin_app/domain/model"
)

type UserRepository interface {
	GetByUserID(userID string) (user *model.User, err error)
}
