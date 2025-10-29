package main

import (
	"log"
	"net/http"

	_ "cloudserver/docs" // Import generated Swagger docs

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title CloudServer REST API
// @version 1.0
// @description This is a sample REST API for CRUD operations on MySQL.
// @host api.devsecops2025-arubacloud.com
// @BasePath /
func main() {
	db := InitDB()
	defer db.Close()

	http.HandleFunc("/ping", PingHandler(db))
	http.HandleFunc("/items", GetItemsHandler(db))         // GET all
	http.HandleFunc("/item", GetItemHandler(db))           // GET one (by ?id=)
	http.HandleFunc("/item/create", CreateItemHandler(db)) // POST
	http.HandleFunc("/item/update", UpdateItemHandler(db)) // PUT
	http.HandleFunc("/item/delete", DeleteItemHandler(db)) // DELETE

	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
