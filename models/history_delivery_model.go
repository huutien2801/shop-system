package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type HistoryDelivery struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	UserName    string             `json:"username" bson:"username"`
	Shipper     string             `json:"shipper" bson:"shipper"`
	PickTime    *time.Time         `json:"pickTime" bson:"pick_time"`
	DeliverTime *time.Time         `json:"deliverTime" bson:"deliver_time"`
	Status      string             `json:"status" bson:"status"`
	OrderCode   string             `json:"orderCode" bson:"order_code"`
	CreatedTime *time.Time         `json:"createdTime" bson:"created_time"`
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
