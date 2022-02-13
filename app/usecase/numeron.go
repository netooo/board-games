package usecase

import (
	"errors"
	"github.com/netooo/board-games/app/domain/model"
	"github.com/netooo/board-games/app/domain/repository"
	validators "github.com/netooo/board-games/app/interfaces/validators/numeron"
)

type NumeronUseCase interface {
	GetNumerons() ([]*model.Numeron, error)
	CreateNumeron(name string, userId string) (string, error)
	ShowNumeron(id string, userId string) (*model.Numeron, error)
	EntryNumeron(id string, userId string) error
	LeaveNumeron(id string, userId string) error
	StartNumeron(id string, userId string, firstId string, secondId string) error
	SetNumeron(id string, userId string, code string) error
	AttackNumeron(id string, userId string, code string) error
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

func (u numeronUseCase) StartNumeron(id string, userId string, firstId string, secondId string) error {
	if id == "" {
		return errors.New("ID Not Found")
	}

	if firstId == "" {
		return errors.New("First ID Not Found")
	}

	if secondId == "" {
		return errors.New("Second ID Not Found")
	}

	err := u.numeronRepository.StartNumeron(id, userId, firstId, secondId)
	if err != nil {
		return err
	}

	return nil
}

func (u numeronUseCase) SetNumeron(id string, userId string, code string) error {
	if id == "" {
		return errors.New("ID Not Found")
	}

	if len(code) != 3 {
		return errors.New("Length is Incorrect")
	}

	validateCode := &validators.SetCode{
		Code: code,
	}

	if err := validators.SetCodeValidate(validateCode); err != nil {
		return errors.New("Invalid Code")
	}

	err := u.numeronRepository.SetNumeron(id, userId, code)
	if err != nil {
		return err
	}

	return nil
}

func (u numeronUseCase) AttackNumeron(id string, userId string, code string) error {
	if id == "" {
		return errors.New("ID Not Found")
	}

	if len(code) != 3 {
		return errors.New("Length is Incorrect")
	}

	validateCode := &validators.AttackCode{
		Code: code,
	}

	if err := validators.AttackCodeValidate(validateCode); err != nil {
		return errors.New("Invalid Code")
	}

	err := u.numeronRepository.AttackNumeron(id, userId, code)
	if err != nil {
		return err
	}

	return nil
}
