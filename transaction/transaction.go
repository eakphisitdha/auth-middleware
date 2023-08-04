package transaction

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ITransaction interface {
	Log(field string, values string, action string) error
}

type Transaction struct {
	collection *mongo.Collection
}

func NewTransaction(mongoDB *mongo.Database) ITransaction {
	return &Transaction{collection: mongoDB.Collection("logs")}
}

func (t *Transaction) Log(field string, value string, action string) error {

	// create log in MongoDB
	doc := bson.D{{Key: "time", Value: time.Now().Format("2-Jan-06 03:04PM")}, {Key: field, Value: value}, {Key: "action", Value: action}}

	_, err := t.collection.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
