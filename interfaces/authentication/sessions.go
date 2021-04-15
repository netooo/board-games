package authentication

import (
	"github.com/gorilla/sessions"
	"net/http"
	"os"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

const (
	SessionName       = "session-name"
	ContextSessionKey = "session"
)

func sessionInit(writer http.ResponseWriter, request *http.Request) {
	// Create New Session
	session := sessions.NewSession(store, SessionName)

	// Session Config
	store.Options = &sessions.Options{
		Secure:   false, // とりあえず開発用に
		MaxAge:   60 * 15,
		HttpOnly: true,
	}

	session.Save(request, writer)
}
