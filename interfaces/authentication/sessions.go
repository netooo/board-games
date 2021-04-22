package authentication

import (
	"github.com/gorilla/sessions"
	"github.com/netooo/board-games/config"
	"github.com/netooo/board-games/domain/model"
	"os"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

const (
	SessionName       = "session-name"
	ContextSessionKey = "session"
)

func SessionCreate(user *model.User) *sessions.Session {
	// Session Config
	store.Options = &sessions.Options{
		Secure:   false, // とりあえず開発用に
		MaxAge:   60 * 60 * 24,
		HttpOnly: true,
	}

	// TODO: interface層でDB操作はしたくないが, いずれredisにするので一旦はこれでOK
	// Create New Session
	sesison := sessions.NewSession(store, SessionName)
	session := model.Session{
		SessionId: sessionId,
		Data:      nil,
		User:      *user,
	}

	db := config.Connect()
	defer config.Close()

	db.Create(&sesison)

	return sesison
}
