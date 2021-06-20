package handler

import (
	"github.com/gorilla/mux"
	"github.com/netooo/board-games/app/interfaces/authentication"
	"github.com/netooo/board-games/app/interfaces/response"
	"github.com/netooo/board-games/app/usecase"
	"net/http"
)

type NumeronHandler interface {
	HandleRoomCreate(http.ResponseWriter, *http.Request)
}

type numeronHandler struct {
	numeronUseCase usecase.NumeronUseCase
}

func NewNumeronHandler(nu usecase.NumeronUseCase) NumeronHandler {
	return &numeronHandler{
		numeronUseCase: nu,
	}
}

func (nh numeronHandler) HandleRoomCreate(writer http.ResponseWriter, request *http.Request) {
	user := authentication.SessionUser(request)
	if user == nil {
		// TODO: redirect login form
		response.Unauthorized(writer, "Invalid Session")
		return
	}

	vars := mux.Vars(request)
	numeronId := vars["id"]

	if numeronId == "" {
		response.BadRequest(writer, "Invalid Request Body")
		return
	}

	//TODO: Check request_user already join other room?
	// もしやるんだったら Userテーブルに Statusカラムを追加しないといけなさそう

	room, err := nh.numeronUseCase.CreateRoom(user)
	if err != nil {
		response.InternalServerError(writer, "Internal Server Error")
		return
	}

	response.Success(writer, room)
}
