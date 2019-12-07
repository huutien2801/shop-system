package api

import(
	"log"
    "github.com/huutien2801/shop-system/models"
    "github.com/huutien2801/shop-system/action"

    "net/http"
    "encoding/json"
// 	"strconv"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func FindAllProductAPI(w http.ResponseWriter, r *http.Request){
    results := action.FindAllProduct()
    if results == nil {
        respondWithError(w, http.StatusBadRequest, "Khong tim thay")
        return
    } else {
        respondWithJson(w, http.StatusOK, results)
    }
}

func CreateProductAPI(w http.ResponseWriter, r *http.Request){
    var product models.Product
	product.ID = primitive.NewObjectID()
    err := json.NewDecoder(r.Body).Decode(&product)
    if err != nil{
        respondWithError(w, http.StatusBadRequest, err.Error())
        return
    }else{
        insertItem := action.CreateProduct(product)
        if insertItem == nil {
            respondWithError(w, http.StatusBadRequest, "Create product failed")
        }else{
            respondWithJson(w, http.StatusOK, insertItem)
        }
    }
}

func DeleteProductAPI(w http.ResponseWriter, r *http.Request){
    keys, ok := r.URL.Query()["id"]

    if !ok || len(keys[0]) < 1 {
        log.Println("Url Param 'key' is missing")
        return
    }

    id := keys[0]
    deleteItem := action.DeleteProduct(id)
    if deleteItem == nil {
        respondWithError(w, http.StatusBadRequest, "Delete product failed")
    }else{
        respondWithJson(w, http.StatusOK, deleteItem)
    }
}

func UpdateProductAPI(w http.ResponseWriter, r *http.Request){
    keys, ok := r.URL.Query()["id"]

    if !ok || len(keys[0]) < 1 {
        log.Println("Url Param 'key' is missing")
        return
    }

    id := keys[0]
    var product models.Product
    err := json.NewDecoder(r.Body).Decode(&product)
    if err != nil{
        respondWithError(w, http.StatusBadRequest, err.Error())
        return
    }else{
        updateItem := action.UpdateProduct(id,product)
        respondWithJson(w, http.StatusOK, updateItem)
    }
}

func FindOneProductAPI(w http.ResponseWriter, r *http.Request){
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