package model

type BalanceInput struct {
	Name    string  `json:"name" binding:"required"`
	Surname string  `json:"surname" binding:"required"`
	Balance float64 `json:"balance" binding:"required"`
}
