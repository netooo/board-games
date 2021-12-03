package usecase

import (
	_ "github.com/go-playground/validator"
	"github.com/netooo/board-games/app/domain/model"
	"github.com/netooo/board-games/app/domain/repository"
	validators "github.com/netooo/board-games/app/interfaces/validators/numeron"
)

type NumeronPlayerUseCase interface {
	SetCode(user *model.User, id string, code string) error
}

type numeronPlayerUseCase struct {
	numeronPlayerRepository repository.NumeronPlayerRepository
}

func NewNumeronPlayerUseCase(npr repository.NumeronPlayerRepository) NumeronPlayerUseCase {
	return &numeronPlayerUseCase{
		numeronPlayerRepository: npr,
	}
}

func (npu numeronPlayerUseCase) SetCode(user *model.User, id string, code string) error {
	// リクエストパラメータのバリデーション
	ValidateCode := &validators.SetCode{
		Code: code,
	}

	if err := validators.SetCodeValidate(ValidateCode); err != nil {
		return err
	}

	// domainを介してinfrastructureで実装した関数を呼び出す。
	// Persistence（Repository）を呼出
	err := npu.numeronPlayerRepository.SetCode(user, id, code)
	if err != nil {
		return err
	}

	return nil
}
