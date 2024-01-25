package repository

import (
	"gin_app/domain/model"
)

//go:generate mockgen -source=C:\dev\gin_app\domain\repository\user_information_repository.go -destination C:\dev\gin_app\domain\repository\repository_mock\user_information_mock.go
type UserInformationRepository interface {
	GetByUserInformation(Id string) (userInformation *model.UserInformation, err error)
}
