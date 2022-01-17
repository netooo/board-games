package usecase

import (
	"errors"
	"github.com/netooo/board-games/app/domain/model"
	"github.com/netooo/board-games/app/domain/repository"
	validators "github.com/netooo/board-games/app/interfaces/validators/numeron"
	"strconv"
)

type NumeronUseCase interface {
	GetNumerons() ([]*model.Numeron, error)
	CreateNumeron(name string, user *model.User) (uint, error)
	ShowNumeron(id string) (*model.Numeron, error)
	EntryNumeron(id string, user *model.User) error
	StartNumeron(id string, user *model.User) error
}

type numeronUseCase struct {
	numeronRepository repository.NumeronRepository
}

func NewNumeronUseCase(r repository.NumeronRepository) NumeronUseCase {
	return &numeronUseCase{
		numeronRepository: r,
	}
}

func (u numeronUseCase) GetNumerons() ([]*model.Numeron, error) {
	numerons, err := u.numeronRepository.GetNumerons()
	if err != nil {
		return nil, err
	}

	return numerons, nil
}

func (u numeronUseCase) CreateNumeron(name string, user *model.User) (uint, error) {
	// リクエストパラメータのバリデーション
	validateNumeron := &validators.CreateNumeron{
		Name: name,
	}

	if err := validators.CreateNumeronValidate(validateNumeron); err != nil {
		return 0, err
	}

	numeronId, err := u.numeronRepository.CreateNumeron(name, user)
	if err != nil {
		return 0, err
	}

	return numeronId, nil
}

func (u numeronUseCase) ShowNumeron(id string) (*model.Numeron, error) {
	if id == "" {
		return nil, errors.New("ID Not Found")
	}

	numeronId_, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, errors.New("Invalid ID")
	}

	var numeronId uint = uint(numeronId_)

	numeron, err := u.numeronRepository.ShowNumeron(numeronId)
	if err != nil {
		return nil, err
	}

	return numeron, nil
}

func (u numeronUseCase) EntryNumeron(id string, user *model.User) error {
	if id == "" {
		return errors.New("ID Not Found")
	}

	numeronId_, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return errors.New("Invalid ID")
	}

	var numeronId uint = uint(numeronId_)

	err = u.numeronRepository.EntryNumeron(numeronId, user)
	if err != nil {
		return err
	}

	return nil
}

func (u numeronUseCase) StartNumeron(id string, user *model.User) error {
	if id == "" {
		return errors.New("ID Not Found")
	}

	numeronId_, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return errors.New("Invalid ID")
	}

	var numeronId uint = uint(numeronId_)

	err = u.numeronRepository.StartNumeron(numeronId, user)
	if err != nil {
		return err
	}

	return nil
}
