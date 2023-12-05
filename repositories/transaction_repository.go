package repositories

import (
	"assignment_4/entities/models"
	"assignment_4/interfaces"
	"context"

	"gorm.io/gorm"
)

type transactionRepository struct {
	db *gorm.DB
}

// BeginTrx implements interfaces.TransactionRepository.
func (repo *transactionRepository) BeginTrx(ctx context.Context) interfaces.TransactionRepository {
	panic("unimplemented")
}

// CommitOrRollback implements interfaces.TransactionRepository.
func (repo *transactionRepository) CommitOrRollback(ctx context.Context) {
	panic("unimplemented")
}

// CreateTransaction implements interfaces.TransactionRepository.
func (repo *transactionRepository) CreateTransaction(ctx context.Context, transaction *models.Transaction) (*models.Transaction, error) {
	// err := repo.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

	// })

	return nil, nil
}

// FindAll implements interfaces.TransactionRepository.
func (repo *transactionRepository) FindAll(ctx context.Context) ([]models.Transaction, error) {
	panic("unimplemented")
}

// GetById implements interfaces.TransactionRepository.
func (repo *transactionRepository) GetById(ctx context.Context, id uint) (*models.Transaction, error) {
	panic("unimplemented")
}

// Save implements interfaces.TransactionRepository.
func (repo *transactionRepository) Save(ctx context.Context, entity *models.Transaction) (*models.Transaction, error) {
	panic("unimplemented")
}

// Update implements interfaces.TransactionRepository.
func (repo *transactionRepository) Update(ctx context.Context, entity *models.Transaction) (*models.Transaction, error) {
	panic("unimplemented")
}

func NewTransactionRepository(db *gorm.DB) *transactionRepository {
	return &transactionRepository{db: db}
}

var _ interfaces.TransactionRepository = &transactionRepository{}
