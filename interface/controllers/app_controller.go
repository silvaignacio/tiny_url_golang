package controllers

type AppController struct {
	Pet interface{ PetController }
}
