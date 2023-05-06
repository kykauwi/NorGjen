package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kykauwi/NorGjen/handlers"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/Invoice", handlers.SendInvoice)
	router.HandleFunc("/Notify", handlers.SendNotification)
	log.Fatal(http.ListenAndServe(":8080", router))
}
