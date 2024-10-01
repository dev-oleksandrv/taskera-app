package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func GetRouteParamUUID(ctx *gin.Context, paramKey string) uuid.UUID {
	spaceID := ctx.Param(paramKey)
	parsedSpaceID, err := uuid.Parse(spaceID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: err.Error(),
		})
		return uuid.Nil
	}
	return parsedSpaceID
}
