package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/netooo/board-games/app/interfaces/authentication"
	"github.com/netooo/board-games/app/interfaces/response"
	"github.com/netooo/board-games/app/usecase"
	"io/ioutil"
	"net/http"
)

type UserHandler interface {
	HandleUserFind(http.ResponseWriter, *http.Request)
	HandleUserSignup(http.ResponseWriter, *http.Request)
	HandleUserSignin(http.ResponseWriter, *http.Request)
}

type userHandler struct {
	userUseCase usecase.UserUseCase
}

type userSignupRequest struct {
	Name     string
	Email    string
	Password string
}

type userSigninRequest struct {
	Email    string
	Password string
}

func NewUserHandler(uu usecase.UserUseCase) UserHandler {
	return &userHandler{
		userUseCase: uu,
	}
}

func (uh userHandler) HandleUserFind(writer http.ResponseWriter, request *http.Request) {
	// Contextから認証済みのユーザIdを取得
	vars := mux.Vars(request)
	userId := vars["user_id"]

	// UseCaseレイヤを操作して、ユーザデータ取得
	user, err := uh.userUseCase.FindByUserId(userId)
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
	session, err := authentication.SessionCreate(user.DisplayId)
	if err != nil {
		response.Unauthorized(writer, "Invalid Session")
	}

	// Set Cookie
	http.SetCookie(writer, sessions.NewCookie(session.Name(), session.ID, session.Options))

	// レスポンスに必要な情報を詰めて返却
	response.Success(writer, "")
}

func (uh userHandler) HandleUserSignin(writer http.ResponseWriter, request *http.Request) {
	_, err := authentication.SessionUser(request)
	if err == nil {
		response.Unauthorized(writer, "Already Logged In")
		return
	}

	// リクエストボディを取得
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		response.BadRequest(writer, "Invalid Request Body")
		return
	}

	// リクエストボディのパース
	var requestBody userSigninRequest
	_ = json.Unmarshal(body, &requestBody)

	// UseCaseの呼び出し
	user, err := uh.userUseCase.BasicSignin(requestBody.Email, requestBody.Password)
	if err != nil {
		response.InternalServerError(writer, "Internal Server Error")
		return
	}

	// Create and Return Session
	session, err := authentication.SessionCreate(user.DisplayId)
	if err != nil {
		response.Unauthorized(writer, "Invalid Session")
	}

	// Set Cookie
	http.SetCookie(writer, sessions.NewCookie(session.Name(), session.ID, session.Options))

	// レスポンスに必要な情報を詰めて返却
	response.Success(writer, "")
}
