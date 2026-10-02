package domain

import "context"

type Wallet struct {
	ID      uint    `gorm:"primaryKey" json:"id"`
	UserID  uint    `json:"user_id"`
	Name    string  `json:"name"` // e.g., Cash, BCA, GoPay
	Balance float64 `json:"balance"`
}

type WalletRepository interface {
	Create(ctx context.Context, wallet *Wallet) error
	GetByUserID(ctx context.Context, userID uint) ([]Wallet, error)
	UpdateBalance(ctx context.Context, id uint, amount float64) error // used if there is a transaction
}

type WalletUsecase interface {
	CreateWallet(ctx context.Context, userID uint, name string, initialBalance float64) error
	GetMyWallets(ctx context.Context, userID uint) ([]Wallet, error)
}
