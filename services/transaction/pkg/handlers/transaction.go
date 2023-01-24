package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/kartpop/mf-folio/services/transaction/pkg/mocks"
	"github.com/kartpop/mf-folio/services/transaction/pkg/models"
)

var counter = len(mocks.Transactions) // used for txn id

func GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	txns := mocks.Transactions

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(txns)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AddTransaction(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var txn models.Transaction
	err = json.Unmarshal(body, &txn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	txn.Id = counter
	counter++                                            // monotonically increasing integer
	mocks.Transactions = append(mocks.Transactions, txn) // use mocks collection as DB for now

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(txn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
