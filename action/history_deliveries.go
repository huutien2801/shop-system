package action

import (
	"context"

	"github.com/huutien2801/shop-system/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindAllHistory(input models.HistoryDelivery, limit int64, offset int64) models.Response {

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
	if input.UserName != "" {
		filter["user_id"] = input.UserName
	}
	if input.Shipper != "" {
		filter["shipper"] = input.Shipper
	}
	var results []*models.HistoryDelivery
	cur, err := models.HistoryDeliveryDB.Collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return models.Response{
			Status:  models.ResponseStatus.ERROR,
			Message: err.Error(),
		}
	}
	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem models.HistoryDelivery
		err := cur.Decode(&elem)
		if err != nil {
			return models.Response{
				Status:  models.ResponseStatus.ERROR,
				Message: err.Error(),
			}
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		return models.Response{
			Status:  models.ResponseStatus.ERROR,
			Message: err.Error(),
		}
	}

	// Close the cursor once finished
	cur.Close(context.TODO())
	return models.Response{
		Status:  models.ResponseStatus.OK,
		Message: "Success",
		Data:    results,
	}
}

func CreateHistory(newHistory models.HistoryDelivery) models.Response {

	insertResult, err := models.HistoryDeliveryDB.Collection.InsertOne(context.TODO(), newHistory)
	if err != nil {
		return models.Response{
			Status:  models.ResponseStatus.ERROR,
			Message: err.Error(),
		}
	}
	return models.Response{
		Status:  models.ResponseStatus.OK,
		Message: "Success",
		Data:    insertResult,
	}
}

func DeleteHistory(id string) models.Response {
	objectId, _ := primitive.ObjectIDFromHex(id)
	deleteResult, err := models.HistoryDeliveryDB.Collection.DeleteOne(context.TODO(), bson.M{"_id": objectId})
	if err != nil {
		return models.Response{
			Status:  models.ResponseStatus.ERROR,
			Message: err.Error(),
		}
	}
	return models.Response{
		Status:  models.ResponseStatus.OK,
		Message: "Success",
		Data:    deleteResult,
	}
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
// 		return models.Response{
//			Status:  models.ResponseStatus.ERROR,
//			Message: err.Error(),
//		}
// 	}
// 	return updateResult
// }

func FindOneHistory(id string) models.Response {
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id": objectId,
	}
	var result models.HistoryDelivery
	err := models.HistoryDeliveryDB.Collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return models.Response{
			Status:  models.ResponseStatus.ERROR,
			Message: err.Error(),
		}
	}
	return models.Response{
		Status:  models.ResponseStatus.OK,
		Message: "Success",
		Data:    result,
	}
}
