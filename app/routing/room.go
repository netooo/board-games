package routing

import (
	"github.com/gorilla/mux"
	"github.com/netooo/board-games/app/config"
	"github.com/netooo/board-games/app/infrastructure/persistence"
	"github.com/netooo/board-games/app/interfaces/handler"
	"github.com/netooo/board-games/app/usecase"
)

func RoomInit(r *mux.Router) {
	roomPersistence := persistence.NewRoomPersistence(config.Connect())
	roomUseCase := usecase.NewRoomUseCase(roomPersistence)
	roomHandler := handler.NewRoomHandler(roomUseCase)

	r.HandleFunc("/rooms", roomHandler.HandleRoomGet).Methods("GET")
	r.HandleFunc("/rooms", roomHandler.HandleRoomCreate).Methods("POST")
	r.HandleFunc("/rooms/{id}", roomHandler.HandleRoomShow).Methods("GET")
	r.HandleFunc("/rooms/{id}", roomHandler.HandleRoomJoin).Methods("POST")
}
