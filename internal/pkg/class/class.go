package class

import (
	"errors"
	"fmt"
	"github.com/anish-yadav/lms-api/internal/pkg/db"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

type Class struct {
	Name     string `json:"name"`
	Section  string `json:"section"`
	Semester string `json:"semester"`
	Stream   string `json:"stream"`
}

const collection = "classes"

func NewClass(section string, sem string, stream string) *Class {
	name := fmt.Sprintf("%s-%s (%s sem)", stream, section, sem)
	return &Class{name, section, sem, stream}
}
func GetClassByName(name string) *Class {
	classDb, err := db.GetByPKey(collection, "name", name)
	if err != nil {
		return nil
	}
	bsonBytes, err := bson.Marshal(classDb)
	if err != nil {
		log.Debugf("class.GetClassByName: marshal bson : %s", err.Error())
		return nil
	}
	var class Class
	if err = bson.Unmarshal(bsonBytes, &class); err != nil {
		log.Debugf("class.GetClassByName: unmarshal to class: %s", err.Error())
		return nil
	}
	return &class
}

func GetAll() []*Class {
	result, err := db.GetAll(collection)
	if err != nil {
		log.Debugf("err : %s", err.Error())
		return nil
	}
	var classes []*Class
	for _, classDb := range result {
		var class *Class
		bsonBytes, err := bson.Marshal(classDb)
		if err != nil {
			log.Debugf("class.GetAll: marshal bson : %s", err.Error())
			return nil
		}
		if err = bson.Unmarshal(bsonBytes, &class); err != nil {
			log.Debugf("class.GetAll: unmarshal to class: %s", err.Error())
			return nil
		}
		classes = append(classes, class)
	}
	return classes
}

func (c *Class) AddToDB() (string, error) {
	bin, err := bson.Marshal(c)
	if err != nil {
		log.Debugf("failed to marshal class: %s", err.Error())
		return "", errors.New("failed to marshal class data")
	}
	var bsonData bson.D
	err = bson.Unmarshal(bin, &bsonData)
	return db.InsertOne(collection, bsonData)
}
