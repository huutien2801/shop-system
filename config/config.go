package config

import (
	//"gopkg.in/mgo.v2/bson"

	//"go.mongodb.org/mongo-driver/bson"
    // "context"
    // "fmt"
    // "os"
  
    // "go.mongodb.org/mongo-driver/mongo"
    // "go.mongodb.org/mongo-driver/mongo/options"
   
)
var username = "admin"
var host1 = "shopdb-svkhj.mongodb.net/test?retryWrites=true&w=majority"  // of the form foo.mongodb.net
// func GetMongoDB() (mongo.Client, error) {

// 	ctx := context.TODO()

//     // pw, ok := os.LookupEnv("MONGO_PW")
//     // if !ok {
//     //     fmt.Println("error: unable to find MONGO_PW in the environment")
//     //     os.Exit(1)
// 	// }
// 	pw := "PBhUtEoEtbV8d3kJ"
//     mongoURI := fmt.Sprintf("mongodb+srv://%s:%s@%s", username, pw, host1)
//     fmt.Println("connection string is:", mongoURI)

//     // Set client options and connect
//     clientOptions := options.Client().ApplyURI(mongoURI)
//     client, err := mongo.Connect(ctx, clientOptions)
//     if err != nil {
//         fmt.Println(err)
//         os.Exit(1)
//     }
  
//     err = client.Ping(ctx, nil)
//     if err != nil {
//         fmt.Println(err)
//         os.Exit(1)
// 	}
// }
