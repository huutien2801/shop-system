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

func FindAllUser(input models.User, limit int64, offset int64) []*models.User {

	//Set query
	findOptions := options.Find()
	findOptions.SetLimit(limit)
	findOptions.SetSkip(offset)
	filter := bson.M{}

	if input.FullName != "" {
		filter["full_name"] = bson.M{"$regex": input.FullName}
	}
	if input.Address != "" {
		filter["address"] = bson.M{"$regex": input.Address}
	}

	var results []*models.User
	cur, err := models.UserDB.Collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem models.User
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

func CreateUser(newUser models.User) *mongo.InsertOneResult {

	insertResult, err := models.UserDB.Collection.InsertOne(context.TODO(), newUser)
	if err != nil {
		log.Fatal(err)
	}
	return insertResult
}

func DeleteUser(id string) *mongo.DeleteResult {
	objectId, _ := primitive.ObjectIDFromHex(id)
	deleteResult, err := models.UserDB.Collection.DeleteOne(context.TODO(), bson.M{"_id": objectId})
	if err != nil {
		log.Fatal(err)
	}
	return deleteResult
	// fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
}

func UpdateUser(id string, newUpdater models.User) *mongo.UpdateResult {

	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id": objectId,
	}

	bsonUpdate := bson.M{}
	if newUpdater.FullName != "" {
		bsonUpdate["full_name"] = newUpdater.FullName
	}

	if newUpdater.PhoneNumber != "" {
		bsonUpdate["phone_number"] = newUpdater.PhoneNumber
	}

	if newUpdater.Address != "" {
		bsonUpdate["address"] = newUpdater.Address
	}

	if newUpdater.DateOfBirth != nil {
		bsonUpdate["date_of_birth"] = newUpdater.DateOfBirth
	}
	update := bson.M{
		"$set": bsonUpdate,
	}
	// var res User
	updateResult, err := models.UserDB.Collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	return updateResult
}

func FindOneUser(id string) models.User {
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id": objectId,
	}
	var result models.User
	err := models.UserDB.Collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
