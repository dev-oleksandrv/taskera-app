package request

type ListCreateRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Emoji       string `json:"emoji" binding:"required"`
}

type ListUpdateRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Emoji       string `json:"emoji" binding:"required"`
}
