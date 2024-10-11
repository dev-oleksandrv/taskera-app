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

type TaskHandlerImpl struct {
	taskService  service.TaskService
	spaceService service.SpaceService
}

func NewTaskHandler(taskService service.TaskService, spaceService service.SpaceService) TaskHandler {
	return &TaskHandlerImpl{taskService, spaceService}
}

func (h TaskHandlerImpl) GetAll(ctx *gin.Context) {
	spaceID := utils.GetRouteParamUUID(ctx, "spaceID")
	if spaceID == uuid.Nil {
		return
	}

	listID := utils.GetRouteParamUUID(ctx, "listID")
	if listID == uuid.Nil {
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

	tasks, err := h.taskService.GetAllByListID(listID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "InternalServerError",
			Errors: map[string]interface{}{},
		})
		return
	}
	var resTasks []response.TaskDto
	for _, task := range tasks {
		resTasks = append(resTasks, response.TaskDto{
			ID:          task.ID,
			Content:     task.Content,
			FullContent: task.FullContent,
			Order:       task.Order,
			Completed:   task.Completed,
		})
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data: &response.TaskGetAllResponse{
			Tasks: resTasks,
		},
	})
}

func (h TaskHandlerImpl) Create(ctx *gin.Context) {
	var req request.TaskCreateRequest
	if err := utils.ExecHandlerValidation(ctx, &req); err != nil {
		return
	}

	spaceID := utils.GetRouteParamUUID(ctx, "spaceID")
	if spaceID == uuid.Nil {
		return
	}

	listID := utils.GetRouteParamUUID(ctx, "listID")
	if listID == uuid.Nil {
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

	task := domain.Task{
		Content:   req.Content,
		ListID:    listID,
		CreatorID: userData.ID,
	}
	if err := h.taskService.Create(&task); err != nil {
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
		Data: &response.TaskCreateResponse{
			Task: response.TaskDto{
				ID:          task.ID,
				Content:     task.Content,
				FullContent: task.FullContent,
				Order:       task.Order,
				Completed:   task.Completed,
			},
		},
	})
}

func (h TaskHandlerImpl) Update(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h TaskHandlerImpl) Delete(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h TaskHandlerImpl) Toggle(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
