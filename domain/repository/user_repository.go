package repository

import (
	"database/sql"
	domain "gin_app/domain/model"
)

type UserRepository interface {
	GetByUserID(DB *sql.DB, userID string) (*domain.User, error)
}
