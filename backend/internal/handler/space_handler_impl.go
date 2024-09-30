package handler

import (
	"dev-oleksandrv/taskera-app/internal/middleware"
	"dev-oleksandrv/taskera-app/internal/model/domain"
	"dev-oleksandrv/taskera-app/internal/model/http/request"
	"dev-oleksandrv/taskera-app/internal/model/http/response"
	"dev-oleksandrv/taskera-app/internal/service"
	"dev-oleksandrv/taskera-app/internal/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"net/http"
)

type SpaceHandlerImpl struct {
	spaceService service.SpaceService
	userService  service.UserService
}

func NewSpaceHandler(spaceService service.SpaceService, userService service.UserService) SpaceHandler {
	return &SpaceHandlerImpl{spaceService, userService}
}

func (h *SpaceHandlerImpl) GetAll(ctx *gin.Context) {
	userData := ctx.MustGet(middleware.UserDataKey).(*utils.UserClaims)
	spaces, err := h.spaceService.GetAllByUser(userData.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: errors.New("cannot retrieve spaces"),
		})
		return
	}
	var resultSpaces []response.SpaceWithRoleDto
	for _, v := range spaces {
		resultSpaces = append(resultSpaces, response.SpaceWithRoleDto{
			Space: response.SpaceDto{
				ID:          v.ID,
				Name:        v.Name,
				Description: v.Description,
			},
			Role: v.Role,
		})
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data: &response.SpaceGetAllResponse{
			Spaces: resultSpaces,
		},
	})
}

func (h *SpaceHandlerImpl) Create(ctx *gin.Context) {
	var req request.SpaceCreateRequest
	if err := validateRequest(ctx, &req); err != nil {
		return
	}

	userData := ctx.MustGet(middleware.UserDataKey).(*utils.UserClaims)
	space := domain.Space{
		Name:        req.Name,
		Description: req.Description,
	}
	if err := h.spaceService.Create(&space, userData.ID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: map[string]interface{}{},
		})
		return
	}

	ctx.JSON(http.StatusCreated, response.SuccessResponse{
		Code:   http.StatusCreated,
		Status: "Created",
		Data: &response.SpaceCreateResponse{
			Space: response.SpaceDto{
				ID:          space.ID,
				Name:        space.Name,
				Description: space.Description,
			},
		},
	})
}

func (h *SpaceHandlerImpl) Update(ctx *gin.Context) {
	var req request.SpaceUpdateRequest
	if err := validateRequest(ctx, &req); err != nil {
		return
	}

	userData := ctx.MustGet(middleware.UserDataKey).(*utils.UserClaims)
	parsedSpaceID := parseSpaceID(ctx)
	role := h.spaceService.GetSpaceRoleByUserID(parsedSpaceID, userData.ID)
	if role == nil || *role != domain.Owner {
		ctx.AbortWithStatusJSON(http.StatusForbidden, response.ErrorResponse{
			Code:   http.StatusForbidden,
			Status: "Forbidden",
			Errors: errors.New("no permissions to update the space"),
		})
		return
	}
	space := domain.Space{
		ID:          parsedSpaceID,
		Name:        req.Name,
		Description: req.Description,
	}
	if err := h.spaceService.Update(&space); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: map[string]interface{}{},
		})
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data: &response.SpaceCreateResponse{
			Space: response.SpaceDto{
				ID:          space.ID,
				Name:        space.Name,
				Description: space.Description,
			},
		},
	})
}

func (h *SpaceHandlerImpl) Delete(ctx *gin.Context) {
	parsedSpaceID := parseSpaceID(ctx)

	userData := ctx.MustGet(middleware.UserDataKey).(*utils.UserClaims)
	role := h.spaceService.GetSpaceRoleByUserID(parsedSpaceID, userData.ID)
	if role == nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, response.ErrorResponse{
			Code:   http.StatusForbidden,
			Status: "Forbidden",
			Errors: errors.New("role is not found for current space"),
		})
		return
	}

	if *role != domain.Owner {
		ctx.AbortWithStatusJSON(http.StatusForbidden, response.ErrorResponse{
			Code:   http.StatusForbidden,
			Status: "Forbidden",
			Errors: errors.New("role is not owner of current space"),
		})
		return
	}

	if err := h.spaceService.Delete(parsedSpaceID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNoContent, response.SuccessResponse{
		Code:   http.StatusNoContent,
		Status: "NoContent",
	})
}

func (h *SpaceHandlerImpl) Invite(ctx *gin.Context) {
	var req request.SpaceInviteRequest
	if err := validateRequest(ctx, &req); err != nil {
		return
	}

	parsedSpaceID := parseSpaceID(ctx)

	userData := ctx.MustGet(middleware.UserDataKey).(*utils.UserClaims)
	role := h.spaceService.GetSpaceRoleByUserID(parsedSpaceID, userData.ID)
	if role == nil || *role != domain.Owner {
		ctx.AbortWithStatusJSON(http.StatusForbidden, response.ErrorResponse{
			Code:   http.StatusForbidden,
			Status: "Forbidden",
			Errors: errors.New("no permissions to update the space"),
		})
		return
	}

	invitedUser, err := h.userService.GetUserByEmail(req.Email)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: err.Error(),
		})
		return
	}

	invitedUserRole := h.spaceService.GetSpaceRoleByUserID(parsedSpaceID, invitedUser.ID)
	if invitedUserRole != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: errors.New("invited user is already a member of this space"),
		})
		return
	}

	spaceUser := domain.SpaceUser{
		UserID:  invitedUser.ID,
		SpaceID: parsedSpaceID,
		Role:    req.Role,
	}
	if err := h.spaceService.CreateSpaceUserRelation(&spaceUser); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNoContent, response.SuccessResponse{
		Code:   http.StatusNoContent,
		Status: "NoContent",
	})
}

func validateRequest[T interface{}](ctx *gin.Context, req *T) error {
	if err := ctx.ShouldBindJSON(&req); err != nil {
		var validationError validator.ValidationErrors
		errors.As(err, &validationError)
		errorMap := utils.ValidationErrorsToMap(validationError)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: errorMap,
		})
		return err
	}

	return nil
}

func parseSpaceID(ctx *gin.Context) uuid.UUID {
	spaceID := ctx.Param("spaceID")
	parsedSpaceID, err := uuid.Parse(spaceID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: err.Error(),
		})
		return uuid.Nil
	}
	return parsedSpaceID
}
