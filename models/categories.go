package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Category struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	Name string             `json:"name" bson:"name"`
	CreatedTime *time.Time  `json:"createdTime" bson:"created_time"`
}

var CategoryDB = DbModel{
	ColName: "categories",
	DbName:  "ShopDB",
}

func InitCategoryDB(client *mongo.Client) {
	CategoryDB.Collection = CategoryDB.GetCollection(client)
}
