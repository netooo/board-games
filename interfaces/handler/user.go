package handler

import (
	"encoding/json"
	"github.com/netooo/board-games/interfaces/authentication"
	"github.com/netooo/board-games/interfaces/response"
	"github.com/netooo/board-games/usecase"
	"io/ioutil"
	"net/http"
)

type UserHandler interface {
	HandleUserGet(http.ResponseWriter, *http.Request)
	HandleUserSignup(http.ResponseWriter, *http.Request)
}

type userHandler struct {
	userUseCase usecase.UserUseCase
}

type userSignupRequest struct {
	Name     string
	Email    string
	Password string
}

func NewUserHandler(uu usecase.UserUseCase) UserHandler {
	return &userHandler{
		userUseCase: uu,
	}
}

func (uh userHandler) HandleUserGet(writer http.ResponseWriter, request *http.Request) {
	// Contextから認証済みのユーザIdを取得
	userId := request.FormValue("userId")

	// UseCaseレイヤを操作して、ユーザデータ取得
	user, err := uh.userUseCase.GetByUserId(userId)
	if err != nil {
		response.InternalServerError(writer, "Internal Server Error")
		return
	}

	// レスポンスに必要な情報を詰めて返却
	response.Success(writer, user)
}

func (uh userHandler) HandleUserSignup(writer http.ResponseWriter, request *http.Request) {
	// リクエストボディを取得
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		response.BadRequest(writer, "Invalid Request Body")
		return
	}

	// リクエストボディのパース
	var requestBody userSignupRequest
	_ = json.Unmarshal(body, &requestBody)

	// UseCaseの呼び出し
	user, err := uh.userUseCase.Insert(requestBody.Name, requestBody.Email, requestBody.Password)
	if err != nil {
		response.InternalServerError(writer, "Internal Server Error")
		return
	}

	// Create and Return Session
	session := authentication.SessionCreate(user)
	_ = session.Save(request, writer)

	// レスポンスに必要な情報を詰めて返却
	response.Success(writer, "")
}
