package repository

import (
	"gin_app/domain/model"
)

//go:generate mockgen -source=C:\dev\gin_app\domain\repository\user_repository.go -destination C:\dev\gin_app\domain\repository\repository_mock\user_mock.go
type UserRepository interface {
	GetByUserID(userID string) (user *model.User, err error)
}
