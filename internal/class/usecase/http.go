package usecase

import (
	"errors"
	"github.com/anish-yadav/lms-api/internal/constants"
	"github.com/anish-yadav/lms-api/internal/pkg/class"
	log "github.com/sirupsen/logrus"
)

type httpManager struct {
}

func NewHttpManager() *httpManager {
	return &httpManager{}
}

func (m *httpManager) CreateClass(class *class.Class) (string, error) {
	id, err := class.AddToDB()
	if err != nil {
		log.Debugf("createClass: %s", err.Error())
		return "", err
	}
	return id, err
}

func (m *httpManager) GetAllClass() ([]*class.Class, error) {
	classes := class.GetAll()
	if len(classes) == 0 {
		return nil, errors.New(constants.ItemNotFound)
	}
	return classes, nil
}
