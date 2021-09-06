package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/anish-yadav/lms-api/internal/constants"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
	"strings"
	"time"
)

var dbURI = ""
var dbName = ""

func Init(dbAddr string, db string) {
	dbURI = dbAddr
	dbName = db

	// ping the db once
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client := connect(ctx)
	err := client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Errorf("failed to ping database")
		os.Exit(1)
	}
	if err = client.Disconnect(ctx); err != nil {
		os.Exit(1)
	}
}

func CreateIndexes(col string) {
	//need to run once when setting up the account
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := connect(ctx)
	_, err := client.Database(dbName).Collection(col).Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true),
		})
	if err != nil {
		log.Errorf("Create indexes: %s", err.Error())
	}
}

func connect(ctx context.Context) *mongo.Client {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURI))
	if err != nil {
		log.Errorf("db.connect: %s", err.Error())
		os.Exit(1)
	}
	log.Debugf("db connection opened")
	return client
}

func GetByID(collNamespace string, id string) (bson.M, error) {
	log.Debugf("db.GeByID: %s , %s, %s", dbName, collNamespace, id)
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

	collection := client.Database(dbName).Collection(collNamespace)
	var result bson.M
	objectId, _ := primitive.ObjectIDFromHex(id)
	err := collection.FindOne(ctx, bson.D{{"_id", objectId}}).Decode(&result)
	if err != nil {
		log.Errorf("db.GetByID: %s", err.Error())
		return bson.M{}, err
	}
	return result, nil
}
func GetByPKey(collNamespace string, pkey string, value string) (bson.M, error) {
	log.Debugf("db.GeByID: %s , %s, %s", dbName, collNamespace, value)
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

	collection := client.Database(dbName).Collection(collNamespace)
	var result bson.M
	err := collection.FindOne(ctx, bson.D{{pkey, value}}).Decode(&result)
	if err != nil {
		log.Errorf("db.GetByPkey: %s", err.Error())
		return bson.M{}, err
	}
	return result, nil
}

func InsertOne(collNamespace string, data bson.D) (string, error) {
	log.Debugf("db.InsertOne: %s , %s", dbName, collNamespace)
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

	collection := client.Database(dbName).Collection(collNamespace)

	res, err := collection.InsertOne(ctx, data)

	if err != nil {
		log.Errorf("db.InsertOne: %s", err.Error())
		return "", err
	}
	log.Debugf("db.insertOne: id : %s", res.InsertedID)
	id := fmt.Sprintf("%s", res.InsertedID)
	id = strings.TrimPrefix(id, "ObjectID")
	id = strings.TrimPrefix(id, "(")
	id = strings.TrimSuffix(id, ")")
	id = strings.Trim(id, "\"")

	return id, nil
}

func UpdateItem(collNamespace string, id string, update bson.D) error {
	log.Debugf("db.Update: %s , %s, %s", dbName, collNamespace, id)
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

	collection := client.Database(dbName).Collection(collNamespace)
	//var result bson.M
	// do not creat new user if not present
	opts := options.Update().SetUpsert(false)
	result, err := collection.UpdateByID(ctx, id, update, opts)
	if err != nil {
		log.Errorf("db.GetByID: %s", err.Error())
		return err
	}
	if result.MatchedCount != 0 {
		return errors.New(constants.ItemNotFound)
	}
	return nil
}

func DelByID(collNamespace string, id string) error {
	log.Debugf("db.DelByID: %s , %s, %s", dbName, id)
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

	collection := client.Database(dbName).Collection(collNamespace)
	objectId, _ := primitive.ObjectIDFromHex(id)
	res, err := collection.DeleteOne(ctx, bson.D{{"_id", objectId}})
	if err != nil {
		log.Errorf("db.DelByID: %s", err.Error())
		return err
	}
	log.Debugf("db.DelByID: deleted entities: %s", res.DeletedCount)
	return nil
}
