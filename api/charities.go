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

func FindAllCharityAPI(w http.ResponseWriter, r *http.Request) {

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

	var input models.Charity
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
	results := action.FindAllCharity(input, limit, offset)

	if results == nil {
		respondWithError(w, http.StatusBadRequest, "No document is match with your query")
		return
	} else {
		respondWithJson(w, http.StatusOK, results)
	}
}

func CreateCharityAPI(w http.ResponseWriter, r *http.Request) {
	var charity models.Charity
	charity.ID = primitive.NewObjectID()
	err := json.NewDecoder(r.Body).Decode(&charity)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	} else {
		insertItem := action.CreateCharity(charity)
		if insertItem == nil {
			respondWithError(w, http.StatusBadRequest, "Create charity failed")
		} else {
			respondWithJson(w, http.StatusOK, insertItem)
		}
	}
}

func DeleteCharityAPI(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	id := keys[0]
	deleteItem := action.DeleteCharity(id)
	if deleteItem == nil {
		respondWithError(w, http.StatusBadRequest, "Delete Charity failed")
	} else {
		respondWithJson(w, http.StatusOK, deleteItem)
	}
}

func UpdateCharityAPI(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	id := keys[0]
	var charity models.Charity
	err := json.NewDecoder(r.Body).Decode(&charity)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	} else {
		updateItem := action.UpdateCharity(id, charity)
		respondWithJson(w, http.StatusOK, updateItem)
	}
}

func FindOneCharityAPI(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]

	if !ok || len(keys[0]) < 1 {
		respondWithError(w, http.StatusBadRequest, "Id is missing")
		return
	}

	id := keys[0]

	result := action.FindOneCharity(id)
	respondWithJson(w, http.StatusOK, result)
}
