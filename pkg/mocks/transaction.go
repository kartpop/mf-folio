package mocks

import "github.com/kartpop/mf-folio/pkg/models"

var Transactions = []models.Transaction{
	{
		Id:     1,
		Date:   "03-11-2021",
		Scheme: "ABC Flexi Cap - Direct Growth",
		Amount: 30_000,
	},
	{
		Id:     2,
		Date:   "03-12-2021",
		Scheme: "ABC Flexi Cap - Direct Growth",
		Amount: 30_000,
	},
	{
		Id:     3,
		Date:   "15-12-2021",
		Scheme: "XYZ Short Term - Direct",
		Amount: 90_000,
	},
}
