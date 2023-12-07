package usecase

import (
	"gin_app/domain/model"
	"gin_app/domain/repository"
)

type UserInformationUseCase interface {
	GetByUserInformation(id string) (userInformation *model.UserInformation, err error)
}

type userInformationUseCase struct {
	userInformationRepository repository.UserInformationRepository
}

func NewUserInformationUseCase(ur repository.UserInformationRepository) UserInformationUseCase {
	return &userInformationUseCase{
		userInformationRepository: ur,
	}
}

func (uu userInformationUseCase) GetByUserInformation(id string) (user *model.UserInformation, err error) {
	// Persistenceを呼び出し
	user, err = uu.userInformationRepository.GetByUserInformation(id)
	if err != nil {
		return nil, err
	}
	return user, err
}
