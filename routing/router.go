package routing

import (
	"github.com/gorilla/mux"
)

//var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
//
//func handleFunc(r *mux.Router, path string, fn http.HandlerFunc) *mux.Route {
//	return r.HandleFunc(path, applicationHandler(fn))
//}
//
//func applicationHandler(fn http.HandlerFunc) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		// セッションの取得
//		session, err := store.Get(r, SessionName)
//		if err != nil {
//			// 不正なセッションだった場合は作り直す
//			session, err = store.New(r, SessionName)
//			checkError(err)
//		}
//		context.Set(r, ContextSessionKey, session)
//		// 個別のハンドラー呼び出し
//		fn(w, r)
//	}
//}

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
