package handler

import (
	"github.com/netooo/board-games/app/usecase"
	"net/http"
)

type CommonHandler interface {
}

type commonHandler struct {
	commonUseCase usecase.CommonUseCase
}

func NewCommonHandler(cu usecase.CommonUseCase) CommonHandler {
	return &commonHandler{
		commonUseCase: cu,
	}
}

func (ch commonHandler) HandleRoomCreate(writer http.ResponseWriter, request *http.Request) {

}
