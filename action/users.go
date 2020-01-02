package action

import (
	"context"

	"github.com/huutien2801/shop-system/models"
	"github.com/patrickmn/go-cache"
	uuid "github.com/satori/go.uuid"

	// uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

//FindAllUser function
func FindAllUser(input models.User, limit int64, offset int64) models.Response {

	//Set query
	findOptions := options.Find()
	findOptions.SetLimit(limit)
	findOptions.SetSkip(offset)
	filter := bson.M{}

	if input.FullName != "" {
		filter["full_name"] = bson.M{"$regex": input.FullName}
	}

	if input.Username != "" {
		filter["username"] = input.Username
	}
	// if input.Address != "" {
	// 	filter["address"] = bson.M{"$regex": input.Address}
	// }
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
		return models.Response{
			Status:  models.ResponseStatus.ERROR,
			Message: err.Error(),
		}
	}
	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem models.User
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
		Message: "Get data successfully",
		Data:    results,
	}
}

//CreateUser function
func CreateUser(newUser models.User) models.Response {

	password := []byte(newUser.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return models.Response{
			Status:  models.ResponseStatus.ERROR,
			Message: err.Error(),
		}
	}
	newUser.Password = string(hashedPassword)
	newUser.Avatar = "https://cdn1.vectorstock.com/i/1000x1000/19/45/user-avatar-icon-sign-symbol-vector-4001945.jpg"

	insertResult, err := models.UserDB.Collection.InsertOne(context.TODO(), newUser)
	if err != nil {
		return models.Response{
			Status:  models.ResponseStatus.ERROR,
			Message: err.Error(),
		}
	}
	return models.Response{
		Status: models.ResponseStatus.OK,
		Data:   insertResult,
	}
}

//DeleteUser function
func DeleteUser(id string) models.Response {
	objectID, _ := primitive.ObjectIDFromHex(id)
	
	deleteResult, err := models.UserDB.Collection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
	if err != nil {
		return models.Response{
			Status:  models.ResponseStatus.ERROR,
			Message: err.Error(),
		}
	}
	return models.Response{
		Status:  models.ResponseStatus.OK,
		Message: "Delete successfully",
		Data:    deleteResult,
	}
	// fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
}

//UpdateUser function
func UpdateUser(id string, newUpdater models.User) models.Response {

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

	if newUpdater.Password != "" {
		password := []byte(newUpdater.Password)
		hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
		if err != nil {
			return models.Response{
				Status:  models.ResponseStatus.ERROR,
				Message: err.Error(),
			}
		}
		bsonUpdate["password"] = string(hashedPassword)
	}
	// if newUpdater.Address != "" {
	// 	bsonUpdate["address"] = newUpdater.Address
	// }

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
		return models.Response{
			Status:  models.ResponseStatus.ERROR,
			Message: err.Error(),
		}
	}
	return models.Response{
		Status:  models.ResponseStatus.OK,
		Message: "Update successfully",
		Data:    updateResult,
	}
}

//FindOneUser function
func FindOneUser(id string) models.Response {
	objectID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id": objectID,
	}
	var result models.User
	err := models.UserDB.Collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return models.Response{
			Status:  models.ResponseStatus.ERROR,
			Message: err.Error(),
		}
	}
	return models.Response{
		Status:  models.ResponseStatus.OK,
		Message: "Update successfully",
		Data:    result,
	}
}

//Login function
func Login(input models.User, availableSessionToken string) models.Response {
	if availableSessionToken != "" {
		item, found := models.UserCache.Get(availableSessionToken)
		if found {
			filter := bson.M{
				"username": item,
			}
			var result models.User
			err := models.UserDB.Collection.FindOne(context.TODO(), filter).Decode(&result)
			if err != nil {
				return models.Response{
					Status:  models.ResponseStatus.ERROR,
					Message: err.Error(),
				}
			}
			result.Session = availableSessionToken
			return models.Response{
				Data:   result,
				Status: models.ResponseStatus.OK,
			}
		}
	}
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
	result.Session = sessionToken.String()

	return models.Response{
		Data:   result,
		Status: models.ResponseStatus.OK,
	}
}

//Logout function
func Logout(sessionToken string) models.Response {
	_, found := models.UserCache.Get(sessionToken)

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
