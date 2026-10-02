package domain

import (
	"context"
	"time"
)

type Transaction struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	UserID      uint      `json:"user_id"`
	WalletID    uint      `json:"wallet_id"`
	Type        string    `json:"type"`     // "INCOME" or "EXPENSE"
	Category    string    `json:"category"` // e.g., "Food", "Salary"
	Amount      float64   `json:"amount"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

// Struct specifically for the Profit and Loss Statement
type ProfitLoss struct {
	TotalIncome  float64 `json:"total_income"`
	TotalExpense float64 `json:"total_expense"`
	NetProfit    float64 `json:"net_profit"`
}

type TransactionRepository interface {
	Create(ctx context.Context, tx *Transaction) error
	GetReport(ctx context.Context, userID uint, startDate, endDate string) ([]Transaction, error)
}

type TransactionUsecase interface {
	RecordTransaction(ctx context.Context, tx *Transaction) error
	GetReport(ctx context.Context, userID uint, startDate, endDate string) ([]Transaction, ProfitLoss, error)
}
