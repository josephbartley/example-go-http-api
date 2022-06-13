package main

import (
	"fmt"
	"net/http"

	"josephbartley.dev/go-api-test/handlers"
)

func main() {
	http.HandleFunc("/ping", handlers.PingHandler)
	http.HandleFunc("/items", handlers.GetAllItemsHandler)
	http.HandleFunc("/item", handlers.AddItemHandler)
	http.HandleFunc("/item/", handlers.ItemRouter)

	fmt.Println("Listening on :8090...")
	http.ListenAndServe(":8090", nil)
}
