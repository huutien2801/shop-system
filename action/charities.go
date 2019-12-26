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

func FindAllCharity(input models.Charity, limit int64, offset int64) []*models.Charity {

	//Set query
	findOptions := options.Find()
	findOptions.SetLimit(limit)
	findOptions.SetSkip(offset)
	filter := bson.M{}

	if input.CharityCode != "" {
		filter["charity_code"] = input.CharityCode
	}
	if input.CharityName != "" {
		filter["charity_name"] = input.CharityName
	}
	if input.Status != "" {
		filter["status"] = input.Status
	}
	if input.Address != "" {
		filter["address"] = input.Address
	}
	if input.Target != 0 {
		filter["target"] = input.Target
	}
	if input.CurrentBudget != 0 {
		filter["current_budget"] = input.Target
	}
	if input.StartTime != nil {
		filter["start_time"] = bson.M{
			"$gte": input.StartTime,
		}
	}
	if input.FinishTime != nil {
		filter["finish_time"] = bson.M{
			"$lte": input.FinishTime,
		}
	}

	var results []*models.Charity
	cur, err := models.CharityDB.Collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem models.Charity
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

func CreateCharity(newCharity models.Charity) *mongo.InsertOneResult {

	insertResult, err := models.CharityDB.Collection.InsertOne(context.TODO(), newCharity)
	if err != nil {
		log.Fatal(err)
	}
	return insertResult
}

func DeleteCharity(id string) *mongo.DeleteResult {
	objectId, _ := primitive.ObjectIDFromHex(id)
	deleteResult, err := models.CharityDB.Collection.DeleteOne(context.TODO(), bson.M{"_id": objectId})
	if err != nil {
		log.Fatal(err)
	}
	return deleteResult
	// fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
}

func UpdateCharity(id string, newUpdater models.Charity) *mongo.UpdateResult {

	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id": objectId,
	}

	bsonUpdate := bson.M{}
	if newUpdater.CharityCode != "" {
		bsonUpdate["charity_code"] = newUpdater.CharityCode
	}
	if newUpdater.CharityName != "" {
		bsonUpdate["charity_name"] = newUpdater.CharityName
	}
	if newUpdater.Status != "" {
		bsonUpdate["status"] = newUpdater.Status
	}
	if newUpdater.Address != "" {
		bsonUpdate["address"] = newUpdater.Address
	}
	if newUpdater.Target != 0 {
		bsonUpdate["target"] = newUpdater.Target
	}
	if newUpdater.CurrentBudget != 0 {
		bsonUpdate["current_budget"] = newUpdater.Target
	}
	if newUpdater.StartTime != nil {
		bsonUpdate["start_time"] = newUpdater.StartTime
	}
	if newUpdater.FinishTime != nil {
		bsonUpdate["finish_time"] = newUpdater.FinishTime
	}

	update := bson.M{
		"$set": bsonUpdate,
	}
	// var res Charity
	updateResult, err := models.CharityDB.Collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	return updateResult
}

func FindOneCharity(id string) models.Charity {
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id": objectId,
	}
	var result models.Charity
	err := models.CharityDB.Collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
