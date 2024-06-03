package photos_route

import (
	"github.com/gin-gonic/gin"
	"github.com/zidanhafiz/go-restapi/controllers/photos_controller"
	"github.com/zidanhafiz/go-restapi/middlewares"
	"github.com/zidanhafiz/go-restapi/middlewares/photos_middleware"
)

func Routes(router *gin.Engine) {
	r := router.Group("/photos")
	{
		r.Use(middlewares.Authentication())
		r.GET("/", photos_controller.Index)
		r.POST("/", photos_controller.Create)
		r.Use(photos_middleware.Authorization())
		r.GET("/:photoId", photos_controller.Show)
		r.PUT("/:photoId", photos_controller.Update)
		r.DELETE("/:photoId", photos_controller.Delete)
	}
}
