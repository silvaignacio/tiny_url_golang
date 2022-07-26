package interactor

import (
	"pet-system/domain/model"
	"pet-system/usecase/presenter"
	"pet-system/usecase/repository"
)

type petInteractor struct {
	PetRepository repository.PetRepository
	PetPresenter  presenter.PetPresenter
	DBRepository  repository.DBRepository
}

func (us *petInteractor) FindByUrl(u string) (*model.TinyUrl, error) {
	data, err := us.PetRepository.FindByUrl(u)
	if err != nil {
		return nil, err
	}
	return us.PetPresenter.ResponseCreated(data), nil
}

func (us *petInteractor) Create(u *model.TinyUrl) (*model.TinyUrl, error) {
	data, err := us.PetRepository.Create(u)
	if err != nil {
		return nil, err
	}
	return us.PetPresenter.ResponseCreated(data), nil
}

type PetInteractor interface {
	Get(u []*model.TinyUrl) ([]*model.TinyUrl, error)
	Create(u *model.TinyUrl) (*model.TinyUrl, error)
	FindByUrl(u string) (*model.TinyUrl, error)
}

func NewPetInteractor(r repository.PetRepository, p presenter.PetPresenter, d repository.DBRepository) PetInteractor {
	return &petInteractor{r, p, d}
}

func (us *petInteractor) Get(u []*model.TinyUrl) ([]*model.TinyUrl, error) {
	u, err := us.PetRepository.FindAll(u)
	if err != nil {
		return nil, err
	}
	return us.PetPresenter.ResponsePets(u), nil
}
