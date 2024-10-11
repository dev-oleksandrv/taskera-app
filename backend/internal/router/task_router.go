package router

import (
	"dev-oleksandrv/taskera-app/internal/handler"
	"dev-oleksandrv/taskera-app/internal/middleware"
	"dev-oleksandrv/taskera-app/internal/repository"
	"dev-oleksandrv/taskera-app/internal/service"
	"github.com/gin-gonic/gin"
)

func TaskRouter(router *gin.RouterGroup) {
	taskRepository := repository.NewTaskRepository()
	taskService := service.NewTaskService(taskRepository)

	spaceRepository := repository.NewSpaceRepository()
	spaceService := service.NewSpaceService(spaceRepository)

	taskHandler := handler.NewTaskHandler(taskService, spaceService)

	r := router.Group("/space/:spaceID/list/:listID/task")

	r.Use(middleware.AuthMiddleware())

	r.GET("/", taskHandler.GetAll)
	r.POST("/", taskHandler.Create)
	r.PUT("/:taskID", taskHandler.Update)
	r.DELETE("/:taskID", taskHandler.Delete)
	r.POST("/:taskID/toggle", taskHandler.Toggle)
}
