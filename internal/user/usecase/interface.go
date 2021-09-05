package usecase

import "github.com/anish-yadav/lms-api/internal/pkg/user"

type HttpManager interface {
	CreateUser(user user.User) (string, error)
	ResetPassword(id string, old string, pwd string) error
}
