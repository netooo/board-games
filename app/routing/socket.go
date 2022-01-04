package routing

import (
	"github.com/gorilla/mux"
	"github.com/netooo/board-games/app/config"
	"github.com/netooo/board-games/app/infrastructure/persistence"
	"github.com/netooo/board-games/app/interfaces/handler"
	"github.com/netooo/board-games/app/usecase"
)

func SocketInit(r *mux.Router) {
	socketPersistence := persistence.NewSocketPersistence(config.Connect())
	socketUseCase := usecase.NewSocketUseCase(socketPersistence)
	socketHandler := handler.NewSocketHandler(socketUseCase)

	r.HandleFunc("/ws/connect", socketHandler.HandleSocketConnect).Methods("GET")
	r.HandleFunc("/ws/disconnect", socketHandler.HandleSocketDisconnect).Methods("GET")
}
