package authentication

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/netooo/board-games/app/config"
	"github.com/netooo/board-games/app/domain/model"
	"net/http"
	"os"
	"strings"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

const (
	SessionName       = "session-name"
	ContextSessionKey = "session"
)

func SessionCreate(user *model.User) (*sessions.Session, error) {
	// Session Config
	store.Options = &sessions.Options{
		Secure:   false, // とりあえず開発用に
		MaxAge:   60 * 60 * 24,
		HttpOnly: true,
	}

	// Create New Session
	randomId, _ := uuid.NewRandom()
	sessionId := strings.Replace(randomId.String(), "-", "", -1)

	newSession := sessions.NewSession(store, SessionName)
	newSession.ID = sessionId

	mc := memcache.New("127.0.0.1:11211")
	err := mc.Set(&memcache.Item{Key: sessionId, Value: []byte(user.UserId)})

	if err != nil {
		return nil, err
	}

	return newSession, nil
}

func SessionUser(r *http.Request) *model.User {
	s, _ := store.Get(r, SessionName)
	sessionId := s.ID

	db := config.Connect()
	defer config.Close()

	var user model.User
	var session model.Session

	if err := db.Where("SessionId = ?", sessionId).Find(&session).Related(&user).Error; err != nil {
		return nil
	}

	return &user
}
