package api

import (
	"log"

	"github.com/huutien2801/shop-system/action"
	"github.com/huutien2801/shop-system/models"

	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
)

func FindAllProductAPI(w http.ResponseWriter, r *http.Request) {

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

	var input models.ClientProductInput
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
	results := action.FindAllProduct(input, limit, offset)

	enableCors(&w)
	if results == nil {
		respondWithError(w, http.StatusBadRequest, "No document is match with your query")
		return
	} else {
		respondWithJson(w, http.StatusOK, results)
	}
}

func CreateProductAPI(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	product.ID = primitive.NewObjectID()
	*product.CreatedTime = time.Now()
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	} else {
		insertItem := action.CreateProduct(product)
		if insertItem == nil {
			respondWithError(w, http.StatusBadRequest, "Create product failed")
		} else {
			respondWithJson(w, http.StatusOK, insertItem)
		}
	}
}

func DeleteProductAPI(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	id := keys[0]
	deleteItem := action.DeleteProduct(id)
	if deleteItem == nil {
		respondWithError(w, http.StatusBadRequest, "Delete product failed")
	} else {
		respondWithJson(w, http.StatusOK, deleteItem)
	}
}

func UpdateProductAPI(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	id := keys[0]
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	} else {
		updateItem := action.UpdateProduct(id, product)
		respondWithJson(w, http.StatusOK, updateItem)
	}
}

func FindOneProductAPI(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]

	if !ok || len(keys[0]) < 1 {
		respondWithError(w, http.StatusBadRequest, "Id is missing")
		return
	}

	id := keys[0]

	result := action.FindOneProduct(id)
	respondWithJson(w, http.StatusOK, result)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
