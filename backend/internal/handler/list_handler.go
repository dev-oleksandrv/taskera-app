package handler

import "github.com/gin-gonic/gin"

type ListHandler interface {
	GetAll(c *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
