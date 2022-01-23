package usecase

import (
	"errors"
	"github.com/netooo/board-games/app/domain/model"
	"github.com/netooo/board-games/app/domain/repository"
	validators "github.com/netooo/board-games/app/interfaces/validators/numeron"
)

type NumeronUseCase interface {
	GetNumerons(userId string) ([]*model.Numeron, error)
	CreateNumeron(name string, userId string) (string, error)
	ShowNumeron(id string, userId string) (*model.Numeron, error)
	EntryNumeron(id string, userId string) error
	LeaveNumeron(id string, userId string) error
	StartNumeron(id string, userId string) error
}

type numeronUseCase struct {
	numeronRepository repository.NumeronRepository
}

func NewNumeronUseCase(r repository.NumeronRepository) NumeronUseCase {
	return &numeronUseCase{
		numeronRepository: r,
	}
}

func (u numeronUseCase) GetNumerons(userId string) ([]*model.Numeron, error) {
	numerons, err := u.numeronRepository.GetNumerons(userId)
	if err != nil {
		return nil, err
	}

	return numerons, nil
}

func (u numeronUseCase) CreateNumeron(name string, userId string) (string, error) {
	// リクエストパラメータのバリデーション
	validateNumeron := &validators.CreateNumeron{
		Name: name,
	}

	if err := validators.CreateNumeronValidate(validateNumeron); err != nil {
		return "", err
	}

	displayId, err := u.numeronRepository.CreateNumeron(name, userId)
	if err != nil {
		return "", err
	}

	return displayId, nil
}

func (u numeronUseCase) ShowNumeron(id string, userId string) (*model.Numeron, error) {
	if id == "" {
		return nil, errors.New("ID Not Found")
	}

	numeron, err := u.numeronRepository.ShowNumeron(id, userId)
	if err != nil {
		return nil, err
	}

	return numeron, nil
}

func (u numeronUseCase) EntryNumeron(id string, userId string) error {
	if id == "" {
		return errors.New("ID Not Found")
	}

	err := u.numeronRepository.EntryNumeron(id, userId)
	if err != nil {
		return err
	}

	return nil
}

func (u numeronUseCase) LeaveNumeron(id string, userId string) error {
	if id == "" {
		return errors.New("ID Not Found")
	}

	err := u.numeronRepository.LeaveNumeron(id, userId)
	if err != nil {
		return err
	}

	return nil
}

func (u numeronUseCase) StartNumeron(id string, userId string) error {
	if id == "" {
		return errors.New("ID Not Found")
	}

	err := u.numeronRepository.StartNumeron(id, userId)
	if err != nil {
		return err
	}

	return nil
}
