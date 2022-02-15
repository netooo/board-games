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

type NumeronHandler interface {
	HandleNumeronGet(http.ResponseWriter, *http.Request)
	HandleNumeronCreate(http.ResponseWriter, *http.Request)
	HandleNumeronShow(http.ResponseWriter, *http.Request)
	HandleNumeronEntry(http.ResponseWriter, *http.Request)
	HandleNumeronLeave(http.ResponseWriter, *http.Request)
	HandleNumeronStart(http.ResponseWriter, *http.Request)
	HandleNumeronSet(http.ResponseWriter, *http.Request)
	HandleNumeronAttack(http.ResponseWriter, *http.Request)
}

type numeronHandler struct {
	numeronUseCase usecase.NumeronUseCase
}

type createRequest struct {
	Name string
}

type startRequest struct {
	First  string
	Second string
}

type codeRequest struct {
	Code string
}

type getResponse struct {
	DisplayId string `json:"display_id"`
	Name      string `json:"name"`
	Owner     string `json:"owner"`
	UsersNum  int    `json:"users_num"`
}

type createResponse struct {
	DisplayId string `json:"display_id"`
}

type showResponse struct {
	DisplayId string             `json:"display_id"`
	Name      string             `json:"name"`
	Status    int                `json:"status"`
	Owner     string             `json:"owner"`
	Users     []showUserResponse `json:"users"`
}

type showUserResponse struct {
	UserId string `json:"user_id"`
	Name   string `json:"name"`
}

func NewNumeronHandler(u usecase.NumeronUseCase) NumeronHandler {
	return &numeronHandler{
		numeronUseCase: u,
	}
}

func (h numeronHandler) HandleNumeronGet(writer http.ResponseWriter, request *http.Request) {
	_, err := authentication.SessionUser(request)
	if err != nil {
		// TODO: redirect login form
		response.Unauthorized(writer, "Invalid Session")
		return
	}

	//TODO: Check request_user already join other numeron?
	// もしやるんだったら Userテーブルに Statusカラムを追加しないといけなさそう

	numerons, err := h.numeronUseCase.GetNumerons()
	if err != nil {
		response.InternalServerError(writer, "Internal Server Error")
		return
	}

	var res []*getResponse
	for _, r_ := range numerons {
		r := getResponse{
			DisplayId: r_.DisplayId,
			Name:      r_.Name,
			Owner:     r_.Owner.Name,
			UsersNum:  len(r_.Users),
		}
		res = append(res, &r)
	}

	response.Success(writer, res)
}

func (h numeronHandler) HandleNumeronCreate(writer http.ResponseWriter, request *http.Request) {
	user, err := authentication.SessionUser(request)
	if err != nil {
		// TODO: redirect login form
		response.Unauthorized(writer, "Invalid Session")
		return
	}

	//TODO: Check request_user already join other numeron?
	// もしやるんだったら Userテーブルに Statusカラムを追加しないといけなさそう

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		response.BadRequest(writer, "Invalid Request Body")
		return
	}

	// リクエストボディのパース
	var requestBody createRequest
	_ = json.Unmarshal(body, &requestBody)

	displayId, err := h.numeronUseCase.CreateNumeron(requestBody.Name, user.DisplayId)
	if err != nil {
		response.InternalServerError(writer, err.Error())
		return
	}

	res := createResponse{
		DisplayId: displayId,
	}
	response.Success(writer, res)
}

func (h numeronHandler) HandleNumeronShow(writer http.ResponseWriter, request *http.Request) {
	user, err := authentication.SessionUser(request)
	if err != nil {
		// TODO: redirect login form
		response.Unauthorized(writer, "Invalid Session")
		return
	}

	vars := mux.Vars(request)
	displayId := vars["display_id"]

	numeron, err := h.numeronUseCase.ShowNumeron(displayId, user.DisplayId)
	if err != nil {
		response.InternalServerError(writer, "Internal Server Error")
		return
	}

	var users []showUserResponse
	for u, _ := range numeron.Users {
		users = append(users, showUserResponse{
			u.DisplayId,
			u.Name,
		})
	}

	res := showResponse{
		DisplayId: numeron.DisplayId,
		Name:      numeron.Name,
		Status:    numeron.Status,
		Owner:     numeron.Owner.Name,
		Users:     users,
	}

	response.Success(writer, res)
}

func (h numeronHandler) HandleNumeronEntry(writer http.ResponseWriter, request *http.Request) {
	user, err := authentication.SessionUser(request)
	if err != nil {
		// TODO: redirect login form
		response.Unauthorized(writer, "Invalid Session")
		return
	}

	// パスパラメータを取得
	vars := mux.Vars(request)
	displayId := vars["display_id"]

	err = h.numeronUseCase.EntryNumeron(displayId, user.DisplayId)
	if err != nil {
		response.InternalServerError(writer, err.Error())
		return
	}

	response.Success(writer, "")
}

func (h numeronHandler) HandleNumeronLeave(writer http.ResponseWriter, request *http.Request) {
	user, err := authentication.SessionUser(request)
	if err != nil {
		// TODO: redirect login form
		response.Unauthorized(writer, "Invalid Session")
		return
	}

	// パスパラメータを取得
	vars := mux.Vars(request)
	displayId := vars["display_id"]

	err = h.numeronUseCase.LeaveNumeron(displayId, user.DisplayId)
	if err != nil {
		response.InternalServerError(writer, err.Error())
		return
	}

	response.Success(writer, "")
}

func (h numeronHandler) HandleNumeronStart(writer http.ResponseWriter, request *http.Request) {
	user, err := authentication.SessionUser(request)
	if err != nil {
		// TODO: redirect login form
		response.Unauthorized(writer, "Invalid Session")
		return
	}

	// パスパラメータを取得
	vars := mux.Vars(request)
	displayId := vars["display_id"]

	// リクエストボディを取得
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		response.BadRequest(writer, "Invalid Request Body")
		return
	}

	var requestBody startRequest
	_ = json.Unmarshal(body, &requestBody)

	err = h.numeronUseCase.StartNumeron(displayId, user.DisplayId, requestBody.First, requestBody.Second)
	if err != nil {
		response.InternalServerError(writer, err.Error())
		return
	}

	response.Success(writer, "")
}

func (h numeronHandler) HandleNumeronSet(writer http.ResponseWriter, request *http.Request) {
	user, err := authentication.SessionUser(request)
	if err != nil {
		// TODO: redirect login form
		response.Unauthorized(writer, "Invalid Session")
		return
	}

	// パスパラメータを取得
	vars := mux.Vars(request)
	displayId := vars["display_id"]

	// リクエストボディを取得
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		response.BadRequest(writer, "Invalid Request Body")
		return
	}

	var requestBody codeRequest
	_ = json.Unmarshal(body, &requestBody)

	err = h.numeronUseCase.SetNumeron(displayId, user.DisplayId, requestBody.Code)
	if err != nil {
		response.InternalServerError(writer, err.Error())
		return
	}

	response.Success(writer, "")
}

func (h numeronHandler) HandleNumeronAttack(writer http.ResponseWriter, request *http.Request) {
	user, err := authentication.SessionUser(request)
	if err != nil {
		// TODO: redirect login form
		response.Unauthorized(writer, "Invalid Session")
		return
	}

	// パスパラメータを取得
	vars := mux.Vars(request)
	displayId := vars["display_id"]

	// リクエストボディを取得
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		response.BadRequest(writer, "Invalid Request Body")
		return
	}

	var requestBody codeRequest
	_ = json.Unmarshal(body, &requestBody)

	err = h.numeronUseCase.AttackNumeron(displayId, user.DisplayId, requestBody.Code)
	if err != nil {
		response.InternalServerError(writer, err.Error())
		return
	}

	response.Success(writer, "")
}
