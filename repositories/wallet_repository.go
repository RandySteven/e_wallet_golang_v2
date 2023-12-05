package repositories

import (
	"assignment_4/entities/models"
	"assignment_4/interfaces"
	"context"

	"gorm.io/gorm"
)

type walletRepository struct {
	db *gorm.DB
}

// BeginTrx implements interfaces.WalletRepository.
func (*walletRepository) BeginTrx(ctx context.Context) interfaces.WalletRepository {
	panic("unimplemented")
}

// CommitOrRollback implements interfaces.WalletRepository.
func (*walletRepository) CommitOrRollback(ctx context.Context) {
	panic("unimplemented")
}

// FindAll implements interfaces.WalletRepository.
func (*walletRepository) FindAll(ctx context.Context) ([]models.Wallet, error) {
	panic("unimplemented")
}

// GetById implements interfaces.WalletRepository.
func (*walletRepository) GetById(ctx context.Context, id uint) (*models.Wallet, error) {
	panic("unimplemented")
}

// Save implements interfaces.WalletRepository.
func (*walletRepository) Save(ctx context.Context, entity *models.Wallet) (*models.Wallet, error) {
	panic("unimplemented")
}

// Update implements interfaces.WalletRepository.
func (*walletRepository) Update(ctx context.Context, entity *models.Wallet) (*models.Wallet, error) {
	panic("unimplemented")
}

func NewWalletRepository(db *gorm.DB) *walletRepository {
	return &walletRepository{db: db}
}

var _ interfaces.WalletRepository = &walletRepository{}
