package router

import (
	"dev-oleksandrv/taskera-app/internal/handler"
	"dev-oleksandrv/taskera-app/internal/middleware"
	"dev-oleksandrv/taskera-app/internal/repository"
	"dev-oleksandrv/taskera-app/internal/service"
	"github.com/gin-gonic/gin"
)

func ListRouter(router *gin.RouterGroup) {
	listRepository := repository.NewListRepository()
	listService := service.NewListService(listRepository)
	spaceRepository := repository.NewSpaceRepository()
	spaceService := service.NewSpaceService(spaceRepository)
	listHandler := handler.NewListHandler(listService, spaceService)

	r := router.Group("/space/:spaceID/list")

	r.Use(middleware.AuthMiddleware())

	r.GET("/", listHandler.GetAll)
	r.POST("/", listHandler.Create)
	r.DELETE("/:listID", listHandler.Delete)
	r.PUT("/:listID", listHandler.Update)
}
