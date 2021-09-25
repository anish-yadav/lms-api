package user

import (
	"errors"
	"github.com/anish-yadav/lms-api/internal/pkg/db"
	"github.com/anish-yadav/lms-api/internal/util"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ResetRequest struct {
	ID             primitive.ObjectID `json:"id" bson:"_id"`
	Token          string             `json:"token"`
	ExpirationDate time.Time          `json:"expirationDate"`
	Used           bool               `json:"used"`
	Username       string             `json:"username"`
}

const resetTokenCollection = "reset-requests"

func NewPasswordResetRequest(user *UserDb) *ResetRequest {
	expirationTime := time.Now().Add(time.Hour * 24 * 31)
	return &ResetRequest{
		ExpirationDate: expirationTime,
		Username:       user.Email,
		Used:           false,
		Token:          uuid.NewString(),
		ID:             primitive.NewObjectID(),
	}
}

func GetReqById(id string) *ResetRequest {
	requestInDB, err := db.GetByID(resetTokenCollection, id)
	if err != nil {
		return nil
	}
	bsonBytes, err := bson.Marshal(requestInDB)
	if err != nil {
		log.Debugf("user.NewUserById: marshal bson : %s", err.Error())
		return nil
	}
	var req ResetRequest
	if err = bson.Unmarshal(bsonBytes, &req); err != nil {
		log.Debugf("req.GetReqByID: unmarshal to req: %s", err.Error())
		return nil
	}
	return &req
}

func (r *ResetRequest) SendRequest() (string, error) {
	bin, err := bson.Marshal(r)
	if err != nil {
		return "", errors.New("failed to marshal request data")
	}
	var bsonData bson.D
	err = bson.Unmarshal(bin, &bsonData)
	id, err := db.InsertOne(resetTokenCollection, bsonData)
	if err != nil {
		return "", err
	}
	data := map[string]string{
		"token_id": id,
		"username": r.Username,
	}
	token, err := util.CreateToken(data)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *ResetRequest) Close() error {
	log.Debugf("closing req")
	resetQuery := bson.D{{"$set", bson.D{{"used", true}}}}
	return db.UpdateItem(resetTokenCollection, r.ID.Hex(), resetQuery)
}

func (r *ResetRequest) IsValid() bool {
	if r.Used == true {
		return false
	}
	if time.Now().UnixNano() > r.ExpirationDate.UnixNano() {
		return false
	}
	return true
}
