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

func FindAllHistory(input models.HistoryDelivery, limit int64, offset int64) []*models.HistoryDelivery {

	//Set query
	findOptions := options.Find()
	findOptions.SetLimit(limit)
	findOptions.SetSkip(offset)
	filter := bson.M{}

	if input.OrderCode != "" {
		filter["order_code"] = input.OrderCode
	}
	if input.Status != "" {
		filter["status"] = input.Status
	}
	if input.UserID != "" {
		filter["user_id"] = input.UserID
	}

	var results []*models.HistoryDelivery
	cur, err := models.HistoryDeliveryDB.Collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem models.HistoryDelivery
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

func CreateHistory(newHistory models.HistoryDelivery) *mongo.InsertOneResult {

	insertResult, err := models.HistoryDeliveryDB.Collection.InsertOne(context.TODO(), newHistory)
	if err != nil {
		log.Fatal(err)
	}
	return insertResult
}

func DeleteHistory(id string) *mongo.DeleteResult {
	objectId, _ := primitive.ObjectIDFromHex(id)
	deleteResult, err := models.HistoryDeliveryDB.Collection.DeleteOne(context.TODO(), bson.M{"_id": objectId})
	if err != nil {
		log.Fatal(err)
	}
	return deleteResult
	// fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
}

// func UpdateHistory(id string, newUpdater models.HistoryDelivery) *mongo.UpdateResult {

// 	objectId, _ := primitive.ObjectIDFromHex(id)
// 	filter := bson.M{
// 		"_id": objectId,
// 	}

// 	bsonUpdate := bson.M{}
// 	if newUpdater.Name != "" {
// 		bsonUpdate["name"] = newUpdater.Name
// 	}

// 	if newUpdater.Price != 0 {
// 		bsonUpdate["price"] = newUpdater.Price
// 	}
// 	update := bson.M{
// 		"$set": bsonUpdate,
// 	}
// 	// var res history
// 	updateResult, err := models.HistoryDeliveryDB.Collection.UpdateOne(context.TODO(), filter, update)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return updateResult
// }

func FindOneHistory(id string) models.HistoryDelivery {
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id": objectId,
	}
	var result models.HistoryDelivery
	err := models.HistoryDeliveryDB.Collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
