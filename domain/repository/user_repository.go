package repository

import (
	"database/sql"
	"gin_app/domain"
)

type UserRepository interface {
	GetByUserID(DB *sql.DB, userID string) (*domain.User, error)
}
