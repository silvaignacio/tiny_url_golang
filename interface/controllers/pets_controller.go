package controllers

import (
	"github.com/gin-gonic/gin"
	"pet-system/domain/model"
	"pet-system/usecase/interactor"
)

type petController struct {
	userInteractor interactor.PetInteractor
}

func (uc *petController) GetPet(c *gin.Context) []*model.TinyUrl {
	//TODO implement me
	data, err := uc.userInteractor.Get(nil)
	if err != nil {
		panic("Error")
	}
	return data
}

type PetController interface {
	GetPet(c *gin.Context) []*model.TinyUrl
}

func NewUserController(us interactor.PetInteractor) PetController {
	return &petController{us}
}
