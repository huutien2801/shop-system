package entities

import (
	"gopkg.in/mgo.v2/bson"
)

type Book struct {
	Id       bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string        `bson:"name"`
	Price    float64       `bson:"price"`
	Category string        `bson:"category"`
}
