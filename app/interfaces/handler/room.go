package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/netooo/board-games/app/interfaces/authentication"
	"github.com/netooo/board-games/app/interfaces/response"
	"github.com/netooo/board-games/app/usecase"
	"io/ioutil"
	"net/http"
)

type RoomHandler interface {
	HandleRoomGet(http.ResponseWriter, *http.Request)
	HandleRoomCreate(http.ResponseWriter, *http.Request)
	HandleRoomShow(http.ResponseWriter, *http.Request)
	HandleRoomJoin(http.ResponseWriter, *http.Request)
}

type roomHandler struct {
	roomUseCase usecase.RoomUseCase
}

type createRoomRequest struct {
	Name string
}

type getResponse struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Owner   string `json:"owner"`
	Players int    `json:"players"`
}

type createResponse struct {
	RoomId uint
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

	var res []*getResponse
	for _, r_ := range rooms {
		r := getResponse{
			Id:      r_.ID,
			Name:    r_.Name,
			Owner:   r_.Owner.Name,
			Players: len(r_.Players),
		}
		res = append(res, &r)
	}

	response.Success(writer, res)
}

func (rh roomHandler) HandleRoomCreate(writer http.ResponseWriter, request *http.Request) {
	user, err := authentication.SessionUser(request)
	if err != nil {
		// TODO: redirect login form
		response.Unauthorized(writer, "Invalid Session")
		return
	}

	//TODO: Check request_user already join other room?
	// もしやるんだったら Userテーブルに Statusカラムを追加しないといけなさそう

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		response.BadRequest(writer, "Invalid Request Body")
		return
	}

	// リクエストボディのパース
	var requestBody createRoomRequest
	_ = json.Unmarshal(body, &requestBody)

	roomId, err := rh.roomUseCase.CreateRoom(requestBody.Name, user)
	if err != nil {
		response.InternalServerError(writer, err.Error())
		return
	}

	res := createResponse{
		RoomId: roomId,
	}
	response.Success(writer, res)
}

func (rh roomHandler) HandleRoomShow(writer http.ResponseWriter, request *http.Request) {
	_, err := authentication.SessionUser(request)
	if err != nil {
		// TODO: redirect login form
		response.Unauthorized(writer, "Invalid Session")
		return
	}

	vars := mux.Vars(request)
	roomId := vars["id"]

	room, err := rh.roomUseCase.ShowRoom(roomId)
	if err != nil {
		response.InternalServerError(writer, "Internal Server Error")
		return
	}

	response.Success(writer, room)
}

func (rh roomHandler) HandleRoomJoin(writer http.ResponseWriter, request *http.Request) {
	user, err := authentication.SessionUser(request)
	if err != nil {
		// TODO: redirect login form
		response.Unauthorized(writer, "Invalid Session")
		return
	}

	// パスパラメータを取得
	vars := mux.Vars(request)
	id := vars["id"]

	err = rh.roomUseCase.JoinRoom(id, user)
	if err != nil {
		response.InternalServerError(writer, err.Error())
		return
	}

	response.Success(writer, "")
}
