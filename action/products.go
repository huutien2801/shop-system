package action

import(
	"log"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/huutien2801/shop-system/models"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"fmt"
)

func FindAllProduct() []*models.Product{

	findOptions := options.Find()
    findOptions.SetLimit(2)

	var results []*models.Product
    cur, err := models.ProductDB.Collection.Find(context.TODO(), bson.M{}, findOptions)
    if err != nil {
        log.Fatal(err)
    }
    // Finding multiple documents returns a cursor
    // Iterating through the cursor allows us to decode documents one at a time
    for cur.Next(context.TODO()) {
    
    // create a value into which the single document can be decoded
    var elem models.Product
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
    for _,value :=range results{
        fmt.Println(*value)
	}
	return results
}

func CreateProduct(newProduct models.Product) *mongo.InsertOneResult{
    
    insertResult, err := models.ProductDB.Collection.InsertOne(context.TODO(), newProduct)
    if err != nil {
        log.Fatal(err)
    }
  	return insertResult
}

func DeleteProduct(id string) *mongo.DeleteResult{
	objectId,_ := primitive.ObjectIDFromHex(id)
    deleteResult, err := models.ProductDB.Collection.DeleteOne(context.TODO(), bson.M{"_id": objectId})
    if err != nil {
        log.Fatal(err)
	}
	return deleteResult
    // fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
}

func UpdateProduct(id string, newUpdater models.Product) *mongo.UpdateResult{

	objectId,_ := primitive.ObjectIDFromHex(id)
    filter := bson.M{
        "_id": objectId,
    }
	
	bsonUpdate := bson.M{}
	if newUpdater.Name != ""{
		bsonUpdate["name"] = newUpdater.Name
	}

	if newUpdater.Price != 0{
		bsonUpdate["price"] = newUpdater.Price
	}
    update := bson.M{
        "$set": bsonUpdate,
    }
    // var res Product
    updateResult, err  := models.ProductDB.Collection.UpdateOne(context.TODO(), filter, update)
    if err != nil {
        log.Fatal(err)
    }   
   return updateResult
}

func FindOneProduct(id string) models.Product{
	objectId,_ := primitive.ObjectIDFromHex(id)
    filter := bson.M{
        "_id": objectId,
    }
    var result models.Product
    err := models.ProductDB.Collection.FindOne(context.TODO(), filter).Decode(&result)
    if err != nil {
        log.Fatal(err)
    }
	return result
}