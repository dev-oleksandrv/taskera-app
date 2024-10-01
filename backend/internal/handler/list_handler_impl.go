package handler

import (
	"dev-oleksandrv/taskera-app/internal/middleware"
	"dev-oleksandrv/taskera-app/internal/model/domain"
	"dev-oleksandrv/taskera-app/internal/model/http/request"
	"dev-oleksandrv/taskera-app/internal/model/http/response"
	"dev-oleksandrv/taskera-app/internal/service"
	"dev-oleksandrv/taskera-app/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type ListHandlerImpl struct {
	listService  service.ListService
	spaceService service.SpaceService
}

func NewListHandler(listService service.ListService, spaceService service.SpaceService) ListHandler {
	return &ListHandlerImpl{listService, spaceService}
}

func (h *ListHandlerImpl) GetAll(ctx *gin.Context) {
	spaceID := utils.GetRouteParamUUID(ctx, "spaceID")
	if spaceID == uuid.Nil {
		return
	}
	userData := ctx.MustGet(middleware.UserDataKey).(*utils.UserClaims)
	role := h.spaceService.GetSpaceRoleByUserID(spaceID, userData.ID)
	if role == nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, response.ErrorResponse{
			Code:   http.StatusForbidden,
			Status: "Forbidden",
			Errors: map[string]interface{}{},
		})
		return
	}
	lists, err := h.listService.GetAllBySpaceID(spaceID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "InternalServerError",
			Errors: map[string]interface{}{},
		})
		return
	}
	var resLists []response.ListDto
	for _, list := range lists {
		resLists = append(resLists, response.ListDto{
			ID:          list.ID,
			Name:        list.Name,
			Description: list.Description,
			Emoji:       list.Emoji,
			CreatorID:   list.CreatorID,
		})
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data: &response.ListGetAllResponse{
			Lists: resLists,
		},
	})
}

func (h *ListHandlerImpl) Create(ctx *gin.Context) {
	var req request.ListCreateRequest
	if err := utils.ExecHandlerValidation(ctx, &req); err != nil {
		return
	}

	spaceID := utils.GetRouteParamUUID(ctx, "spaceID")
	if spaceID == uuid.Nil {
		return
	}

	userData := ctx.MustGet(middleware.UserDataKey).(*utils.UserClaims)
	role := h.spaceService.GetSpaceRoleByUserID(spaceID, userData.ID)
	if role == nil || *role == domain.Editor {
		ctx.AbortWithStatusJSON(http.StatusForbidden, response.ErrorResponse{
			Code:   http.StatusForbidden,
			Status: "Forbidden",
			Errors: map[string]interface{}{},
		})
		return
	}

	list := domain.List{
		Name:        req.Name,
		Description: req.Description,
		Emoji:       req.Emoji,
		CreatorID:   userData.ID,
		SpaceID:     spaceID,
	}
	if err := h.listService.Create(&list); err != nil {
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
		Data: &response.ListCreateResponse{
			List: response.ListDto{
				ID:          list.ID,
				Name:        list.Name,
				Description: list.Description,
				Emoji:       list.Emoji,
				CreatorID:   list.CreatorID,
			},
		},
	})
}

func (h *ListHandlerImpl) Update(ctx *gin.Context) {
	var req request.ListUpdateRequest
	if err := utils.ExecHandlerValidation(ctx, &req); err != nil {
		return
	}

	spaceID := utils.GetRouteParamUUID(ctx, "spaceID")
	listID := utils.GetRouteParamUUID(ctx, "listID")
	if spaceID == uuid.Nil || listID == uuid.Nil {
		return
	}

	userData := ctx.MustGet(middleware.UserDataKey).(*utils.UserClaims)
	role := h.spaceService.GetSpaceRoleByUserID(spaceID, userData.ID)
	if role == nil || *role == domain.Editor {
		ctx.AbortWithStatusJSON(http.StatusForbidden, response.ErrorResponse{
			Code:   http.StatusForbidden,
			Status: "Forbidden",
			Errors: map[string]interface{}{},
		})
		return
	}

	list := domain.List{
		ID:          listID,
		Name:        req.Name,
		Description: req.Description,
		Emoji:       req.Emoji,
		CreatorID:   userData.ID,
		SpaceID:     spaceID,
	}
	if err := h.listService.Update(&list); err != nil {
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
		Data: &response.ListUpdateResponse{
			List: response.ListDto{
				ID:          list.ID,
				Name:        list.Name,
				Description: list.Description,
				Emoji:       list.Emoji,
				CreatorID:   list.CreatorID,
			},
		},
	})
}

func (h *ListHandlerImpl) Delete(ctx *gin.Context) {
	spaceID := utils.GetRouteParamUUID(ctx, "spaceID")
	listID := utils.GetRouteParamUUID(ctx, "listID")
	if spaceID == uuid.Nil || listID == uuid.Nil {
		return
	}

	userData := ctx.MustGet(middleware.UserDataKey).(*utils.UserClaims)
	role := h.spaceService.GetSpaceRoleByUserID(spaceID, userData.ID)
	if role == nil || *role == domain.Editor {
		ctx.AbortWithStatusJSON(http.StatusForbidden, response.ErrorResponse{
			Code:   http.StatusForbidden,
			Status: "Forbidden",
			Errors: map[string]interface{}{},
		})
		return
	}
	if err := h.listService.Delete(listID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "InternalServerError",
			Errors: map[string]interface{}{},
		})
		return
	}
	ctx.JSON(http.StatusNoContent, response.SuccessResponse{
		Code:   http.StatusNoContent,
		Status: "No Content",
	})
}
