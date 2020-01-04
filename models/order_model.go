package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Order struct {
	ID             primitive.ObjectID `json:"id" bson:"_id"`
	OrderCode      string             `json:"orderCode" bson:"order_code"`
	PhoneNumber    string             `json:"phoneNumber" bson:"phone_number"`
	Products       []ProductDetail    `json:"products" bson:"products"`
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
	Name         string `json:"name" bson:"name"`
	Price        int32  `json:"price" bson:"price"`
	Discount     int32  `json:"discount" bson:"discount"`
	Description  string `json:"description" bson:"description"`
	Status       string `json:"status" bson:"status"`
	Quantity     int32  `json:"quantity" bson:"quantity"`
	ImageUrl     string `json:"imageUrl" bson:"image_url"`
	Size         string `json:"size" bson:"size"`
	ColorName    string `json:"colorName" bson:"color_name"`
	Material     string `json:"material" bson:"material"`
	CategoryName string `json:"categoryName" bson:"category_name"`
	Amount       int32  `json:"amount" bson:"amount"`
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
	CANCLE     string //Hủy đơn
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
	CANCLE:     "CANCLE",
}

var OrderDB = DbModel{
	ColName: "orders",
	DbName:  "ShopDB",
}

func InitOrderDB(client *mongo.Client) {
	OrderDB.Collection = OrderDB.GetCollection(client)
}
