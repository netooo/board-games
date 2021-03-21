package config

import (
	"github.com/jinzhu/gorm"
	"os"
)

var (
	db  *gorm.DB
	err error
)

// DB接続
func Connect() *gorm.DB {
	// 実行環境取得
	DBName := os.Getenv("DB_NAME")
	DBUser := os.Getenv("DB_USER")
	DBPass := os.Getenv("DB_PASS")

	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@/"+DBName+"?charset=utf8&parseTime=True&loc=Local")

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
