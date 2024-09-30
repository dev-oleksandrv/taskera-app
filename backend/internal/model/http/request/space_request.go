package request

import "dev-oleksandrv/taskera-app/internal/model/domain"

type SpaceCreateRequest struct {
	Name        string `binding:"required" json:"name"`
	Description string `binding:"required" json:"description"`
}

type SpaceUpdateRequest struct {
	Name        string `binding:"required" json:"name"`
	Description string `binding:"required" json:"description"`
}

type SpaceInviteRequest struct {
	Email string      `binding:"required" json:"email"`
	Role  domain.Role `binding:"required" json:"role"`
}
