package main

import (
	"dev-oleksandrv/taskera-app/internal/config"
	"dev-oleksandrv/taskera-app/internal/database"
	"dev-oleksandrv/taskera-app/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()
	database.Init()

	gin.SetMode(config.GetConfig().GinMode)
	r := gin.Default()
	g := r.Group("/api")
	router.UserRouter(g)

	if err := r.Run("localhost:8080"); err != nil {
		panic(err)
	}
}
