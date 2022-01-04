package handler

import (
	"github.com/gorilla/websocket"
	"github.com/netooo/board-games/app/interfaces/authentication"
	"github.com/netooo/board-games/app/interfaces/response"
	"github.com/netooo/board-games/app/usecase"
	"log"
	"net/http"
)

type SocketHandler interface {
	HandleSocketConnect(http.ResponseWriter, *http.Request)
	HandleSocketDisconnect(http.ResponseWriter, *http.Request)
}

type socketHandler struct {
	socketUseCase usecase.SocketUseCase
}

const (
	socketBufferSize = 1024
)

/* websocket用の変数 */
var upgrader = &websocket.Upgrader{
	ReadBufferSize:  socketBufferSize,
	WriteBufferSize: socketBufferSize,
}

func NewSocketHandler(su usecase.SocketUseCase) SocketHandler {
	return &socketHandler{
		socketUseCase: su,
	}
}

func (sh socketHandler) HandleSocketConnect(writer http.ResponseWriter, request *http.Request) {
	user, err := authentication.SessionUser(request)
	if err != nil {
		// TODO: redirect login form
		response.Unauthorized(writer, "Invalid Session")
		return
	}

	/* websocketの開設 */
	socket, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Fatalln("websocketの開設に失敗しました。:", err)
	}

	// check to already connect socket

	if err := sh.socketUseCase.ConnectSocket(user, socket); err != nil {
		response.InternalServerError(writer, "Internal Server Error")
		return
	}

	response.Success(writer, "")
}

func (sh socketHandler) HandleSocketDisconnect(writer http.ResponseWriter, request *http.Request) {

}
