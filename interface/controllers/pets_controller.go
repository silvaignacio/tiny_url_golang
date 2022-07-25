package controllers

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"pet-system/domain/model"
	"pet-system/usecase/interactor"
)

type petController struct {
	userInteractor interactor.PetInteractor
}

func (uc *petController) FindByUrl(c *gin.Context) *model.TinyUrl {
	//TODO implement me
	data, err := uc.userInteractor.FindByUrl(c.Param("url"))
	if err != nil {
		panic("Error")
	}
	return data
}

type PostBody struct {
	Url string `json:"url"`
}

type ToCreate struct {
	ShortUrl string
	LongUrl  string
}

func (uc *petController) CreateShortUrl(c *gin.Context) *model.TinyUrl {

	var requestBody PostBody

	if err := c.BindJSON(&requestBody); err != nil {
		panic("Error al parsear body")
	}

	decodeUrl := base64.StdEncoding.EncodeToString([]byte(requestBody.Url))
	toCreate := &model.TinyUrl{}
	toCreate.ShortUrl = decodeUrl
	toCreate.LongUrl = requestBody.Url
	data, err := uc.userInteractor.Create(toCreate)
	if err != nil {
		panic("Error al crear data")
	}
	return data
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
	CreateShortUrl(c *gin.Context) *model.TinyUrl
	FindByUrl(c *gin.Context) *model.TinyUrl
}

func NewUserController(us interactor.PetInteractor) PetController {
	return &petController{us}
}
