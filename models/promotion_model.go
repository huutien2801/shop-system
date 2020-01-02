package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Promotion struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	PromotionCode string             `json:"promotionCode" bson:"promotion_code"`
	PromotionName string             `json:"promotionName" bson:"promotion_name"`
	PromotionType string             `json:"promotionType" bson:"promotion_type"`
	Status        string             `json:"status" bson:"status"`
	StartTime     *time.Time         `json:"startTime" bson:"start_time"`
	FinishTime    *time.Time         `json:"finishTime" bson:"finish_time"`
	ValueDiscount *int32             `json:"valueDiscount" bson:"value_discount"`
	CreatedTime *time.Time         `json:"createdTime" bson:"created_time"`
}

type PromotionTypeEnum struct {
	DISCOUNT  string
	FREE_SHIP string
}
type PromotionStatusEnum struct {
	ACTIVE   string
	EXPIRED  string
	INACTIVE string
}

var PromotionStatus = PromotionStatusEnum{
	ACTIVE:   "ACTIVE",
	EXPIRED:  "EXPIRED",
	INACTIVE: "INACTIVE",
}

var PromotionType = PromotionTypeEnum{
	DISCOUNT:  "DISCOUNT",
	FREE_SHIP: "FREE_SHIP",
}

var PromotionDB = DbModel{
	ColName: "promotions",
	DbName:  "ShopDB",
}

func InitPromotionDB(client *mongo.Client) {
	PromotionDB.Collection = PromotionDB.GetCollection(client)
}
