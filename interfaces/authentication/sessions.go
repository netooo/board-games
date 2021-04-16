package authentication

import (
	"github.com/gorilla/sessions"
	"os"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

const (
	SessionName       = "session-name"
	ContextSessionKey = "session"
)

func SessionInit() *sessions.Session {
	// Session Config
	store.Options = &sessions.Options{
		Secure:   false, // とりあえず開発用に
		MaxAge:   60 * 15,
		HttpOnly: true,
	}

	// Create New Session
	return sessions.NewSession(store, SessionName)
}
