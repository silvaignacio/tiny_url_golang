package presenter

import "pet-system/domain/model"

type PetPresenter interface {
	ResponsePets(u []*model.TinyUrl) []*model.TinyUrl
	ResponseCreated(u *model.TinyUrl) *model.TinyUrl
}
