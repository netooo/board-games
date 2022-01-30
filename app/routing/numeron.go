package routing

import (
	"github.com/gorilla/mux"
	"github.com/netooo/board-games/app/config"
	"github.com/netooo/board-games/app/infrastructure/persistence"
	"github.com/netooo/board-games/app/interfaces/handler"
	"github.com/netooo/board-games/app/usecase"
)

func NumeronInit(r *mux.Router) {
	numeronPersistence := persistence.NewNumeronPersistence(config.Connect())
	numeronUseCase := usecase.NewNumeronUseCase(numeronPersistence)
	numeronHandler := handler.NewNumeronHandler(numeronUseCase)

	r.HandleFunc("/numerons", numeronHandler.HandleNumeronGet).Methods("GET")
	r.HandleFunc("/numerons", numeronHandler.HandleNumeronCreate).Methods("POST")
	r.HandleFunc("/numerons/{display_id}", numeronHandler.HandleNumeronShow).Methods("GET")
	r.HandleFunc("/numerons/{display_id}/entry", numeronHandler.HandleNumeronEntry).Methods("POST")
	r.HandleFunc("/numerons/{display_id}/leave", numeronHandler.HandleNumeronLeave).Methods("POST")
	r.HandleFunc("/numerons/{display_id}/start", numeronHandler.HandleNumeronStart).Methods("POST")
	r.HandleFunc("/numerons/{display_id}/set", numeronHandler.HandleNumeronSet).Methods("POST")
	r.HandleFunc("/numerons/{display_id}/attack", numeronHandler.HandleNumeronAttack).Methods("POST")
}
