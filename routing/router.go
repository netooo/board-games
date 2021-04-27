package routing

import (
	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	// Set Routing
	r := mux.NewRouter()
	s := r.PathPrefix("/api").
		Host("localhost"). // TODO: 環境によって変えるようにする
		Schemes("http").   // TODO: localではhttp, test/productionではhttpsを使う
		Subrouter()

	// Read Various APIs
	UserInit(s)

	return s
}
