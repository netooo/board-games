package routing

import (
	"github.com/gorilla/mux"
	"github.com/netooo/board-games/app/config"
	"github.com/netooo/board-games/app/infrastructure/persistence"
	"github.com/netooo/board-games/app/interfaces/handler"
	"github.com/netooo/board-games/app/usecase"
)

func RoomInit(r *mux.Router) {
	// Room
	roomPersistence := persistence.NewRoomPersistence(config.Connect())
	roomUseCase := usecase.NewRoomUseCase(roomPersistence)
	roomHandler := handler.NewRoomHandler(roomUseCase)

	r.HandleFunc("/rooms", roomHandler.HandleRoomGet).Methods("GET")
	r.HandleFunc("/ws/rooms", roomHandler.HandleRoomCreate).Methods("GET")
	r.HandleFunc("/ws/rooms/{id}/join", roomHandler.HandleRoomJoin).Methods("GET")
}
