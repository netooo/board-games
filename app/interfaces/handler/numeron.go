package handler

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/netooo/board-games/app/interfaces/authentication"
	"github.com/netooo/board-games/app/interfaces/response"
	"github.com/netooo/board-games/app/usecase"
	"log"
	"net/http"
)

type NumeronHandler interface {
	HandleRoomCreate(http.ResponseWriter, *http.Request)
	HandleRoomJoin(http.ResponseWriter, *http.Request)
}

type numeronHandler struct {
	numeronUseCase usecase.NumeronUseCase
}

const (
	socketBufferSize = 1024
)

/* websocket用の変数 */
var upgrader = &websocket.Upgrader{
	ReadBufferSize:  socketBufferSize,
	WriteBufferSize: socketBufferSize,
}

func NewNumeronHandler(nu usecase.NumeronUseCase) NumeronHandler {
	return &numeronHandler{
		numeronUseCase: nu,
	}
}

func (nh numeronHandler) HandleRoomCreate(writer http.ResponseWriter, request *http.Request) {
	/* websocketの開設 */
	socket, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Fatalln("websocketの開設に失敗しました。:", err)
	}

	user, err := authentication.SessionUser(request)

	if err != nil {
		// TODO: redirect login form
		response.Unauthorized(writer, "Invalid Session")
		return
	}

	//TODO: Check request_user already join other room?
	// もしやるんだったら Userテーブルに Statusカラムを追加しないといけなさそう

	room, err := nh.numeronUseCase.CreateRoom(user, socket)
	if err != nil {
		response.InternalServerError(writer, "Internal Server Error")
		return
	}

	response.Success(writer, room)
}

func (nh numeronHandler) HandleRoomJoin(writer http.ResponseWriter, request *http.Request) {
	/* websocketの開設 */
	socket, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Fatalln("websocketの開設に失敗しました。:", err)
	}

	user, err := authentication.SessionUser(request)

	if err != nil {
		// TODO: redirect login form
		response.Unauthorized(writer, "Invalid Session")
		return
	}

	// パスパラメータを取得
	vars := mux.Vars(request)
	id := vars["id"]

	if id == "" {
		response.StatusNotFound(writer, "Status Not Found")
		return
	}

	err = nh.numeronUseCase.JoinRoom(id, user, socket)
	if err != nil {
		response.InternalServerError(writer, "Internal Server Error")
		return
	}

	response.Success(writer, "")
}
