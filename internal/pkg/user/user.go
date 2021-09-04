package user

import (
	"errors"
	"github.com/anish-yadav/lms-api/internal/pkg/db"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email    string      `json:"email"`
	Password string      `json:"password"`
	Type     string      `json:"type"`
	Detail   interface{} `json:"details"`
}

func NewUser(email string, password string, typ string) *User {
	hashedPwd, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return &User{email, string(hashedPwd), typ, nil}
}

func (user *User) AddToDB() error {
	bin, err := bson.Marshal(user)
	if err != nil {
		return errors.New("failed to marshal user data")
	}
	var bsonData bson.D
	err = bson.Unmarshal(bin, &bsonData)
	return db.InsertOne("user", bsonData)
}
