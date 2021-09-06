package usecase

import (
	"errors"
	"github.com/anish-yadav/lms-api/internal/constants"
	"github.com/anish-yadav/lms-api/internal/pkg/user"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

type httpManager struct {
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func NewHttpManager() *httpManager {
	rand.Seed(time.Now().UnixNano())
	return &httpManager{}
}

func (m *httpManager) CreateUser(u user.User) (string, error) {
	password := randStringRunes(8)
	log.Debugf("email: %s, password: %s", u.Email, password)
	newUser := user.NewUser(u.Name, u.Email, password, u.Type)
	// TODO need to send mail to user to reset its password
	return newUser.AddToDB()
}

func (m *httpManager) ResetPassword(id string, old string, new string) error {
	currUser := user.GetUserById(id)
	if currUser == nil {
		return errors.New(constants.ItemNotFound)
	}
	return currUser.ResetPassword(old, new)
}

func (m *httpManager) DeleteUser(id string) error {
	return user.DeleteUserByID(id)
}

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
