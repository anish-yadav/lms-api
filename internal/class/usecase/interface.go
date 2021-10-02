package usecase

import (
	"github.com/anish-yadav/lms-api/internal/pkg/class"
)

type HttpManager interface {
	CreateClass(class *class.Class) (string, error)
	GetAllClass() ([]*class.Class, error)
}
