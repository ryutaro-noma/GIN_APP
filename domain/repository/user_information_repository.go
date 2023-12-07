package repository

import (
	"gin_app/domain/model"
)

type UserInformationRepository interface {
	GetByUserInformation(Id string) (userInformation *model.UserInformation, err error)
}
