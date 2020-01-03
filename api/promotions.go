package api

import (
	"log"

	"github.com/huutien2801/shop-system/action"
	"github.com/huutien2801/shop-system/models"

	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strconv"
	"time"
)

func FindAllPromotionAPI(w http.ResponseWriter, r *http.Request) {

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

	var input models.Promotion
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
	results := action.FindAllPromotion(input, limit, offset)

	if results.Status == models.ResponseStatus.ERROR {
		respondWithError(w, http.StatusBadRequest, "No document is match with your query")
		return
	} else {
		respondWithJson(w, http.StatusOK, results)
	}
}

func CreatePromotionAPI(w http.ResponseWriter, r *http.Request) {
	var promotion models.Promotion
	promotion.ID = primitive.NewObjectID()
	now := time.Now()
	promotion.CreatedTime = &now
	err := json.NewDecoder(r.Body).Decode(&promotion)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	} else {
		insertItem := action.CreatePromotion(promotion)
		if insertItem.Status == models.ResponseStatus.ERROR {
			respondWithError(w, http.StatusBadRequest, "Create Promotion failed")
		} else {
			respondWithJson(w, http.StatusOK, insertItem)
		}
	}
}

func DeletePromotionAPI(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	id := keys[0]
	deleteItem := action.DeletePromotion(id)
	if deleteItem.Status == models.ResponseStatus.ERROR {
		respondWithError(w, http.StatusBadRequest, "Delete Promotion failed")
	} else {
		respondWithJson(w, http.StatusOK, deleteItem)
	}
}

func UpdatePromotionAPI(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	id := keys[0]
	var promotion models.Promotion
	err := json.NewDecoder(r.Body).Decode(&promotion)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	} else {
		updateItem := action.UpdatePromotion(id, promotion)
		respondWithJson(w, http.StatusOK, updateItem)
	}
}

func FindOnePromotionAPI(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]

	if !ok || len(keys[0]) < 1 {
		respondWithError(w, http.StatusBadRequest, "Id is missing")
		return
	}

	id := keys[0]

	result := action.FindOnePromotion(id)
	respondWithJson(w, http.StatusOK, result)
}
