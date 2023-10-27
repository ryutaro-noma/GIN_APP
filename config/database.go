package config

import (
	"gin_app/domain/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db  *gorm.DB
	err error
)

// DB接続
func Connect() *gorm.DB {

	// DB接続
	db, err = gorm.Open("mysql", "user:p@ssw0rd@tcp(127.0.0.1:3306)/hoge-db?charset=utf8mb4")
	if err != nil {
		panic(err)
	}

	autoMigration()

	return db
}

// DB終了
func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func autoMigration() {
	db.AutoMigrate(&model.User{})
}
