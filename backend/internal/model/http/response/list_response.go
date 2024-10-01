package response

import (
	"github.com/google/uuid"
)

type ListDto struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Emoji       string    `json:"emoji"`
	CreatorID   uuid.UUID `json:"creator_id"`
}

type ListGetAllResponse struct {
	Lists []ListDto `json:"lists"`
}

type ListCreateResponse struct {
	List ListDto `json:"list"`
}

type ListUpdateResponse struct {
	List ListDto `json:"list"`
}
