package main

import (
    "context"
    "fmt"
    "os"
    "net/http"
    "github.com/huutien2801/shop-system/models"
    "github.com/huutien2801/shop-system/api"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    //routes "github.com/huutien2801/shop-system/routes"
    "github.com/gorilla/mux"
)

var username = "admin"
var host1 = "shopdb-svkhj.mongodb.net/test?retryWrites=true&w=majority" 

func onConnectedDB(client *mongo.Client){
    models.InitProductDB(client)
    fmt.Println("Connected to MongoDB successfully")
}
func main() {

    ctx := context.TODO()
	pw := "PBhUtEoEtbV8d3kJ"
    mongoURI := fmt.Sprintf("mongodb+srv://%s:%s@%s", username, pw, host1)
    fmt.Println("connection string is:", mongoURI)

    // Set client options and connect
    clientOptions := options.Client().ApplyURI(mongoURI)
    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
  
    err = client.Ping(ctx, nil)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    onConnectedDB(client)
    // api.FindAllProduct()

    r := mux.NewRouter()
    http.Handle("/", r)
    port := "5000"
    fmt.Printf("Server is running at port: %s",port)

    //API for products
    r.HandleFunc("/products", api.FindAllProductAPI).Methods("GET")
    r.HandleFunc("/products/find-one", api.FindOneProductAPI).Methods("GET")
    r.HandleFunc("/products", api.DeleteProductAPI).Methods("DELETE")
    r.HandleFunc("/products", api.CreateProductAPI).Methods("POST")
    r.HandleFunc("/products", api.UpdateProductAPI).Methods("PUT")
    //API for category
    //TODO

    //API for history-trip
    //TODO

    err2 := http.ListenAndServe(":5000", r)
	if err2 != nil {
		fmt.Println(err2)
    }
}

