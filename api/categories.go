package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/huutien2801/shop-system/action"
	"github.com/huutien2801/shop-system/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindAllCategoriesAPI(w http.ResponseWriter, r *http.Request) {
	results := action.FindAllCategories()
	if results == nil {
		respondWithError(w, http.StatusBadRequest, "Khong tim thay")
		return
	} else {
		respondWithJson(w, http.StatusOK, results)
	}
}

func CreateCategoryAPI(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	category.ID = primitive.NewObjectID()
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	} else {
		insertItem := action.CreateCategory(category)
		if insertItem == nil {
			respondWithError(w, http.StatusBadRequest, "Create product failed")
		} else {
			respondWithJson(w, http.StatusOK, insertItem)
		}
	}
}

func DeleteCategoryAPI(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	id := keys[0]
	deleteItem := action.DeleteCategory(id)
	if deleteItem == nil {
		respondWithError(w, http.StatusBadRequest, "Delete product failed")
	} else {
		respondWithJson(w, http.StatusOK, deleteItem)
	}
}

func UpdateCategoryAPI(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	id := keys[0]
	var category models.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	} else {
		updateItem := action.UpdateCategory(id, category)
		respondWithJson(w, http.StatusOK, updateItem)
	}
}

func FindOneCategoryAPI(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]

	if !ok || len(keys[0]) < 1 {
		respondWithError(w, http.StatusBadRequest, "Id is missing")
		return
	}

	id := keys[0]

	result := action.FindOneCategory(id)
	respondWithJson(w, http.StatusOK, result)
}
