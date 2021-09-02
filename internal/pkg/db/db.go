package db

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var dbURI = ""

func Init(db string) {
	dbURI = db
}

func connect(ctx context.Context) *mongo.Client {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURI))
	if err != nil {
		log.Errorf("db.connect: %s", err.Error())
		panic(err)
	}
	log.Debugf("db connection opened")
	return client
}

func GetByID(dbNamespace string, collNamespace string, id string) bson.M {
	log.Debugf("db.GeByID: %s , %s, %s", dbNamespace, collNamespace, id)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client := connect(ctx)

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Errorf("failed to close db connection")
			panic(err)
		}
		log.Debugf("db connection closed")
	}()

	collection := client.Database(dbNamespace).Collection(collNamespace)
	var result bson.M
	err := collection.FindOne(ctx, bson.D{{"_id", id}}).Decode(&result)
	if err != nil {
		log.Errorf("db.GetByID: %s", err.Error())
		return bson.M{}
	}
	return result
}

func InsertOne(dbNamespace string, collNamespace string, data bson.D) error {
	log.Debugf("db.InsertOne: %s , %s", dbNamespace, collNamespace)
	log.Debugf("db.InsertOne: %+v", data)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client := connect(ctx)

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Errorf("failed to close db connection")
			panic(err)
		}
		log.Debugf("db connection closed")
	}()

	collection := client.Database(dbNamespace).Collection(collNamespace)

	res, err := collection.InsertOne(ctx, data)

	if err != nil {
		log.Errorf("db.GetByID: %s", err.Error())
		return err
	}
	fmt.Printf("db.insertOne: id : %s", res.InsertedID)
	return nil
}
