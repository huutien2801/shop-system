package models

import (
	"time"

	"github.com/patrickmn/go-cache"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Username    string             `json:"username" bson:"username"`
	Password    string             `json:"password" bson:"password"`
	DateOfBirth *time.Time         `json:"dateOfBirth" bson:"date_of_birth"`
	FullName    string             `json:"fullName" bson:"full_name"`
	PhoneNumber string             `json:"phoneNumber" bson:"phone_number"`
	Address     AddressDetailEnum  `json:"address" bson:"address"`
	UserRole    string             `json:"userRole" bson:"user_role"`
	Email       string             `json:"email" bson:"email"`
	Avatar      string             `json:"avatar" bson:"avatar"`
	Session     string             `json:"session" bson:"session"`
}

type AddressDetailEnum struct {
	Address  string `json:"address" bson:"address"`
	Province string `json:"province" bson:"province"`
	District string `json:"district" bson:"district"`
}

type UserRoleEnum struct {
	SHIPPER  string
	CUSTOMER string
	ADMIN    string
}

var UserRoleType = UserRoleEnum{
	SHIPPER:  "SHIPPER",
	CUSTOMER: "CUSTOMER",
	ADMIN:    "ADMIN",
}

var UserDB = DbModel{
	ColName: "users",
	DbName:  "ShopDB",
}
var UserCache *cache.Cache

func InitUserDB(client *mongo.Client) {
	UserDB.Collection = UserDB.GetCollection(client)
}

func InitUserCache() {
	UserCache = cache.New(5*time.Minute, 10*time.Minute)
}
