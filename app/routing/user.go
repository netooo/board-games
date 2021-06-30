package routing

import (
	"github.com/gorilla/mux"
	"github.com/netooo/board-games/app/config"
	"github.com/netooo/board-games/app/infrastructure/persistence"
	"github.com/netooo/board-games/app/interfaces/handler"
	"github.com/netooo/board-games/app/usecase"
)

func UserInit(r *mux.Router) {
	// 依存関係を注入
	userPersistence := persistence.NewUserPersistence(config.Connect())
	userUseCase := usecase.NewUserUseCase(userPersistence)
	userHandler := handler.NewUserHandler(userUseCase)

	r.HandleFunc("/users/{user_id}", userHandler.HandleUserGet).Methods("GET")
	r.HandleFunc("/users", userHandler.HandleUserSignup).Methods("POST")
}
