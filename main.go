package main

import (
	"goweb/config"
	"goweb/controllers/categorycontroller"
	"goweb/controllers/homecontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	// homepage
	http.HandleFunc("/", homecontroller.Welcome)

	// categories
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	log.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
