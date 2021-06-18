package routing

import (
	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	// Set Routing
	r := mux.NewRouter()
	s := r.PathPrefix("/api").
		Schemes("http"). // TODO: localではhttp, test/productionではhttpsを使う
		Subrouter()

	// Read Various APIs
	UserInit(s)

	// Numeron
	NumeronInit(s)
	NumeronPlayerInit(s)

	return s
}
