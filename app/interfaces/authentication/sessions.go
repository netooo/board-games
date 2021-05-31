package authentication

import (
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/netooo/board-games/app/config"
	"github.com/netooo/board-games/app/domain/model"
	"net/http"
	"os"
	"strings"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
var session model.Session

const (
	SessionName       = "session-name"
	ContextSessionKey = "session"
)

func CreateSession(user *model.User) *sessions.Session {
	// Session Config
	store.Options = &sessions.Options{
		Secure:   false, // とりあえず開発用に
		MaxAge:   60 * 60 * 24,
		HttpOnly: true,
	}

	// TODO: interface層でDB操作はしたくないが, いずれredisにするので一旦はこれでOK
	// Create New Session
	randomId, _ := uuid.NewRandom()
	sessionId := strings.Replace(randomId.String(), "-", "", -1)

	newSession := sessions.NewSession(store, SessionName)
	newSession.ID = sessionId
	s := model.Session{
		SessionId: sessionId,
		Data:      "",
		User:      *user,
	}

	db := config.Connect()
	defer config.Close()

	db.Create(&s)

	return newSession
}

func GetSessionUser(r *http.Request) (*model.User, error) {
	s, err := store.Get(r, SessionName)
	if err != nil {
		return nil, err
	}

	sessionId := s.ID

	db := config.Connect()
	defer config.Close()

	db.Where("session_id = ?", sessionId).First(&session)
	user := &session.User
	db.Model(&session).Related(&user)

	return user, nil
}
