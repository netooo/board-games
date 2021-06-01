package routing

import (
	"github.com/gorilla/mux"
	"github.com/netooo/board-games/config"
	"github.com/netooo/board-games/infrastructure/persistence"
	"github.com/netooo/board-games/interfaces/handler"
	"github.com/netooo/board-games/usecase"
)

func CommonInit(r *mux.Router) {
	// 依存関係を注入
	commonPersistence := persistence.NewCommonPersistence(config.Connect())
	commonUseCase := usecase.NewCommonUseCase(commonPersistence)
	commonHandler := handler.NewCommonHandler(commonUseCase)

	r.HandleFunc("/room", commonHandler.HandleRoomPost).Methods("POST")
}
