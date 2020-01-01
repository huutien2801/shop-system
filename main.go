package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/huutien2801/shop-system/api"
	"github.com/huutien2801/shop-system/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	//routes "github.com/huutien2801/shop-system/routes"
	"github.com/gorilla/mux"
)

var username = "admin"
var host1 = "shopdb-svkhj.mongodb.net/test?retryWrites=true&w=majority"

func onConnectedDB(client *mongo.Client) {
	models.InitProductDB(client)
	models.InitCategoryDB(client)
	models.InitCharityDB(client)
	models.InitPromotionDB(client)
	models.InitOrderDB(client)
	models.InitUserDB(client)
	models.InitHistoryDeliveryDB(client)
	fmt.Println("Connected to MongoDB successfully")
}

func onInitCache() {
	models.InitUserCache()
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

	//Init Cache
	onInitCache()

	r := mux.NewRouter()
	http.Handle("/", r)

	//API for products
	r.HandleFunc("/products", api.FindAllProductAPI).Methods("GET")
	r.HandleFunc("/products/find-one", api.FindOneProductAPI).Methods("GET")
	r.HandleFunc("/products", api.DeleteProductAPI).Methods("DELETE")
	r.HandleFunc("/products", api.CreateProductAPI).Methods("POST")
	r.HandleFunc("/products", api.UpdateProductAPI).Methods("PUT")
	//API for category
	r.HandleFunc("/categories", api.FindAllCategoriesAPI).Methods("GET")
	r.HandleFunc("/categories/find-one", api.FindOneCategoryAPI).Methods("GET")
	r.HandleFunc("/categories", api.DeleteCategoryAPI).Methods("DELETE")
	r.HandleFunc("/categories", api.CreateCategoryAPI).Methods("POST")
	r.HandleFunc("/categories", api.UpdateCategoryAPI).Methods("PUT")
	//API for user
	r.HandleFunc("/user", api.FindAllUserAPI).Methods("GET")
	r.HandleFunc("/user/find-one", api.FindOneUserAPI).Methods("GET")
	r.HandleFunc("/user", api.DeleteUserAPI).Methods("DELETE")
	r.HandleFunc("/user", api.CreateUserAPI).Methods("POST")
	r.HandleFunc("/user", api.UpdateUserAPI).Methods("PUT")
	r.HandleFunc("/user/login", api.LoginAPI).Methods("POST")
	r.HandleFunc("/user/logout", api.LogoutAPI).Methods("POST")
	//API for history
	r.HandleFunc("/history", api.FindAllHistoryAPI).Methods("GET")
	r.HandleFunc("/history/find-one", api.FindOneHistoryAPI).Methods("GET")
	r.HandleFunc("/history", api.DeleteHistoryAPI).Methods("DELETE")
	r.HandleFunc("/history", api.CreateHistoryAPI).Methods("POST")
	//API for charity
	r.HandleFunc("/charity", api.FindAllCharityAPI).Methods("GET")
	r.HandleFunc("/charity/find-one", api.FindOneCharityAPI).Methods("GET")
	r.HandleFunc("/charity", api.DeleteCharityAPI).Methods("DELETE")
	r.HandleFunc("/charity", api.CreateCharityAPI).Methods("POST")
	r.HandleFunc("/charity", api.UpdateCharityAPI).Methods("PUT")
	//API for promotion
	r.HandleFunc("/promotion", api.FindAllPromotionAPI).Methods("GET")
	r.HandleFunc("/promotion/find-one", api.FindOnePromotionAPI).Methods("GET")
	r.HandleFunc("/promotion", api.DeletePromotionAPI).Methods("DELETE")
	r.HandleFunc("/promotion", api.CreatePromotionAPI).Methods("POST")
	r.HandleFunc("/promotion", api.UpdatePromotionAPI).Methods("PUT")
	//API for history-trip
	//TODO
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	fmt.Printf("Server is running at port: %s", port)
	err2 := http.ListenAndServe(":"+port, r)
	if err2 != nil {
		fmt.Println(err2)
	}
}
