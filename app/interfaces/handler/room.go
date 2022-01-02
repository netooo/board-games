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

type RoomHandler interface {
	HandleRoomGet(http.ResponseWriter, *http.Request)
	HandleRoomCreate(http.ResponseWriter, *http.Request)
	HandleRoomJoin(http.ResponseWriter, *http.Request)
}

type roomHandler struct {
	roomUseCase usecase.RoomUseCase
}

const (
	socketBufferSize = 1024
)

/* websocket用の変数 */
var upgrader = &websocket.Upgrader{
	ReadBufferSize:  socketBufferSize,
	WriteBufferSize: socketBufferSize,
}

func NewRoomHandler(ru usecase.RoomUseCase) RoomHandler {
	return &roomHandler{
		roomUseCase: ru,
	}
}
func (rh roomHandler) HandleRoomGet(writer http.ResponseWriter, request *http.Request) {
	_, err := authentication.SessionUser(request)
	if err != nil {
		// TODO: redirect login form
		response.Unauthorized(writer, "Invalid Session")
		return
	}

	//TODO: Check request_user already join other room?
	// もしやるんだったら Userテーブルに Statusカラムを追加しないといけなさそう

	rooms, err := rh.roomUseCase.GetRooms()
	if err != nil {
		response.InternalServerError(writer, "Internal Server Error")
		return
	}

	response.Success(writer, rooms)
}

func (rh roomHandler) HandleRoomCreate(writer http.ResponseWriter, request *http.Request) {
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

	//TODO: Check request_user already join other room?
	// もしやるんだったら Userテーブルに Statusカラムを追加しないといけなさそう

	if err := rh.roomUseCase.CreateRoom(user, socket); err != nil {
		response.InternalServerError(writer, "Internal Server Error")
		return
	}

	response.Success(writer, "")
}

func (rh roomHandler) HandleRoomJoin(writer http.ResponseWriter, request *http.Request) {
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

	err = rh.roomUseCase.JoinRoom(id, user, socket)
	if err != nil {
		response.InternalServerError(writer, err.Error())
		return
	}

	response.Success(writer, "")
}
