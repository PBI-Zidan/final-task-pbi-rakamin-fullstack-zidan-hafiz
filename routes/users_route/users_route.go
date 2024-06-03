package users_route

import (
	"github.com/gin-gonic/gin"
	"github.com/zidanhafiz/go-restapi/controllers/users_controller"

	"github.com/zidanhafiz/go-restapi/middlewares"
	"github.com/zidanhafiz/go-restapi/middlewares/users_middleware"
)

func Routes(router *gin.Engine) {
	r := router.Group("/users")
	{
		r.POST("/register", users_controller.Register)
		r.POST("/login", users_controller.Login)
		r.Use(middlewares.Authentication())
		r.GET("/", users_controller.Index)
		r.Use(users_middleware.Authorization())
		r.GET("/:userId", users_controller.Show)
		r.PUT("/:userId", users_controller.Update)
		r.DELETE("/:userId", users_controller.Delete)
	}
}
