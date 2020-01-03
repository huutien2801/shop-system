package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Product struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	Price       int32              `json:"price" bson:"price"`
	Discount    int32              `json:"discount" bson:"discount"`
	Description string             `json:"description" bson:"description"`
	Status      string             `json:"status" bson:"status"`
	ImageURL    string             `json:"imageUrl" bson:"image_url"`
	ImageName   string             `json:"imageName" bson:"image_name"`
	CreatedTime *time.Time         `json:"createdTime" bson:"created_time"`
}

type ClientProductInput struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	Name         string             `json:"name" bson:"name"`
	Price        int32              `json:"price" bson:"price"`
	Description  string             `json:"description" bson:"description"`
	Status       string             `json:"status" bson:"status"`
	ImageURL     string             `json:"imageUrl" bson:"image_url"`
	ImageName    string             `json:"imageName" bson:"image_name"`
	ActionFilter string             `json:"actionFilter" bson:"action_filter"`
}

type ActionEnum struct {
	PRICE_ASC  string
	PRICE_DESC string
	TIME_ASC   string
	TIME_DESC  string
}

var ActionType = ActionEnum{
	PRICE_ASC:  "PRICE_ASC",
	PRICE_DESC: "PRICE_DESC",
	TIME_ASC:   "TIME_ASC",
	TIME_DESC:  "TIME_DESC",
}

var ProductDB = DbModel{
	ColName: "products",
	DbName:  "ShopDB",
}

func InitProductDB(client *mongo.Client) {
	ProductDB.Collection = ProductDB.GetCollection(client)
}
