package repository

import "pet-system/domain/model"

type PetRepository interface {
	FindAll(u []*model.TinyUrl) ([]*model.TinyUrl, error)
	Create(u *model.TinyUrl) (*model.TinyUrl, error)
	FindByUrl(u string) (*model.TinyUrl, error)
}
