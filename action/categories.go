package action

import (
	"context"
	"fmt"
	

	"github.com/huutien2801/shop-system/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindAllCategories() models.Response {

	findOptions := options.Find()
	findOptions.SetLimit(100)

	var results []*models.Category
	cur, err := models.CategoryDB.Collection.Find(context.TODO(), bson.M{}, findOptions)
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
		var elem models.Category
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
	for _, value := range results {
		fmt.Println(*value)
	}
	return models.Response{
		Status:  models.ResponseStatus.OK,
		Data: results,
		Message: "Success",
	}
}

func CreateCategory(newCategory models.Category) models.Response {

	insertResult, err := models.CategoryDB.Collection.InsertOne(context.TODO(), newCategory)

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

func DeleteCategory(id string) models.Response {
	objectId, _ := primitive.ObjectIDFromHex(id)
	deleteResult, err := models.CategoryDB.Collection.DeleteOne(context.TODO(), bson.M{"_id": objectId})
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

func UpdateCategory(id string, newUpdater models.Category) models.Response {

	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id": objectId,
	}

	bsonUpdate := bson.M{}
	if newUpdater.Name != "" {
		bsonUpdate["name"] = newUpdater.Name
	}

	update := bson.M{
		"$set": bsonUpdate,
	}
	// var res Product
	updateResult, err := models.CategoryDB.Collection.UpdateOne(context.TODO(), filter, update)
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

func FindOneCategory(id string) models.Response {
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id": objectId,
	}
	var result models.Category
	err := models.CategoryDB.Collection.FindOne(context.TODO(), filter).Decode(&result)
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
