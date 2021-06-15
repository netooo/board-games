package routing

import (
	"github.com/gorilla/mux"
	"github.com/netooo/board-games/app/config"
	"github.com/netooo/board-games/app/infrastructure/persistence"
	"github.com/netooo/board-games/app/interfaces/handler"
	"github.com/netooo/board-games/app/usecase"
)

func NumeronInit(r *mux.Router) {
	// NumeronPlayer
	numeronPlayerPersistence := persistence.NewNumeronPlayerPersistence(config.Connect())
	numeronPlayerUseCase := usecase.NewNumeronPlayerUseCase(numeronPlayerPersistence)
	numeronPlayerHandler := handler.NewNumeronPlayerHandler(numeronPlayerUseCase)
}
