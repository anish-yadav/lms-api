package permission

import (
	"errors"
	"github.com/anish-yadav/lms-api/internal/pkg/db"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

type Permission struct {
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
}

const collection = "permissions"

func NewPermission(name string, p []string) *Permission {
	return &Permission{name, p}
}

func GetPermissionByName(name string) *Permission {
	permissionDB, err := db.GetByPKey(collection, "name", name)
	if err != nil {
		return nil
	}
	bsonBytes, err := bson.Marshal(permissionDB)
	if err != nil {
		log.Debugf("getPermissionByName: marshal bson : %s", err.Error())
		return nil
	}
	var permission Permission
	if err = bson.Unmarshal(bsonBytes, &permission); err != nil {
		log.Debugf("getPermissionByName: unmarshal to permission: %s", err.Error())
		return nil
	}
	return &permission
}

func (p *Permission) AddToDB() (string, error) {
	bin, err := bson.Marshal(p)
	if err != nil {
		log.Debugf("failed to marshal permission: %s", err.Error())
		return "", errors.New("failed to marshal permission data")
	}
	var bsonData bson.D
	err = bson.Unmarshal(bin, &bsonData)
	return db.InsertOne(collection, bsonData)
}

func (p *Permission) HasPermission(permission string) bool {
	for _, p := range p.Permissions {
		if permission == p {
			return true
		}
	}
	return false
}
