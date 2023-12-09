package repositories

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/entities/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/enums"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/interfaces"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/utils"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type gameRepository struct {
	db *gorm.DB
}

// CreateRewardTransaction implements interfaces.GameRepository.
func (repo *gameRepository) CreateRewardTransaction(ctx context.Context, game *models.Game) (*models.Game, error) {
	err := repo.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var (
			wallet *models.Wallet
			winBox *models.Box
		)

		err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Model(&models.Wallet{}).
			Where("user_id = ?", game.UserID).
			Find(&wallet).Error
		if err != nil || wallet == nil {
			return err
		}

		err = tx.Model(&models.Box{}).
			Where("id = ?", game.WinBoxID).
			Find(&winBox).Error
		if err != nil || winBox == nil {
			return err
		}

		err = tx.Model(&models.User{}).Where("id = ?", game.UserID).
			Update("chance", gorm.Expr("chance - ?", 1)).Error
		if err != nil {
			return err
		}

		wallet.Balance = decimal.Sum(wallet.Balance, winBox.Amount)
		err = tx.Table("wallets").
			Where("id = ?", wallet.ID).
			Update("balance", wallet.Balance).Error
		if err != nil {
			return err
		}

		transaction := &models.Transaction{
			SenderID:       wallet.ID,
			ReceiverID:     wallet.ID,
			SourceOfFundID: 4,
			Amount:         winBox.Amount,
			Description:    "Top up from Reward",
		}

		err = tx.Create(&transaction).Error
		if err != nil {
			return err
		}

		err = tx.Table("games").Where("id = ?", game.ID).Update("win_box_id", game.WinBoxID).Error
		if err != nil {
			return err
		}

		return nil
	})

	return game, err
}

// Count implements interfaces.GameRepository.
func (repo *gameRepository) Count(ctx context.Context) (uint, error) {
	return utils.CountTotalItems[models.Game](ctx, repo.db, &models.Game{})
}

// BeginTrx implements interfaces.GameRepository.
func (*gameRepository) BeginTrx(ctx context.Context) interfaces.GameRepository {
	panic("unimplemented")
}

// CommitOrRollback implements interfaces.GameRepository.
func (*gameRepository) CommitOrRollback(ctx context.Context) {
	panic("unimplemented")
}

// FindAll implements interfaces.GameRepository.
func (*gameRepository) FindAll(ctx context.Context) ([]models.Game, error) {
	panic("unimplemented")
}

// GetById implements interfaces.GameRepository.
func (repo *gameRepository) GetById(ctx context.Context, id uint) (*models.Game, error) {
	return utils.GetById[models.Game](ctx, repo.db, id)
}

// Save implements interfaces.GameRepository.
func (repo *gameRepository) Save(ctx context.Context, entity *models.Game) (*models.Game, error) {
	return utils.SaveQuery[models.Game](ctx, repo.db, entity, enums.Create)
}

// Update implements interfaces.GameRepository.
func (*gameRepository) Update(ctx context.Context, entity *models.Game) (*models.Game, error) {
	panic("unimplemented")
}

func NewGameRepository(db *gorm.DB) *gameRepository {
	return &gameRepository{db: db}
}

var _ interfaces.GameRepository = &gameRepository{}
