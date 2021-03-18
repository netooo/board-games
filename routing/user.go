package routing

import (
	"github.com/gorilla/mux"
	"github.com/netooo/board-games/config"
	"github.com/netooo/board-games/infrastructure/persistence"
	"github.com/netooo/board-games/interfaces/handler"
	"github.com/netooo/board-games/usecase"
)

func UserInit(r *mux.Router) {
	// 依存関係を注入
	userPersistence := persistence.NewUserPersistence(config.Connect())
	userUseCase := usecase.NewUserUseCase(userPersistence)
	userHandler := handler.NewUserHandler(userUseCase)
	r.HandleFunc("/users/{id}", userHandler.HandleUserGet).Methods("GET")
	r.HandleFunc("/users", userHandler.HandleUserSignup).Methods("POST")
}
