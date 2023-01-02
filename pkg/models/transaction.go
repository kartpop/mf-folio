package models

type Transaction struct {
	Id     int    `json:"id"`
	Date   string `json:"date"` // "dd-mm-yyyy"
	Scheme string `json:"scheme"`
	Amount int    `json:"amount"`
}
