package usecase

import (
	"gin_app/domain/model"
	"gin_app/domain/repository"
)

type UserUseCase interface {
	GetByUserID(id string) (user *model.User, err error)
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(ur repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: ur,
	}
}

func (uu userUseCase) GetByUserID(id string) (user *model.User, err error) {
	// Persistenceを呼び出し
	user, err = uu.userRepository.GetByUserID(id)
	if err != nil {
		return nil, err
	}
	return user, err
}
