package handler

import "github.com/anish-yadav/lms-api/internal/pkg/user"

type StudentResponse struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	RegistrationNo string `json:"registration_no"`
	Role           string `json:"string"`
}

type UserRequest struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Type  string `json:"type" validate:"required"`
}

type UserRequestResponse struct {
	ID string `json:"id"`
}

type ResetPasswordRequest struct {
	ID          string `json:"id" validate:"required"`
	OldPassword string `json:"old" validate:"required"`
	NewPassword string `json:"new" validate:"required"`
}

func (u *UserRequest) toUser() user.User {
	return user.User{
		Name:  u.Name,
		Email: u.Email,
		Type:  u.Type,
	}
}
