package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kartpop/mf-folio/pkg/handlers"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/transactions", handlers.GetAllTransactions).Methods(http.MethodGet)
	router.HandleFunc("/transactions", handlers.AddTransaction).Methods(http.MethodPost)

	log.Println("Transaction server is running!")
	http.ListenAndServe(":4000", router)
}
