package models
import (
    "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type Product struct {
    ID  primitive.ObjectID `json:"id" bson:"_id"`    
    Name string  `bson:"name"`
    Price  int32   `bson:"price"`
}

var ProductDB = DbModel{
	ColName: "products",
	DbName: "ShopDB",
}

func InitProductDB(client *mongo.Client){
	ProductDB.Collection =  ProductDB.GetCollection(client)
}