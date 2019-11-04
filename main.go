package main

import (
	"fmt"
	"net/http"

	utils "github.com/huutien2801/shop-system/utils"
	routes "github.com/huutien2801/shop-system/routes"
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
