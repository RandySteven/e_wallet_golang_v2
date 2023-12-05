package res

import "assignment_4/entities/models"

type TransactionPaginationResponses struct {
	Page         string               `json:"page"`
	Total        uint                 `json:"total"`
	Transactions []models.Transaction `json:"transactions"`
}
