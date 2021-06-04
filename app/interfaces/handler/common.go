package handler

import (
	"encoding/json"
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

	var requestBody userSignupRequest
	_ = json.Unmarshal(body, &requestBody)

}
