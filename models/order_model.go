package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Order struct {
	ID             primitive.ObjectID `json:"id" bson:"_id"`
	OrderCode      string             `json:"orderCode" bson:"order_code"`
	Products       *ProductDetail     `json:"productCode" bson:"product_code"`
	TotalPrice     int32              `json:"totalPrice" bson:"total_price"`
	PurchaseTime   *time.Time         `json:"purchaseDate" bson:"purchase_date"`
	DeliverTime    *time.Time         `json:"deliverTime" bson:"deliver_time"`
	UserID         string             `json:"userID" bson:"user_id"`
	UserName       string             `json:"userName" bson:"user_name"`
	Shipper        string             `json:"shipper" bson:"shipper"`
	Status         string             `json:"status" bson:"status"`
	DeliverAddress string             `json:"deliverAddress" bson:"deliver_address"`
	CreatedTime    *time.Time         `json:"createdTime" bson:"created_time"`
}

type ProductDetail struct {
	ProductCode string `json:"productCode" bson:"product_code"`
	Amount      int    `json:"amount" bson:"amount"`
}

type OrderStatusEnum struct {
	DELIVERED  string //Đã giao cho khách hàng
	PICKED     string //Tài xế đã lấy hàng giao
	RETURNED   string //Tài xê trả hàng cho shop
	ORDERED    string //Khách hàng đã đặt hàng
	WRAPPED    string //Shop đã đóng gói
	DELIVERING string //Tài xế đang giao
	PICKING    string //Tài xê đang lấy hàng
	RETURNING  string //Tài xế đang trả hàng
}

var OrderStatus = OrderStatusEnum{
	DELIVERED:  "DELIVERED",
	PICKED:     "PICKED",
	RETURNED:   "RETURNED",
	ORDERED:    "ORDERED",
	WRAPPED:    "WRAPPED",
	DELIVERING: "DELIVERING",
	PICKING:    "PICKING",
	RETURNING:  "RETURNING",
}

var OrderDB = DbModel{
	ColName: "orders",
	DbName:  "ShopDB",
}

func InitOrderDB(client *mongo.Client) {
	OrderDB.Collection = OrderDB.GetCollection(client)
}
