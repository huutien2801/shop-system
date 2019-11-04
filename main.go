package main

import (
	"fmt"
	"net/http"

	routes "github.com/huutien2801/shop-system/routes"
)

func main() {
	// utils.LoadTemplates("templates/*.html")
	r := routes.NewRouter()
	http.Handle("/", r)
	err := http.ListenAndServe(":5000", r)

	if err != nil {
		fmt.Println(err)
	}
}
