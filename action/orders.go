package action

import (
	"context"

	"log"

	"github.com/huutien2801/shop-system/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindAllOrder(input models.Order, limit int64, offset int64) []*models.Order {

	//Set query
	findOptions := options.Find()
	findOptions.SetLimit(limit)
	findOptions.SetSkip(offset)
	filter := bson.M{}

	if input.OrderCode != "" {
		filter["order_code"] = input.OrderCode
	}
	if input.Products != nil {
		filter["product"] = input.Products
	}
	if input.TotalPrice != 0 {
		filter["total_price"] = input.TotalPrice
	}
	if input.PurchaseTime != nil {
		filter["purchase_time"] = input.PurchaseTime
	}
	if input.DeliverTime != nil {
		filter["deliver_time"] = input.DeliverTime
	}
	if input.UserID != "" {
		filter["user_id"] = input.UserID
	}
	if input.UserName != "" {
		filter["user_name"] = input.UserName
	}
	if input.Status != "" {
		filter["status"] = input.Status
	}
	if input.DeliverAddress != "" {
		filter["deliver_address"] = input.DeliverAddress
	}

	var results []*models.Order
	cur, err := models.OrderDB.Collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem models.Order
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())
	return results
}

func CreateOrder(newOrder models.Order) *mongo.InsertOneResult {

	insertResult, err := models.OrderDB.Collection.InsertOne(context.TODO(), newOrder)
	if err != nil {
		log.Fatal(err)
	}
	return insertResult
}

func DeleteOrder(id string) *mongo.DeleteResult {
	objectId, _ := primitive.ObjectIDFromHex(id)
	deleteResult, err := models.OrderDB.Collection.DeleteOne(context.TODO(), bson.M{"_id": objectId})
	if err != nil {
		log.Fatal(err)
	}
	return deleteResult
	// fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
}

func UpdateOrder(id string, newUpdater models.Order) *mongo.UpdateResult {

	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id": objectId,
	}

	bsonUpdate := bson.M{}
	if newUpdater.DeliverAddress != "" {
		bsonUpdate["deliver_address"] = newUpdater.DeliverAddress
	}

	if newUpdater.UserName != "" {
		bsonUpdate["user_name"] = newUpdater.UserName
	}

	if newUpdater.DeliverTime != nil {
		bsonUpdate["deliver_time"] = newUpdater.DeliverTime
	}

	if newUpdater.Products != nil {
		bsonUpdate["product"] = newUpdater.Products
	}
	update := bson.M{
		"$set": bsonUpdate,
	}
	// var res Order
	updateResult, err := models.OrderDB.Collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	return updateResult
}

func FindOneOrder(id string) models.Order {
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id": objectId,
	}
	var result models.Order
	err := models.OrderDB.Collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
