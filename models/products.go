package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Product struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	Name         string             `json:"name" bson:"name"`
	Price        int32              `json:"price" bson:"price"`
	Discount     int32              `json:"discount" bson:"discount"`
	Description  string             `json:"description" bson:"description"`
	Status       string             `json:"status" bson:"status"`
	ImageURL     string             `json:"imageUrl" bson:"image_url"`
	ImageName    string             `json:"imageName" bson:"image_name"`
	Size         []string           `json:"size" bson:"size"`
	ColorName    []string           `json:"colorName" bson:"color_name"`
	ColorCode    []string           `json:"colorCode" bson:"color_code"`
	Material     string             `json:"material" bson:"material"`
	Rate         int32              `json:"rate" bson:"rate"`
	CategoryCode string             `json:"categoryCode" bson:"category_code"`
	CategoryName string             `json:"categoryName" bson:"category_name"`
	SelledAmount int32              `json:"selledAmount" bson:"selled_amount"`
	CreatedTime  *time.Time         `json:"createdTime" bson:"created_time"`
}

type ClientProductInput struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	Name         string             `json:"name" bson:"name"`
	Price        int32              `json:"price" bson:"price"`
	Description  string             `json:"description" bson:"description"`
	Status       string             `json:"status" bson:"status"`
	ImageURL     string             `json:"imageUrl" bson:"image_url"`
	ImageName    string             `json:"imageName" bson:"image_name"`
	CategoryCode string             `json:"categoryCode" bson:"category_code"`
	CategoryName string             `json:"categoryName" bson:"category_name"`
	SelledAmount int32              `json:"selledAmount" bson:"selled_amount"`
	ActionFilter string             `json:"actionFilter" bson:"action_filter"`
	CreatedTime  *time.Time         `json:"createdTime" bson:"created_time"`
}

type ActionEnum struct {
	PRICE_ASC  string
	PRICE_DESC string
	TIME_ASC   string
	TIME_DESC  string
	TOP_SELLER string
}

var ActionType = ActionEnum{
	PRICE_ASC:  "PRICE_ASC",
	PRICE_DESC: "PRICE_DESC",
	TIME_ASC:   "TIME_ASC",
	TIME_DESC:  "TIME_DESC",
	TOP_SELLER: "TOP_SELLER",
}

var ProductDB = DbModel{
	ColName: "products",
	DbName:  "ShopDB",
}

func InitProductDB(client *mongo.Client) {
	ProductDB.Collection = ProductDB.GetCollection(client)
}
