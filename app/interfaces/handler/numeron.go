package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/netooo/board-games/app/interfaces/authentication"
	"github.com/netooo/board-games/app/interfaces/response"
	"github.com/netooo/board-games/app/usecase"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type NumeronHandler interface {
	HandleRoomCreate(http.ResponseWriter, *http.Request)
	HandleGameStart(http.ResponseWriter, *http.Request)
}

type numeronHandler struct {
	numeronUseCase usecase.NumeronUseCase
}

type gameStartRequest struct {
	First  string
	Second string
	Third  string
	Fourth string
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

func (nh numeronHandler) HandleGameStart(writer http.ResponseWriter, request *http.Request) {
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

	vars := mux.Vars(request)
	id := vars["id"]

	if id == "" {
		response.BadRequest(writer, "Invalid path parameter")
		return
	}

	var numeronId int
	numeronId, err = strconv.Atoi(id)

	if err != nil {
		response.BadRequest(writer, "Invalid path patameter")
		return
	}

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		response.BadRequest(writer, "Invalid Request Body")
		return
	}

	// request Body から順番を受け取る
	var requestBody gameStartRequest
	_ = json.Unmarshal(body, &requestBody)

	//TODO: Check request_user already join other room?
	// もしやるんだったら Userテーブルに Statusカラムを追加しないといけなさそう

	err = nh.numeronUseCase.GameStart(user.ID, socket, numeronId, requestBody)
	if err != nil {
		response.InternalServerError(writer, "Internal Server Error")
		return
	}

	response.Success(writer, nil)
}
