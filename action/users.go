package action

import (
	"context"
	"fmt"

	"log"

	"github.com/huutien2801/shop-system/models"
	"github.com/patrickmn/go-cache"
	uuid "github.com/satori/go.uuid"
	// uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

//FindAllUser function
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
	if input.UserRole != "" {
		filter["user_role"] = input.UserRole
	}
	if input.Email != "" {
		filter["email"] = input.Email
	}
	if input.PhoneNumber != "" {
		filter["phone_number"] = input.PhoneNumber
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

//CreateUser function
func CreateUser(newUser models.User) *mongo.InsertOneResult {

	password := []byte(newUser.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	newUser.Password = string(hashedPassword)

	insertResult, err := models.UserDB.Collection.InsertOne(context.TODO(), newUser)
	if err != nil {
		log.Fatal(err)
	}
	return insertResult
}

//DeleteUser function
func DeleteUser(id string) *mongo.DeleteResult {
	objectID, _ := primitive.ObjectIDFromHex(id)
	deleteResult, err := models.UserDB.Collection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
	if err != nil {
		log.Fatal(err)
	}
	return deleteResult
	// fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
}

//UpdateUser function
func UpdateUser(id string, newUpdater models.User) *mongo.UpdateResult {

	objectID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id": objectID,
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

	if newUpdater.Avatar != "" {
		bsonUpdate["avatar"] = newUpdater.Avatar
	}

	if newUpdater.UserRole != "" {
		bsonUpdate["user_role"] = newUpdater.UserRole
	}

	if newUpdater.Email != "" {
		bsonUpdate["email"] = newUpdater.Email
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

//FindOneUser function
func FindOneUser(id string) models.User {
	objectID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id": objectID,
	}
	var result models.User
	err := models.UserDB.Collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

//Login function
func Login(input models.User) models.Response {

	filter := bson.M{
		"username": input.Username,
	}
	var result models.User
	err := models.UserDB.Collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return models.Response{
			Status:  models.ResponseStatus.ERROR,
			Message: err.Error(),
		}
	}

	password := []byte(input.Password)
	err1 := bcrypt.CompareHashAndPassword([]byte(result.Password), password)
	if err != nil {
		return models.Response{
			Status:  models.ResponseStatus.ERROR,
			Message: err.Error(),
		}
	}
	if err1 != nil {
		return models.Response{
			Status:  models.ResponseStatus.ERROR,
			Message: err1.Error(),
		}
	}
	sessionToken := uuid.NewV4()
	if err != nil {
		return models.Response{
			Status:  models.ResponseStatus.ERROR,
			Message: err.Error(),
		}
	}

	models.UserCache.Set(sessionToken.String(), result.Username, cache.DefaultExpiration)

	return models.Response{
		Data:   sessionToken.String(),
		Status: models.ResponseStatus.OK,
	}
}

//Logout function
func Logout(sessionToken string) models.Response {
	item, found := models.UserCache.Get(sessionToken)
	fmt.Println(item)
	if found {
		models.UserCache.Delete(sessionToken)
		return models.Response{
			Status:  models.ResponseStatus.OK,
			Message: "Logout Successfully",
		}
	}
	return models.Response{
		Status:  models.ResponseStatus.ERROR,
		Message: "Logout Failed",
	}
}
