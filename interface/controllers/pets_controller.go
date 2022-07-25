package controllers

import (
	"pet-system/usecase/interactor"
)

type petController struct {
	userInteractor interactor.PetInteractor
}

func (uc *petController) GetPet(c Context) error {
	//TODO implement me
	panic("implement me")
}

type PetController interface {
	GetPet(c Context) error
}

func NewUserController(us interactor.PetInteractor) PetController {
	return &petController{us}
}
