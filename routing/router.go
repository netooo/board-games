package routing

import (
	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	// Set Routing
	r := mux.NewRouter()
	r.Host("localhost") // TODO: 環境によって変えるようにする
	r.PathPrefix("/api")
	r.Schemes("http") // TODO: localではhttp, test/productionではhttpsを使う

	// Read Various APIs
	UserInit(r)

	return r
}
