package board_games

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/netooo/board-games/config"
	"github.com/netooo/board-games/infrastructure/persistence"
	"github.com/netooo/board-games/interfaces/handler"
	"github.com/netooo/board-games/usecase"
	"log"
	"net/http"
)

func main() {
	// 依存関係を注入
	userPersistence := persistence.NewUserPersistence(config.Connect())
	userUseCase := usecase.NewUserUseCase(userPersistence)
	userHandler := handler.NewUserHandler(userUseCase)

	// ルーティングの設定
	router := httprouter.New()
	router.GET("/api/users", userHandler.HandleUserGet)
	router.POST("/api/users", userHandler.HandleUserSignup)

	// サーバ起動
	fmt.Println("------------------------")
	fmt.Println("サーバ起動 http://localhost:8080")
	fmt.Println("------------------------")

	_ = http.ListenAndServe(":8080", &Server{router})
	log.Fatal(http.ListenAndServe(":8080", router))
}

type Server struct {
	r *httprouter.Router
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET POST PUT DELETE")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Add("Access-Control-Allow-Headers", "Origin")
	w.Header().Add("Access-Control-Allow-Headers", "X-Requested-With")
	w.Header().Add("Access-Control-Allow-Headers", "Accept")
	w.Header().Add("Access-Control-Allow-Headers", "Accept-Language")
	w.Header().Set("Content-Type", "application/json")
	s.r.ServeHTTP(w, r)
}
