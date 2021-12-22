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
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	handler := c.Handler(r)

	// サーバ起動
	log.Fatal(http.ListenAndServe(":9000", handler))
}
