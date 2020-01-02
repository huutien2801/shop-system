package action

import (
	"context"


	"github.com/huutien2801/shop-system/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindAllPromotion(input models.Promotion, limit int64, offset int64) models.Response {

	//Set query
	findOptions := options.Find()
	findOptions.SetLimit(limit)
	findOptions.SetSkip(offset)
	filter := bson.M{}

	if input.PromotionCode != "" {
		filter["promotion_code"] = input.PromotionCode
	}
	if input.Status != "" {
		filter["status"] = input.Status
	}
	if input.PromotionName != "" {
		filter["promotion_name"] = input.PromotionName
	}
	if input.PromotionType != "" {
		filter["promotion_type"] = input.PromotionType
	}
	if input.PromotionType != "" {
		filter["promotion_type"] = input.PromotionType
	}

	var results []*models.Promotion
	cur, err := models.PromotionDB.Collection.Find(context.TODO(), filter, findOptions)
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
		var elem models.Promotion
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
		Data: results,
		Message: "Success",
	}
}

func CreatePromotion(newPromotion models.Promotion) models.Response {

	insertResult, err := models.PromotionDB.Collection.InsertOne(context.TODO(), newPromotion)
	if err != nil {
		return models.Response{
			Status:  models.ResponseStatus.ERROR,
			Message: err.Error(),
		}
	}
	return models.Response{
		Status:  models.ResponseStatus.OK,
		Data: insertResult,
		Message: "Success",
	}
}

func DeletePromotion(id string) models.Response {
	objectID, _ := primitive.ObjectIDFromHex(id)
	deleteResult, err := models.PromotionDB.Collection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
	if err != nil {
		return models.Response{
			Status:  models.ResponseStatus.ERROR,
			Message: err.Error(),
		}
	}
	return models.Response{
		Status:  models.ResponseStatus.OK,
		Data: deleteResult,
		Message: "Success",
	}
	// fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
}

func UpdatePromotion(id string, newUpdater models.Promotion) models.Response {

	objectID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id": objectID,
	}

	bsonUpdate := bson.M{}
	if newUpdater.PromotionCode != "" {
		bsonUpdate["promotion_code"] = newUpdater.PromotionCode
	}
	if newUpdater.Status != "" {
		bsonUpdate["status"] = newUpdater.Status
	}
	if newUpdater.PromotionName != "" {
		bsonUpdate["promotion_name"] = newUpdater.PromotionName
	}
	if newUpdater.PromotionType != "" {
		bsonUpdate["promotion_type"] = newUpdater.PromotionType
	}
	if newUpdater.StartTime != nil {
		bsonUpdate["start_time"] = newUpdater.StartTime
	}
	if newUpdater.FinishTime != nil {
		bsonUpdate["finish_time"] = newUpdater.FinishTime
	}
	if newUpdater.ValueDiscount != nil {
		bsonUpdate["value_discount"] = newUpdater.ValueDiscount
	}

	update := bson.M{
		"$set": bsonUpdate,
	}
	// var res Promotion
	updateResult, err := models.PromotionDB.Collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return models.Response{
			Status:  models.ResponseStatus.ERROR,
			Message: err.Error(),
		}
	}
	return models.Response{
		Status:  models.ResponseStatus.OK,
		Data: updateResult,
		Message: "Success",
	}
}

func FindOnePromotion(id string) models.Response {
	objectID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id": objectID,
	}
	var result models.Promotion
	err := models.PromotionDB.Collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return models.Response{
			Status:  models.ResponseStatus.ERROR,
			Message: err.Error(),
		}
	}
	return models.Response{
		Status:  models.ResponseStatus.OK,
		Data: result,
		Message: "Success",
	}
}
