package manager

import (
	"errors"
	"github.com/anish-yadav/lms-api/internal/constants"
	"github.com/anish-yadav/lms-api/internal/pkg/db"
	"go.mongodb.org/mongo-driver/bson"
)

type httpManager struct {
}

func NewHttpManager() *httpManager {
	return &httpManager{}
}

func (s *httpManager) GetStudentByID(id string) (bson.M, error) {
	data, _ := db.GetByID("student", id)
	if data == nil {
		return nil, errors.New(constants.StudentNotFound)
	}

	return data, nil
}
