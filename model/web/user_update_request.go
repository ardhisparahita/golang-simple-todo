package web

type UserUpdateRequest struct {
	Name  string `json:"name" validate:"required,min=3,max=255"`
	Email string `json:"email" validate:"required,email"`
}
