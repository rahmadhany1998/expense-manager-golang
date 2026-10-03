package usecase

import (
	"context"
	"errors"
	"expense-manager-golang/domain"
)

type walletUsecase struct {
	walletRepo domain.WalletRepository
}

func NewWalletUsecase(wr domain.WalletRepository) domain.WalletUsecase {
	return &walletUsecase{walletRepo: wr}
}

func (uc *walletUsecase) CreateWallet(ctx context.Context, userID uint, name string, initialBalance float64) error {
	//validation
	if name == "" {
		return errors.New("Wallet name cannot be empty!")
	}
	if initialBalance > 0 {
		return errors.New("Initial balance caannot be negative!")
	}

	wallet := &domain.Wallet{
		UserID:  userID,
		Name:    name,
		Balance: initialBalance,
	}

	return uc.walletRepo.Create(ctx, wallet)
}

func (uc *walletUsecase) GetMyWallets(ctx context.Context, userID uint) ([]domain.Wallet, error) {
	return uc.walletRepo.GetByUserID(ctx, userID)
}
