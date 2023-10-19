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
	// 実行環境取得
	/*env := os.Getenv("HAKEN_ENV")

	if "production" == env {
		env = "production"
	} else {
		env = "development"
	}

	// 環境変数取得
	godotenv.Load(".env." + env)
	godotenv.Load()*/

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

/*func init() {
	var err error
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		"user", "p@ssw0rd", "mysql:3306", "hoge-db",
	)
	Db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
}*/
