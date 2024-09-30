package router

import (
	"dev-oleksandrv/taskera-app/internal/handler"
	"dev-oleksandrv/taskera-app/internal/middleware"
	"dev-oleksandrv/taskera-app/internal/repository"
	"dev-oleksandrv/taskera-app/internal/service"
	"github.com/gin-gonic/gin"
)

func SpaceRouter(router *gin.RouterGroup) {
	spaceRepository := repository.NewSpaceRepository()
	spaceService := service.NewSpaceService(spaceRepository)

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)

	spaceHandler := handler.NewSpaceHandler(spaceService, userService)

	ur := router.Group("/space")

	ur.Use(middleware.AuthMiddleware())

	ur.GET("/", spaceHandler.GetAll)
	ur.POST("/", spaceHandler.Create)
	ur.POST("/invite/:spaceID", spaceHandler.Invite)
	ur.PUT("/:spaceID", spaceHandler.Update)
	ur.DELETE("/:spaceID", spaceHandler.Delete)
}
