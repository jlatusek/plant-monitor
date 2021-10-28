package models

import (
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"plant_monitor/database"
	"time"
)

type Location struct {
	Latitude  float32 `bson:"latitude,omitempty" json:"latitude"`
	Longitude float32 `bson:"longitude,omitempty" json:"longitude"`
}

type Plant struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Type_       string             `bson:"type" json:"type"`
	Name        string             `bson:"name" json:"name"`
	Location    Location           `bson:"location" json:"location"`
	Description string             `bson:"description" json:"description"`
}

type Sensor struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name string             `bson:"name" json:"name"`
	Unit string             `bson:"unit" json:"unit"`
}

type Measurement struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Timestamp time.Time          `bson:"timestamp" json:"timestamp"`
	Sensor    primitive.ObjectID `bson:"sensor" json:"sensor"`
}

func (plant *Plant) insert(db database.MongoInstance) *mongo.InsertOneResult {
	ctx, cancel := db.GetControllerContext()
	defer cancel()
	insertResult, err := db.Db.Collection("Plant").InsertOne(ctx, plant)
	if err != nil {
		panic(err)
	}
	return insertResult
}

func GetPlantById(db database.MongoInstance, id string) (*Plant, error) {
	ctx, cancel := db.GetControllerContext()
	defer cancel()
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return nil, errors.New("wrong id value")
	}
	var plant Plant
	if err = db.Db.Collection("Plant").FindOne(ctx, bson.M{"_id": objectId}).Decode(&plant); err != nil {
		log.Println(err)
		return nil, errors.New("plant with id not exist")
	}
	return &plant, nil
}

func GetPlantAll(db database.MongoInstance) ([]Plant, error) {
	ctx, cancel := db.GetControllerContext()
	defer cancel()
	cursor, err := db.Db.Collection("Plant").Find(ctx, bson.D{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	plants := make([]Plant, 0)
	for i := 0; cursor.Next(ctx); i++ {
		plants = append(plants, Plant{})
		var episode bson.M
		if err = cursor.Decode(&plants[i]); err != nil {
			log.Println(err)
			return nil, errors.New("decode object problem")
		}
		fmt.Println(episode)
	}
	err = cursor.Close(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return plants, nil
}

func (measurement Measurement) create() {

}
