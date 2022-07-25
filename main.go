package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"pet-system/config"
	"pet-system/infrastructure/datastore"
	"pet-system/infrastructure/router"
	"pet-system/registry"
)

func main() {
	config.ReadConfig()

	db := datastore.NewDB()
	r := registry.NewRegistry(db)

	g := gin.Default()

	e := router.NewRouter(g, r.NewAppController())

	fmt.Println("Server listen at http://localhost" + ":" + config.C.Server.Address)
	if err := e.Run(":" + config.C.Server.Address); err != nil {
		log.Fatalln(err)
	}
}
