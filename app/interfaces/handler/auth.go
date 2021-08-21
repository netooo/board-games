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

type AuthHandler interface {
	HandleAuthSignin(http.ResponseWriter, *http.Request)
}

type authHandler struct {
	authUseCase usecase.AuthUseCase
}

type authSignupRequest struct {
	Email    string
	Password string
}

func NewAuthHandler(ua usecase.AuthUseCase) AuthHandler {
	return &authHandler{
		authUseCase: ua,
	}
}

func (ah authHandler) HandleAuthSignup(writer http.ResponseWriter, request *http.Request) {
	// リクエストボディを取得
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		response.BadRequest(writer, "Invalid Request Body")
		return
	}

	// リクエストボディのパース
	var requestBody authSignupRequest
	_ = json.Unmarshal(body, &requestBody)

	// UseCaseの呼び出し
	user, err := ah.authUseCase.Signin(requestBody.Email, requestBody.Password)
	if err != nil {
		response.InternalServerError(writer, "Internal Server Error")
		return
	}

	// Create and Return Session
	session, err := authentication.SessionCreate(user.UserId)
	if err != nil {
		response.Unauthorized(writer, "Invalid Session")
	}
	_ = session.Save(request, writer)

	// レスポンスに必要な情報を詰めて返却
	response.Success(writer, "")
}
