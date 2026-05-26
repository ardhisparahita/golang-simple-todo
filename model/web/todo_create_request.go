package web

type TodoCreateRequest struct {
	Title       string `json:"title" validate:"required,max=255"`
	Description string `json:"description"`
}
