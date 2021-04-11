package authentication

import (
	"github.com/gorilla/sessions"
	"os"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
