package handler

import (
	"dev-oleksandrv/taskera-app/internal/model/domain"
	"dev-oleksandrv/taskera-app/internal/model/http/request"
	"dev-oleksandrv/taskera-app/internal/model/http/response"
	"dev-oleksandrv/taskera-app/internal/service"
	"dev-oleksandrv/taskera-app/internal/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

type UserHandlerImpl struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &UserHandlerImpl{userService}
}

func (h *UserHandlerImpl) Register(ctx *gin.Context) {
	var req request.UserRegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		var validationError validator.ValidationErrors
		errors.As(err, &validationError)
		errorMap := utils.ValidationErrorsToMap(validationError)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: errorMap,
		})
		return
	}

	user := domain.User{
		Email:    req.Email,
		Username: req.Username,
		Password: req.Password,
	}

	if err := h.userService.Register(&user); err != nil {
		fieldErrorResponse := make(map[string]interface{})

		if strings.Contains(err.Error(), "uni_users_email") {
			fieldErrorResponse["email"] = "Email is already used"
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: fieldErrorResponse,
		})

		return
	}

	ctx.JSON(http.StatusCreated, response.SuccessResponse{
		Code:   http.StatusCreated,
		Status: "Created",
		Data: response.UserRegisterResponse{
			ID:       user.ID,
			Email:    user.Email,
			Username: user.Username,
		},
	})
}

func (h *UserHandlerImpl) Login(ctx *gin.Context) {
	var req request.UserLoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		var validationError validator.ValidationErrors
		errors.As(err, &validationError)
		errorMap := utils.ValidationErrorsToMap(validationError)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: errorMap,
		})
		return
	}

	user := domain.User{
		Email:    req.Email,
		Password: req.Password,
	}

	if err := h.userService.Login(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
			Errors: err.Error(),
		})
		return
	}

	token := utils.GenerateJWTToken(user.ID, user.Email)
	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data: response.UserLoginResponse{
			Token: token,
		},
	})
}
