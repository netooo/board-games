package authentication

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/netooo/board-games/app/config"
	"github.com/netooo/board-games/app/domain/model"
	"log"
	"net/http"
	"strings"
)

var store = sessions.NewCookieStore([]byte("asdaskdhasdhgsajdgasdsadksakdhasidoajsdousahdopj"))

const (
	SessionName       = "sessionId"
	ContextSessionKey = "session"
)

func SessionCreate(userId string) (*sessions.Session, error) {
	// Session Config
	store.Options = &sessions.Options{
		Secure:   false, // とりあえず開発用に
		MaxAge:   60 * 60 * 24 * 1,
		HttpOnly: true,
	}

	// Create New Session
	randomId, _ := uuid.NewRandom()
	sessionId := strings.Replace(randomId.String(), "-", "", -1)

	newSession := sessions.NewSession(store, SessionName)
	newSession.ID = sessionId

	mc := memcache.New("memcached:11211")
	err := mc.Set(&memcache.Item{
		Key:        sessionId,
		Value:      []byte(userId),
		Expiration: 60 * 60 * 24 * 1,
	})

	if err != nil {
		return nil, err
	}

	return newSession, nil
}

func SessionUser(r *http.Request) (*model.User, error) {
	session, err := store.Get(r, SessionName)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	fmt.Println(session.ID)
	sessionId := session.ID

	var user model.User

	mc := memcache.New("memcached:11211")
	byteUserId, err := mc.Get(sessionId)

	if err != nil {
		return nil, err
	}

	db := config.Connect()
	defer config.Close()

	userId := string(byteUserId.Value)
	if err := db.Where("user_id = ?", userId).Find(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
