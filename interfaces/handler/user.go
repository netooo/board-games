package handler

import (
	"github.com/julienschmidt/httprouter"
	"github.com/netooo/board-games/usecase"
	"net/http"
)

type UserHandler interface {
	HandleUserGet(http.ResponseWriter, *http.Request, httprouter.Params)
	HandleUserSignup(http.ResponseWriter, *http.Request, httprouter.Params)
}

type userHandler struct {
	userUseCase usecase.UserUseCase
}
