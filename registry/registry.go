package registry

import (
	"gorm.io/gorm"
	"pet-system/interface/controllers"
)

type registry struct {
	db *gorm.DB
}

type Registry interface {
	NewAppController() controllers.AppController
}

func NewRegistry(db *gorm.DB) Registry {
	return &registry{db}
}

func (r *registry) NewAppController() controllers.AppController {
	return controllers.AppController{
		Pet: r.NewUserController(),
	}
}
