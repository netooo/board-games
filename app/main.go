package main

import (
	"github.com/netooo/board-games/app/routing"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	// ルーティングの設定
	r := routing.Init()
	c := cors.Default().Handler(r)

	// サーバ起動
	log.Fatal(http.ListenAndServe(":9000", c))
}
