package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type User struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Username    string             `json:"username" bson:"username"`
	Password    string             `json:"password" bson:"password"`
	DateOfBirth *time.Time         `json:"dateOfBirth" bson:"date_of_birth"`
	FullName    string             `json:"fullName" bson:"full_name"`
	PhoneNumber string             `json:"phoneNumber" bson:"phone_number"`
	Address     string             `json:"address" bson:"address"`
}

var UserDB = DbModel{
	ColName: "users",
	DbName:  "ShopDB",
}

func InitUserDB(client *mongo.Client) {
	UserDB.Collection = ProductDB.GetCollection(client)
}
