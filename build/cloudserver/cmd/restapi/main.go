package main

import (
	"log"
	"net/http"

	"cloudserver/internal/controller"
	"cloudserver/internal/db"

	_ "cloudserver/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	database := db.InitDB()
	defer database.Close()

	log.Println("Welcome from DevSecOps-Bologna 2025 - Aruba")
	
	http.HandleFunc("/ping", controller.PingHandler(database))
	http.HandleFunc("/items", controller.GetItemsHandler(database))
	http.HandleFunc("/item", controller.GetItemHandler(database))
	http.HandleFunc("/item/create", controller.CreateItemHandler(database))
	http.HandleFunc("/item/update", controller.UpdateItemHandler(database))
	http.HandleFunc("/item/delete", controller.DeleteItemHandler(database))
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
