package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pet-system/interface/controllers"
)

func NewRouter(e *gin.Engine, c controllers.AppController) *gin.Engine {

	e.GET("/:url", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, c.Pet.FindByUrl(ctx)) })
	e.POST("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, c.Pet.CreateShortUrl(ctx))
	})
	return e
}
