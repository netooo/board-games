package routing

import (
	"github.com/gorilla/mux"
	"github.com/netooo/board-games/app/config"
	"github.com/netooo/board-games/app/infrastructure/persistence"
	"github.com/netooo/board-games/app/interfaces/handler"
	"github.com/netooo/board-games/app/usecase"
)

func NumeronInit(r *mux.Router) {
	// Numeron
	numeronPersistence := persistence.NewNumeronPersistence(config.Connect())
	numeronUseCase := usecase.NewNumeronUseCase(numeronPersistence)
	numeronHandler := handler.NewNumeronHandler(numeronUseCase)

	r.HandleFunc("/numerons", numeronHandler.HandleRoomCreate).Methods("POST")
	r.HandleFunc("/numerons/{id}/start", numeronHandler.HandleGameStart).Methods("POST")
}
