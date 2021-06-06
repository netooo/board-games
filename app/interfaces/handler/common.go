package handler

import (
	"encoding/json"
	"github.com/netooo/board-games/app/domain/model"
	"github.com/netooo/board-games/app/interfaces/response"
	"github.com/netooo/board-games/app/usecase"
	"io/ioutil"
	"net/http"
)

type CommonHandler interface {
	HandleRoomCreate(http.ResponseWriter, *http.Request)
}

type commonHandler struct {
	commonUseCase usecase.CommonUseCase
}

type roomCreateRequest struct {
	Game string
}

func NewCommonHandler(cu usecase.CommonUseCase) CommonHandler {
	return &commonHandler{
		commonUseCase: cu,
	}
}

func (ch commonHandler) HandleRoomCreate(writer http.ResponseWriter, request *http.Request) {
	//TODO: session check
	//TODO: request_user already join room check
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		response.BadRequest(writer, "Invalid Request Body")
		return
	}

	var requestBody roomCreateRequest
	_ = json.Unmarshal(body, &requestBody)

	room, err := ch.commonUseCase.CreateRoom(user, requestBody.Game)
	if err != nil {
		response.InternalServerError(writer, "Internal Server Error")
		return
	}

	response.Success(writer, room)
}
