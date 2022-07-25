package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pet-system/interface/controllers"
)

func NewRouter(e *gin.Engine, c controllers.AppController) *gin.Engine {

	e.GET("/pets", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, c.Pet.GetPet(ctx)) })

	return e
}
