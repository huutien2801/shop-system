package api

import (
	"log"
	"time"

	"github.com/huutien2801/shop-system/action"
	"github.com/huutien2801/shop-system/models"

	"encoding/json"
	"net/http"

	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindAllOrderAPI(w http.ResponseWriter, r *http.Request) {

	keysQ, ok1 := r.URL.Query()["q"]
	q := ""
	if ok1 {
		q = keysQ[0]
	}

	keyLimits, ok2 := r.URL.Query()["limit"]
	limitStr := ""
	if ok2 {
		limitStr = keyLimits[0]
	}

	keyOffsets, ok3 := r.URL.Query()["offset"]
	offsetStr := ""
	if ok3 {
		offsetStr = keyOffsets[0]
	}

	var input models.Order
	// if get query
	if q != "" {
		// Unmarshal string to struct
		err := json.Unmarshal([]byte(q), &input)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid Json Query")
			return
		}
	}

	//Convert string to int64
	limit, _ := strconv.ParseInt(limitStr, 0, 64)
	offset, _ := strconv.ParseInt(offsetStr, 0, 64)
	results := action.FindAllOrder(input, limit, offset)

	if results.Status == models.ResponseStatus.ERROR {
		respondWithError(w, http.StatusBadRequest, "No document is match with your query")
		return
	} else {
		respondWithJson(w, http.StatusOK, results)
	}
}

func CreateOrderAPI(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	order.ID = primitive.NewObjectID()
	*order.CreatedTime = time.Now()
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	} else {
		insertItem := action.CreateOrder(order)
		if insertItem.Status == models.ResponseStatus.ERROR {
			respondWithError(w, http.StatusBadRequest, "Create order failed")
		} else {
			respondWithJson(w, http.StatusOK, insertItem)
		}
	}
}

func DeleteOrderAPI(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	id := keys[0]
	deleteItem := action.DeleteOrder(id)
	if deleteItem.Status == models.ResponseStatus.ERROR {
		respondWithError(w, http.StatusBadRequest, "Delete order failed")
	}else{
		respondWithJson(w, http.StatusOK, deleteItem)
	}
}

func UpdateOrderAPI(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	id := keys[0]
	var order models.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	} else {
		updateItem := action.UpdateOrder(id, order)
		respondWithJson(w, http.StatusOK, updateItem)
	}
}

func FindOneOrderAPI(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]

	if !ok || len(keys[0]) < 1 {
		respondWithError(w, http.StatusBadRequest, "Id is missing")
		return
	}

	id := keys[0]

	result := action.FindOneOrder(id)
	respondWithJson(w, http.StatusOK, result)
}
