package main

import (
	"github.com/netooo/board-games/routing"
	"log"
	"net/http"
)

func main() {
	// ルーティングの設定
	r := routing.Init()

	// サーバ起動
	log.Fatal(http.ListenAndServe(":8080", r))
}
