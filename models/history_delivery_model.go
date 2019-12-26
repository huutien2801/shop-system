package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type HistoryDelivery struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	UserID      string             `json:"userID" bson:"user_id"`
	PickTime    *time.Time         `json:"pickTime" bson:"pick_time"`
	DeliverTime *time.Time         `json:"deliverTime" bson:"deliver_time"`
	Status      string             `json:"status" bson:"status"`
	OrderCode   string             `json:"orderCode" bson:"order_code"`
}

type HistoryStatusEnum struct {
	SUCCESS string
	FAILED  string
}

var HistoryStatus = HistoryStatusEnum{
	SUCCESS: "SUCCESS",
	FAILED:  "FAILED",
}

var HistoryDeliveryDB = DbModel{
	ColName: "history_deliveries",
	DbName:  "ShopDB",
}

func InitHistoryDeliveryDB(client *mongo.Client) {
	HistoryDeliveryDB.Collection = HistoryDeliveryDB.GetCollection(client)
}
