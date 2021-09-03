package manager

import "go.mongodb.org/mongo-driver/bson"

type HttpManager interface {
	GetStudentByID(id string) (bson.M, error)
}
