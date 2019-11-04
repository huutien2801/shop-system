package main

import (
	"fmt"
	"net/http"

	"CRUD/routes"
	"CRUD/utils"
)

func main() {
	utils.LoadTemplates("templates/*.html")
	r := routes.NewRouter()
	http.Handle("/", r)
	err := http.ListenAndServe(":5000", r)

	if err != nil {
		fmt.Println(err)
	}
}
