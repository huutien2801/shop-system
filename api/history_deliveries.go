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

func FindAllHistoryAPI(w http.ResponseWriter, r *http.Request) {

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

	var input models.HistoryDelivery
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
	results := action.FindAllHistory(input, limit, offset)

	if results.Status == models.ResponseStatus.ERROR {
		respondWithError(w, http.StatusBadRequest, "No document is match with your query")
		return
	} else {
		respondWithJson(w, http.StatusOK, results)
	}
}

func CreateHistoryAPI(w http.ResponseWriter, r *http.Request) {
	var history models.HistoryDelivery
	history.ID = primitive.NewObjectID()
	*history.CreatedTime = time.Now()
	err := json.NewDecoder(r.Body).Decode(&history)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	} else {
		insertItem := action.CreateHistory(history)
		if insertItem.Status == models.ResponseStatus.ERROR {
			respondWithError(w, http.StatusBadRequest, "Create History failed")
		} else {
			respondWithJson(w, http.StatusOK, insertItem)
		}
	}
}

func DeleteHistoryAPI(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	id := keys[0]
	deleteItem := action.DeleteHistory(id)
	if deleteItem.Status == models.ResponseStatus.ERROR {
		respondWithError(w, http.StatusBadRequest, "Delete History failed")
	} else {
		respondWithJson(w, http.StatusOK, deleteItem)
	}
}

// func UpdateHistoryAPI(w http.ResponseWriter, r *http.Request) {
// 	keys, ok := r.URL.Query()["id"]

// 	if !ok || len(keys[0]) < 1 {
// 		log.Println("Url Param 'key' is missing")
// 		return
// 	}

// 	id := keys[0]
// 	var history models.HistoryDelivery
// 	err := json.NewDecoder(r.Body).Decode(&history)
// 	if err != nil {
// 		respondWithError(w, http.StatusBadRequest, err.Error())
// 		return
// 	} else {
// 		updateItem := action.UpdateHistory(id, history)
// 		respondWithJson(w, http.StatusOK, updateItem)
// 	}
// }

func FindOneHistoryAPI(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]

	if !ok || len(keys[0]) < 1 {
		respondWithError(w, http.StatusBadRequest, "Id is missing")
		return
	}

	id := keys[0]

	result := action.FindOneHistory(id)
	respondWithJson(w, http.StatusOK, result)
}
