package board_games

import (
	"github.com/gorilla/mux"
	"github.com/netooo/board-games/config"
	"github.com/netooo/board-games/infrastructure/persistence"
	"github.com/netooo/board-games/interfaces/handler"
	"github.com/netooo/board-games/usecase"
	"net/http"
)

func main() {
	// 依存関係を注入
	userPersistence := persistence.NewUserPersistence(config.Connect())
	userUseCase := usecase.NewUserUseCase(userPersistence)
	userHandler := handler.NewUserHandler(userUseCase)

	// ルーティングの設定
	r := mux.NewRouter()
	r.Host("localhost") // TODO: 環境によって変えるようにする
	r.PathPrefix("/api")
	r.Schemes("http") // TODO: localではhttp, test/productionではhttpsを使う

	r.HandleFunc("/users/{id}", userHandler.HandleUserGet).Methods("GET")
	r.HandleFunc("/users", userHandler.HandleUserSignup).Methods("POST")

	// サーバ起動
	_ = http.ListenAndServe(":8080", r) // TODO: localのみにしたい
}
