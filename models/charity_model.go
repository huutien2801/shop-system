package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Charity struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	CharityCode   string             `json:"charityCode" bson:"charity_code"`
	CharityName   string             `json:"charityName" bson:"charity_name"`
	Target        int32              `json:"target" bson:"target"`
	CurrentBudget int32              `json:"currentBudget" bson:"current_budget"`
	Status        string             `json:"status" bson:"status"`
	Address       string             `json:"address" bson:"address"`
	StartTime     *time.Time         `json:"startTime" bson:"start_time"`
	FinishTime    *time.Time         `json:"finishTime" bson:"finish_time"`
	Sponsor       *Sponsor           `json:"sponsor" bson:"sponsor"`
	CreatedTime   *time.Time         `json:"createdTime" bson:"created_time"`
}

type Sponsor struct {
	UserName string `json:"userName" bson:"user_name"`
	Price    int32  `json:"price" bson:"price"`
}

type CharityStatusEnum struct {
	ACTIVE   string
	EXPIRED  string
	INACTIVE string
}

var CharityStatus = CharityStatusEnum{
	ACTIVE:   "ACTIVE",
	EXPIRED:  "EXPIRED",
	INACTIVE: "INACTIVE",
}

var CharityDB = DbModel{
	ColName: "charities",
	DbName:  "ShopDB",
}

func InitCharityDB(client *mongo.Client) {
	CharityDB.Collection = CharityDB.GetCollection(client)
}
