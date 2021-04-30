package routing

import (
	"github.com/gorilla/mux"
	"github.com/netooo/board-games/config"
	"github.com/netooo/board-games/infrastructure/persistence"
	"github.com/netooo/board-games/interfaces/handler"
	"github.com/netooo/board-games/usecase"
)

func NumeronInit(r *mux.Router) {
	// 依存関係を注入
	numeronPersistence := persistence.NewNumeronPersistence(config.Connect())
	numeronUseCase := usecase.NewNumeronUseCase(numeronPersistence)
	numeronHandler := handler.NewNumeronHandler(numeronUseCase)
}
