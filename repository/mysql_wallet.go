package repository

import (
	"context"
	"expense-manager-golang/domain"

	"gorm.io/gorm"
)

type gormWalletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) domain.WalletRepository {
	return &gormWalletRepository{db: db}
}

func (r *gormWalletRepository) getDB(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value("txKey").(*gorm.DB); ok {
		return tx
	}
	return r.db.WithContext(ctx)
}

func (r *gormWalletRepository) Create(ctx context.Context, wallet *domain.Wallet) error {
	return r.db.WithContext(ctx).Create(wallet).Error
}

func (r *gormWalletRepository) GetByUserID(ctx context.Context, userID uint) ([]domain.Wallet, error) {
	var wallets []domain.Wallet
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&wallets).Error
	return wallets, err
}

func (r *gormWalletRepository) UpdateBalance(ctx context.Context, id uint, amount float64) error {
	db := r.getDB(ctx)
	return db.Model(&domain.Wallet{}).Where("id = ?", id).Update("balance", gorm.Expr("balance + ?", amount)).Error
}
