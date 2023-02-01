package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/kartpop/mf-folio/services/transaction/pkg/handlers"
)

func main() {
	// load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err.Error())
	}

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	h, err := handlers.New(dbURL)
	if err != nil {
		log.Fatalln(err.Error())
	}

	router := mux.NewRouter()
	router.HandleFunc("/transactions", h.GetAllTransactions).Methods(http.MethodGet)
	router.HandleFunc("/transactions", h.AddTransaction).Methods(http.MethodPost)

	log.Println("Transaction server is running!")
	http.ListenAndServe(":6000", router)
}
