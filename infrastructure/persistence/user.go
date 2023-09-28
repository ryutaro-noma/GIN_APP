package persistence

//"database/sql"

//"gin_app/domain/repository"

type userPersistence struct {
	UserID string
	Name   string
	Email  string
}

/*func NewUserPersistence() repository.UserRepository {
	return &userPersistence{}
}*/

// userIDによってユーザ情報を取得する TODO:構造体を返すように修正。いったん単純に返答するように
func /*(up userPersistence)*/ GetByUserID(userID string) string {

	user := userPersistence{
		UserID: "A",
		Name:   "野間です。",
		Email:  "~~@gmail.com",
	}

	return user.UserID
}
