package repositories

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/entities/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/interfaces"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/utils"

	"gorm.io/gorm"
)

type walletRepository struct {
	db *gorm.DB
}

// Count implements interfaces.WalletRepository.
func (repo *walletRepository) Count(ctx context.Context) (uint, error) {
	return utils.CountTotalItems[models.Wallet](ctx, repo.db, &models.Wallet{})
}

// GetByUserId implements interfaces.WalletRepository.
func (repo *walletRepository) GetByUserId(ctx context.Context, id uint) (*models.Wallet, error) {
	var wallet *models.Wallet
	err := repo.db.WithContext(ctx).Model(&models.Wallet{}).
		Preload("User").
		Where("user_id = ?", id).
		Find(&wallet).
		Error

	if err != nil {
		return nil, err
	}

	return wallet, nil
}

// GetByNumber implements interfaces.WalletRepository.
func (repo *walletRepository) GetByNumber(ctx context.Context, number string) (*models.Wallet, error) {
	var wallet *models.Wallet
	err := repo.db.WithContext(ctx).Model(&models.Wallet{}).
		Where("number = ? ", number).
		Scan(&wallet).
		Error
	if err != nil {
		return nil, err
	}
	return wallet, nil
}

// BeginTrx implements interfaces.WalletRepository.
func (repo *walletRepository) BeginTrx(ctx context.Context) interfaces.WalletRepository {
	panic("unimplemented")
}

// CommitOrRollback implements interfaces.WalletRepository.
func (repo *walletRepository) CommitOrRollback(ctx context.Context) {
	panic("unimplemented")
}

// FindAll implements interfaces.WalletRepository.
func (repo *walletRepository) FindAll(ctx context.Context) ([]models.Wallet, error) {
	panic("unimplemented")
}

// GetById implements interfaces.WalletRepository.
func (repo *walletRepository) GetById(ctx context.Context, id uint) (*models.Wallet, error) {
	panic("unimplemented")
}

// Save implements interfaces.WalletRepository.
func (repo *walletRepository) Save(ctx context.Context, entity *models.Wallet) (*models.Wallet, error) {
	panic("unimplemented")
}

// Update implements interfaces.WalletRepository.
func (repo *walletRepository) Update(ctx context.Context, entity *models.Wallet) (*models.Wallet, error) {
	panic("unimplemented")
}

func NewWalletRepository(db *gorm.DB) *walletRepository {
	return &walletRepository{db: db}
}

var _ interfaces.WalletRepository = &walletRepository{}
