package routing

import (
	"github.com/gorilla/mux"
	"github.com/netooo/board-games/app/config"
	"github.com/netooo/board-games/app/infrastructure/persistence"
	"github.com/netooo/board-games/app/interfaces/handler"
	"github.com/netooo/board-games/app/usecase"
)

func NumeronPlayerInit(r *mux.Router) {
	// NumeronPlayer
	numeronPlayerPersistence := persistence.NewNumeronPlayerPersistence(config.Connect())
	numeronPlayerUseCase := usecase.NewNumeronPlayerUseCase(numeronPlayerPersistence)
	numeronPlayerHandler := handler.NewNumeronPlayerHandler(numeronPlayerUseCase)

	r.HandleFunc("/numerons/{id}/code", numeronPlayerHandler.HandleNumeronSetCode).Methods("POST")
	r.HandleFunc("/numerons/{id}/join", numeronPlayerHandler.HandleNumeronJoinRoom).Methods("POST")
}
