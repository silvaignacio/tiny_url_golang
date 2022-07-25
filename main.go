package main

import (
	"fmt"
	"log"
	"pet-system/config"
	"pet-system/infrastructure/datastore"
	"pet-system/infrastructure/router"
	"pet-system/registry"

	"github.com/labstack/echo"
)

func main() {
	config.ReadConfig()

	db := datastore.NewDB()
	r := registry.NewRegistry(db)

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	fmt.Println("Server listen at http://localhost" + ":" + config.C.Server.Address)
	if err := e.Start(":" + config.C.Server.Address); err != nil {
		log.Fatalln(err)
	}
}
