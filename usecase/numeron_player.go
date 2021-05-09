package usecase

import (
	_ "github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/netooo/board-games/domain/model"
	"github.com/netooo/board-games/domain/repository"
	validators "github.com/netooo/board-games/interfaces/validators/numeron"
)

type NumeronPlayerUseCase interface {
	SetCode(code string) (*model.NumeronPlayer, error)
}

type numeronPlayerUseCase struct {
	numeronPlayerRepository repository.NumeronPlayerRepository
}

func NewNumeronPlayerUseCase(npr repository.NumeronPlayerRepository) NumeronPlayerUseCase {
	return &numeronPlayerUseCase{
		numeronPlayerRepository: npr,
	}
}

func (npu numeronPlayerUseCase) SetCode(code string) (*model.NumeronPlayer, error) {
	// リクエストパラメータのバリデーション
	ValidateCode := &validators.SetCode{
		Code: code,
	}

	if err := validators.SetCodeValidate(ValidateCode); err != nil {
		return nil, err
	}

	userId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	// domainを介してinfrastructureで実装した関数を呼び出す。
	// Persistence（Repository）を呼出
	user, err := uu.userRepository.Insert(userId.String(), name, email, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
