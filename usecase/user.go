package usecase

import (
	"database/sql"
)

// User における UseCase のインターフェース
/*type UserUseCase interface {
	GetByUesrID(DB *sql.DB, userID string) (domain.User, error)
	Insert(DB *sql.DB, userID, name, email string) error
}*/

/*type userUseCase struct {
	userRepository repository.UserRepository
}*/

// Userデータに対するusecaseを生成
/*func NewUserUseCase(ur repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: ur,
	}
}*/

func /*(uu UserUsecase)*/ GetByUserID(DB *sql.DB, userID string) string /**domain.User, error*/ {
	return "A"
}
