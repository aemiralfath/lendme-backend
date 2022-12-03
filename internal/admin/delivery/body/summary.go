package body

import "final-project-backend/internal/models"

type SummaryResponse struct {
	UserTotal     int64             `json:"user_total"`
	LendingAmount float64           `json:"lending_amount"`
	ReturnAmount  float64           `json:"return_amount"`
	LendingTotal  int64             `json:"lending_total"`
	LendingAction []*models.Lending `json:"lending_action"`
	UserAction    []*models.Debtor  `json:"user_action"`
}
