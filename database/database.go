package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"plant_monitor/configuration"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	client   *mongo.Client
	Db       *mongo.Database
	ctx      context.Context
	database *mongo.Database
}

var MI MongoInstance

func (db *MongoInstance) Connect() {

	var err error

	db.ctx = context.TODO()
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%d",
		configuration.ServerConfiguration.RootUsername,
		configuration.ServerConfiguration.RootPassword,
		configuration.ServerConfiguration.MongoUri,
		configuration.ServerConfiguration.Port))

	db.client, err = mongo.Connect(db.ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = db.client.Ping(db.ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	db.Db = db.client.Database(configuration.ServerConfiguration.DbName)
}

func (db *MongoInstance) DropAllCollections() {
	ctx, cancel := db.GetControllerContext()
	defer cancel()
	result, err := db.Db.ListCollectionNames(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	for i := range result {
		fmt.Println(i)
		err = db.Db.Collection(result[i]).Drop(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (db MongoInstance) ListDatabases() {
	databases, err := db.client.ListDatabases(db.ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)
}

func (db *MongoInstance) Disconnect() {
	db.client.Disconnect(db.ctx)
}

func (db *MongoInstance) GetControllerContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(db.ctx, 10*time.Second)
}
