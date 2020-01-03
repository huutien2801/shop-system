package action

import (
	"context"

	"github.com/huutien2801/shop-system/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindAllProduct(input models.ClientProductInput, limit int64, offset int64) models.Response {

	//Set query
	findOptions := options.Find()
	findOptions.SetLimit(limit)
	findOptions.SetSkip(offset)
	filter := bson.M{}

	if input.Name != "" {
		filter["name"] = input.Name
	}
	if input.Status != "" {
		filter["status"] = input.Status
	}

	if input.ActionFilter != "" {
		if input.ActionFilter == models.ActionType.PRICE_ASC {
			findOptions.SetSort(bson.M{"price": 1})
		}
		if input.ActionFilter == models.ActionType.PRICE_DESC {
			findOptions.SetSort(bson.M{"price": -1})
		}
		if input.ActionFilter == models.ActionType.TIME_ASC {
			findOptions.SetSort(bson.M{"created_time": 1})
		}
		if input.ActionFilter == models.ActionType.TIME_DESC {
			findOptions.SetSort(bson.M{"created_time": -1})
		}
		//TODO: Sort by time
	}

	var results []*models.Product
	cur, err := models.ProductDB.Collection.Find(context.TODO(), filter, findOptions)
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
		var elem models.Product
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
		Data:    results,
		Message: "Success",
	}
}

func CreateProduct(newProduct models.Product) models.Response {

	insertResult, err := models.ProductDB.Collection.InsertOne(context.TODO(), newProduct)
	if err != nil {
		return models.Response{
			Status:  models.ResponseStatus.ERROR,
			Message: err.Error(),
		}
	}
	return models.Response{
		Status:  models.ResponseStatus.OK,
		Data:    insertResult,
		Message: "Success",
	}
}

func DeleteProduct(id string) models.Response {
	objectId, _ := primitive.ObjectIDFromHex(id)
	deleteResult, err := models.ProductDB.Collection.DeleteOne(context.TODO(), bson.M{"_id": objectId})
	if err != nil {
		return models.Response{
			Status:  models.ResponseStatus.ERROR,
			Message: err.Error(),
		}
	}
	return models.Response{
		Status:  models.ResponseStatus.OK,
		Data:    deleteResult,
		Message: "Success",
	}
	// fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
}

func UpdateProduct(id string, newUpdater models.Product) models.Response {

	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id": objectId,
	}

	bsonUpdate := bson.M{}
	if newUpdater.Name != "" {
		bsonUpdate["name"] = newUpdater.Name
	}
	if newUpdater.Price != 0 {
		bsonUpdate["price"] = newUpdater.Price
	}
	if newUpdater.Discount != 0 {
		bsonUpdate["discount"] = newUpdater.Price
	}
	if newUpdater.Description != "" {
		bsonUpdate["description"] = newUpdater.Description
	}
	if newUpdater.ImageURL != "" {
		bsonUpdate["image_url"] = newUpdater.ImageURL
	}
	if newUpdater.ImageName != "" {
		bsonUpdate["image_name"] = newUpdater.ImageName
	}

	update := bson.M{
		"$set": bsonUpdate,
	}
	// var res Product
	updateResult, err := models.ProductDB.Collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return models.Response{
			Status:  models.ResponseStatus.ERROR,
			Message: err.Error(),
		}
	}
	return models.Response{
		Status:  models.ResponseStatus.OK,
		Data:    updateResult,
		Message: "Success",
	}
}

func FindOneProduct(id string) models.Response {
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id": objectId,
	}
	var result models.Product
	err := models.ProductDB.Collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return models.Response{
			Status:  models.ResponseStatus.ERROR,
			Message: err.Error(),
		}
	}
	return models.Response{
		Status:  models.ResponseStatus.OK,
		Data:    result,
		Message: "Success",
	}
}
