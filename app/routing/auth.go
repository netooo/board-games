package routing

import (
	"github.com/gorilla/mux"
	"github.com/netooo/board-games/app/config"
	"github.com/netooo/board-games/app/infrastructure/persistence"
	"github.com/netooo/board-games/app/interfaces/handler"
	"github.com/netooo/board-games/app/usecase"
)

func AuthInit(r *mux.Router) {
	// 依存関係を注入
	authPersistence := persistence.NewAuthPersistence(config.Connect())
	authUseCase := usecase.NewAuthUseCase(authPersistence)
	authHandler := handler.NewAuthHandler(authUseCase)

	r.HandleFunc("/signin", authHandler.HandleSignin).Methods("POST")
}
