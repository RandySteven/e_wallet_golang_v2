package repositories

import (
	"context"
	"strconv"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/entities"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/entities/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/enums"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/interfaces"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/utils"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type transactionRepository struct {
	db *gorm.DB
}

// GetTransactionCountBasedUserId implements interfaces.TransactionRepository.
func (repo *transactionRepository) GetTransactionCountBasedUserId(ctx context.Context, id uint) (uint, error) {
	var res int64 = 0
	err := repo.db.WithContext(ctx).Model(&models.Transaction{}).Where("receiver_id = ? OR sender_id = ?", id, id).Count(&res).Error
	out := uint(res)
	if err != nil {
		return out, err
	}
	return out, nil

}

// Count implements interfaces.TransactionRepository.
func (repo *transactionRepository) Count(ctx context.Context) (uint, error) {
	return utils.CountTotalItems[models.Transaction](ctx, repo.db, &models.Transaction{})
}

// GetAllTransactions implements interfaces.TransactionRepository.
func (repo *transactionRepository) GetAllTransactions(ctx context.Context, query *entities.QueryCondition, walletId uint) ([]models.Transaction, error) {
	var transactions []models.Transaction
	limit, _ := strconv.Atoi(query.Limit)
	page, _ := strconv.Atoi(query.Page)
	desc := false
	if query.Sort == enums.Desc {
		desc = true
	}
	sql := repo.db.WithContext(ctx).Model(&models.Transaction{}).
		Preload("Receiver.User").
		Preload("Sender.User").
		Where("sender_id = ? OR receiver_id = ?", walletId, walletId).
		Offset((page - 1) * limit)

	if query.SortedBy != "" {
		sql.Order(clause.OrderByColumn{
			Column: clause.Column{Name: query.SortedBy},
			Desc:   desc,
		})
	}

	if query.StartDate != "" && query.EndDate != "" {
		sql.Where("created_at BETWEEN ? AND ?", query.StartDate, query.EndDate)
	}

	if limit != 0 {
		sql.Limit(limit)
	}

	err := sql.Find(&transactions).Error
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

// GetTransactionByUserId implements interfaces.TransactionRepository.
func (repo *transactionRepository) GetTransactionsByWalletId(ctx context.Context, walletId uint) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := repo.db.WithContext(ctx).Model(&models.Transaction{}).
		Preload("Receiver").
		Preload("Sender").
		Where("sender_id = ? OR receiver_id = ?", walletId, walletId).
		Order("created_at DESC").
		Limit(10).
		Find(&transactions).
		Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

// BeginTrx implements interfaces.TransactionRepository.
func (repo *transactionRepository) BeginTrx(ctx context.Context) interfaces.TransactionRepository {
	panic("unimplemented")
}

// CommitOrRollback implements interfaces.TransactionRepository.
func (repo *transactionRepository) CommitOrRollback(ctx context.Context) {
	panic("unimplemented")
}

// CreateTransferTransaction implements interfaces.TransactionRepository.
func (repo *transactionRepository) CreateTransferTransaction(ctx context.Context, transaction *models.Transaction) (*models.Transaction, error) {
	err := repo.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var (
			receiverWallet *models.Wallet
			senderWallet   *models.Wallet
		)

		err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Model(&models.Wallet{}).
			Where("id = ?", transaction.SenderID).
			Find(&senderWallet).Error
		if err != nil || senderWallet == nil {
			return err
		}

		err = tx.Clauses(clause.Locking{Strength: "UPDATE"}).Model(&models.Wallet{}).
			Where("id = ?", transaction.ReceiverID).
			Find(&receiverWallet).Error
		if err != nil || receiverWallet == nil {
			return err
		}

		senderWallet.Balance = senderWallet.Balance.Sub(transaction.Amount)
		receiverWallet.Balance = decimal.Sum(receiverWallet.Balance, transaction.Amount)

		err = tx.Table("wallets").
			Where("id = ?", senderWallet.ID).
			Update("balance", senderWallet.Balance).
			Error
		if err != nil {
			return err
		}

		err = tx.Table("wallets").
			Where("id = ?", receiverWallet.ID).
			Update("balance", receiverWallet.Balance).
			Error
		if err != nil {
			return err
		}

		err = tx.Create(&transaction).Error
		if err != nil {
			return err
		}

		return nil
	})
	return transaction, err
}

// CreateTransaction implements interfaces.TransactionRepository.
func (repo *transactionRepository) CreateTopupTransaction(ctx context.Context, transaction *models.Transaction) (*models.Transaction, error) {
	err := repo.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		var wallet *models.Wallet

		err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Model(&models.Wallet{}).
			Where("id = ?", transaction.ReceiverID).
			Find(&wallet).Error
		if err != nil || wallet == nil {
			return err
		}

		wallet.Balance = decimal.Sum(wallet.Balance, transaction.Amount)
		err = tx.Table("wallets").
			Where("id = ?", wallet.ID).
			Update("balance", wallet.Balance).Error
		if err != nil {
			return err
		}

		err = tx.Save(&transaction).Error
		if err != nil {
			return err
		}

		return nil
	})
	return transaction, err
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
