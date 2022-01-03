package routing

import (
	"github.com/gorilla/mux"
	"github.com/netooo/board-games/app/interfaces/handler"
)

func SocketInit(r *mux.Router) {
	socketHandler := handler.NewSocketHandler()

	r.HandleFunc("/ws/connect", socketHandler.HandleSocketConnect).Methods("GET")
	r.HandleFunc("/ws/disconnect", socketHandler.HandleSocketDisconnect).Methods("GET")
}
