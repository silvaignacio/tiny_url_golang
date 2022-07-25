package registry

import (
	"pet-system/interface/controllers"
	ip "pet-system/interface/presenter"
	ir "pet-system/interface/repository"
	"pet-system/usecase/interactor"
	up "pet-system/usecase/presenter"
	ur "pet-system/usecase/repository"
)

func (r *registry) NewUserController() controllers.PetController {
	return controllers.NewUserController(r.NewUserInteractor())
}

func (r *registry) NewUserInteractor() interactor.PetInteractor {
	return interactor.NewPetInteractor(r.NewUserRepository(), r.NewUserPresenter(), ir.NewDBRepository(r.db))
}

func (r *registry) NewUserRepository() ur.PetRepository {
	return ir.NewPetRepository(r.db)
}

func (r *registry) NewUserPresenter() up.PetPresenter {
	return ip.NewUserPresenter()
}
