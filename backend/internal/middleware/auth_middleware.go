package middleware

import (
	"dev-oleksandrv/taskera-app/internal/model/http/response"
	"dev-oleksandrv/taskera-app/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

const UserDataKey = "userData"

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := utils.ExtractBearerToken(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse{
				Code:   http.StatusUnauthorized,
				Status: "Unauthorized",
				Errors: err.Error(),
			})
			return
		}
		claims, err := utils.ParseJWTToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse{
				Code:   http.StatusUnauthorized,
				Status: "Unauthorized",
				Errors: err.Error(),
			})
			return
		}

		ctx.Set(UserDataKey, claims)
		ctx.Next()
	}
}
