package handler

import "github.com/gin-gonic/gin"

type TaskHandler interface {
	GetAll(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Toggle(ctx *gin.Context)
}
