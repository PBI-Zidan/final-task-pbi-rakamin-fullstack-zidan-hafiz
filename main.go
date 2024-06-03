package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zidanhafiz/go-restapi/models"
	"github.com/zidanhafiz/go-restapi/routes/photos_route"
	"github.com/zidanhafiz/go-restapi/routes/users_route"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	photos_route.Routes(r)
	users_route.Routes(r)

	r.Run()
}
