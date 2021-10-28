package models

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"plant_monitor/database"
	"time"
)

type Location struct {
	Latitude  int `bson:"latitude,omitempty" json:"latitude"`
	Longitude int `bson:"longitude,omitempty" json:"longitude"`
}

type Plant struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Type_       string             `bson:"type__" json:"type__"`
	Name        string             `bson:"name" json:"name"`
	Location    Location           `bson:"location" json:"location"`
	Description string             `bson:"description" json:"description"`
}

type Sensor struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name string
	Unit string
}

type Measurement struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Timestamp time.Time          `bson:"timestamp" json:"timestamp"`
	Sensor    primitive.ObjectID `bson:"sensor" json:"sensor"`
}
