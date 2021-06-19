package handler

import (
	"encoding/json"
	"github.com/netooo/board-games/app/interfaces/authentication"
	"github.com/netooo/board-games/app/interfaces/response"
	"github.com/netooo/board-games/app/usecase"
	"io/ioutil"
	"net/http"
)

type NumeronPlayerHandler interface {
	HandleNumeronSetCode(http.ResponseWriter, *http.Request)
}

type numeronPlayerHandler struct {
	numeronPlayerUseCase usecase.NumeronPlayerUseCase
}

type numeronSetCodeRequest struct {
	Code string
}

func NewNumeronPlayerHandler(npu usecase.NumeronPlayerUseCase) NumeronPlayerHandler {
	return &numeronPlayerHandler{
		numeronPlayerUseCase: npu,
	}
}

func (nph numeronPlayerHandler) HandleNumeronSetCode(writer http.ResponseWriter, request *http.Request) {
	user := authentication.SessionUser(request)
	if user == nil {
		// TODO: redirect login form
		response.Unauthorized(writer, "Invalid Session")
		return
	}

	// リクエストボディを取得
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		response.BadRequest(writer, "Invalid Request Body")
		return
	}

	// リクエストボディのパース
	var requestBody numeronSetCodeRequest
	_ = json.Unmarshal(body, &requestBody)

	// UseCaseの呼び出し
	err = nph.numeronPlayerUseCase.SetCode(user, requestBody.Code)
	if err != nil {
		response.InternalServerError(writer, "Internal Server Error")
		return
	}

	// レスポンスに必要な情報を詰めて返却
	response.Success(writer, "")
}
