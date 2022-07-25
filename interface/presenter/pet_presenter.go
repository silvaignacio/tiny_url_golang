package presenter

import (
	"pet-system/config"
	"pet-system/domain/model"
	"pet-system/usecase/presenter"
)

type petPresenter struct{}

func NewUserPresenter() presenter.PetPresenter {
	return &petPresenter{}
}

func (up *petPresenter) ResponsePets(us []*model.TinyUrl) []*model.TinyUrl {
	config.ReadConfig()

	for _, u := range us {
		u.ShortUrl = "http://" + config.C.Server.BaseUrl + "/" + u.ShortUrl
	}
	return us
}
