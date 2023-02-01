package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/kartpop/mf-folio/services/transaction/pkg/helpers"
	"github.com/kartpop/mf-folio/services/transaction/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	*gorm.DB
	minTxnNum, maxTxnNum int
}

func New(dbURL string) (*Handler, error) {
	gormdb, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = gormdb.AutoMigrate(&models.Transaction{})
	if err != nil {
		return nil, err
	}

	minTxnNum, err := strconv.Atoi(os.Getenv("MIN_TXN_NUM"))
	if err != nil {
		return nil, err
	}

	maxTxnNum, err := strconv.Atoi(os.Getenv("MAX_TXN_NUM"))
	if err != nil {
		return nil, err
	}

	return &Handler{
		DB:        gormdb,
		minTxnNum: minTxnNum,
		maxTxnNum: maxTxnNum,
	}, nil
}

func (h *Handler) GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	var txns []*models.Transaction
	result := h.DB.Find(&txns)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(txns)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) AddTransaction(w http.ResponseWriter, r *http.Request) {
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

	err = h.writeTxnToDB(&txn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(txn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h Handler) writeTxnToDB(txn *models.Transaction) error {
	for { // run till txn Id is unique
		txn.Id = helpers.RandInt(h.minTxnNum, h.maxTxnNum)
		result := h.DB.Create(txn)
		if result.Error != nil {
			if !strings.Contains(result.Error.Error(), "SQLSTATE 23505") { // only return if it is not duplicate key error
				return result.Error
			}
		} else {
			return nil
		}
	}
}
