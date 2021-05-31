package routing

import (
	"github.com/gorilla/mux"
	"github.com/netooo/board-games/config"
	"github.com/netooo/board-games/infrastructure/persistence"
	"github.com/netooo/board-games/interfaces/handler"
	"github.com/netooo/board-games/usecase"
)

func NumeronInit(r *mux.Router) {
	// NumeronPlayer
	numeronPlayerPersistence := persistence.NewNumeronPlayerPersistence(config.Connect())
	numeronPlayerUseCase := usecase.NewNumeronPlayerUseCase(numeronPlayerPersistence)
	numeronPlayerHandler := handler.NewNumeronPlayerHandler(numeronPlayerUseCase)
}
