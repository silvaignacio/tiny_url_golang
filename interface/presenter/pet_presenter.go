package presenter

import (
	"pet-system/domain/model"
	"pet-system/usecase/presenter"
)

type petPresenter struct{}

func NewUserPresenter() presenter.PetPresenter {
	return &petPresenter{}
}

func (up *petPresenter) ResponsePets(us []*model.TinyUrl) []*model.TinyUrl {
	for _, u := range us {
		u.ShortUrl = "Mr." + u.ShortUrl
	}
	return us
}
