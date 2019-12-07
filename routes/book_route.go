package routes

// import (
// 	 "github.com/huutien2801/shop-system/config"
// 	 "github.com/huutien2801/shop-system/entities"
// 	 "github.com/huutien2801/shop-system/models"
// 	 "github.com/huutien2801/shop-system/utils"
// 	"encoding/json"
// 	"net/http"
// 	"strconv"

// 	"github.com/gorilla/mux"
// 	"gopkg.in/mgo.v2/bson"
// )

// func NewRouter() *mux.Router {

// 	r := mux.NewRouter()

// 	//API CRUD
// 	r.HandleFunc("/api/book/findall", FindAllAPI).Methods("GET")
// 	r.HandleFunc("/api/book/find/{id}", FindByIdAPI).Methods("GET")
// 	r.HandleFunc("/api/book/create", CreateAPI).Methods("POST")
// 	r.HandleFunc("/api/book/update", UpdateAPI).Methods("PUT")
// 	r.HandleFunc("/api/book/delete/{id}", DeleteAPI).Methods("DELETE")

// 	//CRUD
// 	r.HandleFunc("/book/view", ViewAllBook).Methods("GET")
// 	r.HandleFunc("/book/add", AddBook).Methods("GET")
// 	r.HandleFunc("/book/add", AddBookPOST).Methods("POST")
// 	r.HandleFunc("/book/update/{id}", UpdateBook).Methods("GET")
// 	r.HandleFunc("/book/update/{id}", UpdateBookPOST).Methods("POST")
// 	r.HandleFunc("/book/delete/{id}", DeleteBook)

// 	//Create link for static file
// 	fs := http.FileServer(http.Dir("./statics/"))
// 	r.PathPrefix("/statics/").Handler(http.StripPrefix("/statics/", fs))
// 	return r
// }

// //Render view add book
// func AddBook(w http.ResponseWriter, r *http.Request) {
// 	utils.ExecuteTemplate(w, "add.html", nil)
// }

// func AddBookPOST(w http.ResponseWriter, r *http.Request) {
// 	db, err := config.GetMongoDB()

// 	if err != nil {
// 		respondWithError(w, http.StatusBadRequest, err.Error())
// 	} else {
// 		bookModel := models.BookModel{
// 			Db:         db,
// 			Collection: "books",
// 		}

// 		r.ParseForm()
// 		var book entities.Book
// 		book.Id = bson.NewObjectId()
// 		book.Name = r.PostForm.Get("name")
// 		book.Category = r.PostForm.Get("category")
// 		book.Price, _ = strconv.ParseFloat(r.PostForm.Get("price"), 64)

// 		err3 := bookModel.Create(&book)
// 		if err3 != nil {
// 			respondWithError(w, http.StatusBadRequest, err3.Error())
// 			return
// 		} else {
// 			http.Redirect(w, r, "/book/view", http.StatusFound)
// 		}
// 	}
// }

// func DeleteBook(w http.ResponseWriter, r *http.Request) {
// 	db, err := config.GetMongoDB()

// 	if err != nil {
// 		respondWithError(w, http.StatusBadRequest, err.Error())
// 	} else {
// 		bookModel := models.BookModel{
// 			Db:         db,
// 			Collection: "books",
// 		}

// 		vars := mux.Vars(r)
// 		id := vars["id"]
// 		book, _ := bookModel.FindById(id)
// 		err2 := bookModel.Delete(book)

// 		if err2 != nil {
// 			respondWithError(w, http.StatusBadRequest, err.Error())
// 			return
// 		} else {
// 			http.Redirect(w, r, "/book/view", http.StatusFound)
// 		}
// 	}
// }
// func UpdateBookPOST(w http.ResponseWriter, r *http.Request) {
// 	db, err := config.GetMongoDB()

// 	if err != nil {
// 		respondWithError(w, http.StatusBadRequest, err.Error())
// 	} else {
// 		bookModel := models.BookModel{
// 			Db:         db,
// 			Collection: "books",
// 		}

// 		vars := mux.Vars(r)
// 		id := vars["id"]
// 		book, _ := bookModel.FindById(id)

// 		r.ParseForm()

// 		book.Name = r.PostForm.Get("name")
// 		book.Category = r.PostForm.Get("category")
// 		book.Price, _ = strconv.ParseFloat(r.PostForm.Get("price"), 64)

// 		err3 := bookModel.Update(&book)
// 		if err3 != nil {
// 			respondWithError(w, http.StatusBadRequest, err3.Error())
// 			return
// 		} else {
// 			http.Redirect(w, r, "/book/view", http.StatusFound)
// 		}
// 	}
// }

// func UpdateBook(w http.ResponseWriter, r *http.Request) {

// 	db, err := config.GetMongoDB()

// 	if err != nil {
// 		respondWithError(w, http.StatusBadRequest, err.Error())
// 	} else {
// 		bookModel := models.BookModel{
// 			Db:         db,
// 			Collection: "books",
// 		}

// 		vars := mux.Vars(r)
// 		id := vars["id"]
// 		book, err2 := bookModel.FindById(id)

// 		if err2 != nil {
// 			respondWithError(w, http.StatusBadRequest, err.Error())
// 			return
// 		} else {
// 			utils.ExecuteTemplate(w, "update.html", book)
// 		}
// 	}
// }

// func ViewAllBook(w http.ResponseWriter, r *http.Request) {
// 	db, err := config.GetMongoDB()
// 	if err != nil {
// 		respondWithError(w, http.StatusBadRequest, err.Error())
// 	} else {
// 		bookModel := models.BookModel{
// 			Db:         db,
// 			Collection: "books",
// 		}
// 		books, err2 := bookModel.FindAll()

// 		if err2 != nil {
// 			respondWithError(w, http.StatusBadRequest, err.Error())
// 			return
// 		} else {
// 			utils.ExecuteTemplate(w, "view.html", books)
// 		}
// 	}
// }

// //API CRUD

// //API find all
// func FindAllAPI(w http.ResponseWriter, r *http.Request) {
// 	db, err := config.GetMongoDB()

// 	if err != nil {
// 		respondWithError(w, http.StatusBadRequest, err.Error())
// 	} else {
// 		bookModel := models.BookModel{
// 			Db:         db,
// 			Collection: "books",
// 		}

// 		books, err2 := bookModel.FindAll()

// 		if err2 != nil {
// 			respondWithError(w, http.StatusBadRequest, err.Error())
// 			return
// 		} else {
// 			respondWithJson(w, http.StatusOK, books)
// 		}
// 	}
// }

// //API find book by id
// func FindByIdAPI(w http.ResponseWriter, r *http.Request) {
// 	db, err := config.GetMongoDB()

// 	if err != nil {
// 		respondWithError(w, http.StatusBadRequest, err.Error())
// 	} else {
// 		bookModel := models.BookModel{
// 			Db:         db,
// 			Collection: "books",
// 		}

// 		vars := mux.Vars(r)
// 		id := vars["id"]
// 		book, err2 := bookModel.FindById(id)

// 		if err2 != nil {
// 			respondWithError(w, http.StatusBadRequest, err.Error())
// 			return
// 		} else {
// 			respondWithJson(w, http.StatusOK, book)
// 		}
// 	}
// }

// //API delete book
// func DeleteAPI(w http.ResponseWriter, r *http.Request) {
// 	db, err := config.GetMongoDB()

// 	if err != nil {
// 		respondWithError(w, http.StatusBadRequest, err.Error())
// 	} else {
// 		bookModel := models.BookModel{
// 			Db:         db,
// 			Collection: "books",
// 		}

// 		vars := mux.Vars(r)
// 		id := vars["id"]
// 		book, _ := bookModel.FindById(id)
// 		err2 := bookModel.Delete(book)

// 		if err2 != nil {
// 			respondWithError(w, http.StatusBadRequest, err.Error())
// 			return
// 		} else {
// 			respondWithJson(w, http.StatusOK, entities.Book{})
// 		}
// 	}
// }

// //API create book
// func CreateAPI(w http.ResponseWriter, r *http.Request) {
// 	db, err := config.GetMongoDB()

// 	if err != nil {
// 		respondWithError(w, http.StatusBadRequest, err.Error())
// 	} else {
// 		bookModel := models.BookModel{
// 			Db:         db,
// 			Collection: "books",
// 		}

// 		var book entities.Book
// 		book.Id = bson.NewObjectId()
// 		err2 := json.NewDecoder(r.Body).Decode(&book)
// 		if err2 != nil {
// 			respondWithError(w, http.StatusBadRequest, err2.Error())
// 			return
// 		} else {
// 			err3 := bookModel.Create(&book)
// 			if err3 != nil {
// 				respondWithError(w, http.StatusBadRequest, err3.Error())
// 				return
// 			} else {
// 				respondWithJson(w, http.StatusOK, book)
// 			}
// 		}
// 	}
// }

// //API update book
// func UpdateAPI(w http.ResponseWriter, r *http.Request) {
// 	db, err := config.GetMongoDB()

// 	if err != nil {
// 		respondWithError(w, http.StatusBadRequest, err.Error())
// 	} else {

// 		bookModel := models.BookModel{
// 			Db:         db,
// 			Collection: "books",
// 		}

// 		var book entities.Book
// 		err2 := json.NewDecoder(r.Body).Decode(&book)

// 		if err2 != nil {
// 			respondWithError(w, http.StatusBadRequest, err2.Error())
// 			return
// 		} else {
// 			err3 := bookModel.Update(&book)
// 			if err3 != nil {
// 				respondWithError(w, http.StatusBadRequest, err3.Error())
// 				return
// 			} else {
// 				respondWithJson(w, http.StatusOK, book)
// 			}
// 		}
// 	}
// }

// //When having err, respond error with Json
// func respondWithError(w http.ResponseWriter, code int, msg string) {
// 	respondWithJson(w, code, map[string]string{"error": msg})
// }

// func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
// 	response, _ := json.Marshal(payload)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(code)
// 	w.Write(response)
// }
