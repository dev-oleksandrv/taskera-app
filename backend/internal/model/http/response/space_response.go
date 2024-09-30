package response

import (
	"dev-oleksandrv/taskera-app/internal/model/domain"
	"github.com/google/uuid"
)

type SpaceDto struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

type SpaceCreateResponse struct {
	Space SpaceDto `json:"space"`
}

type SpaceUpdateResponse struct {
	Space SpaceDto `json:"space"`
}

type SpaceWithRoleDto struct {
	Space SpaceDto    `json:"space"`
	Role  domain.Role `json:"role"`
}

type SpaceGetAllResponse struct {
	Spaces []SpaceWithRoleDto `json:"spaces"`
}
