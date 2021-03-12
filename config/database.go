package config

import (
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"os"
)

var (
	db  *gorm.DB
	err error
)

// DB接続
func Connect() *gorm.DB {
	// 実行環境取得
	env := os.Getenv("ENV")

	if "production" == env {
		env = "production"
	} else {
		env = "development"
	}

	// 環境変数取得
	_ = godotenv.Load(".env." + env)
	_ = godotenv.Load()

	// DB接続
	db, err = gorm.Open("mysql", os.Getenv("CONNECT"))

	if err != nil {
		panic(err)
	}

	return db
}

// DB終了
func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}
