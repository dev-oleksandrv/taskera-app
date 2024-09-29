package router

import (
	"dev-oleksandrv/taskera-app/internal/handler"
	"dev-oleksandrv/taskera-app/internal/repository"
	"dev-oleksandrv/taskera-app/internal/service"
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.RouterGroup) {
	r := repository.NewUserRepository()
	s := service.NewUserService(r)
	h := handler.NewUserHandler(s)

	ur := router.Group("/user")

	ur.POST("/register", h.Register)
	ur.POST("/login", h.Login)
}
