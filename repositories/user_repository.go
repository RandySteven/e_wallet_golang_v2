package repositories

import (
	"assignment_4/entities/models"
	"assignment_4/enums"
	"assignment_4/interfaces"
	"assignment_4/utils"
	"context"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// RegisterUser implements interfaces.UserRepository.
func (repo *userRepository) RegisterUser(ctx context.Context, user *models.User) (*models.User, error) {
	err := repo.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		user, err := repo.Save(ctx, user)
		if err != nil {
			return err
		}

		wallet := &models.Wallet{
			Balance: decimal.NewFromInt(0),
			UserID:  user.ID,
		}

		err = repo.db.Save(&wallet).Error
		if err != nil {
			return err
		}

		return nil
	})

	return user, err
}

// BeginTrx implements interfaces.UserRepository.
func (repo *userRepository) BeginTrx(ctx context.Context) interfaces.UserRepository {
	tx := repo.db.Begin()
	return &userRepository{
		db: tx,
	}
}

// CommitOrRollback implements interfaces.UserRepository.
func (repo *userRepository) CommitOrRollback(ctx context.Context) {
	err := recover()
	if err != nil {
		repo.db.Rollback()
		return
	}
	repo.db.Commit()
}

// GetByEmail implements interfaces.UserRepository.
func (repo *userRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user *models.User
	err := repo.db.WithContext(ctx).Model(&models.User{}).
		Where("email = ?", email).
		Scan(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// FindAll implements interfaces.UserRepository.
func (repo *userRepository) FindAll(ctx context.Context) ([]models.User, error) {
	return utils.SelectQuery[models.User](ctx, repo.db)
}

// GetById implements interfaces.UserRepository.
func (repo *userRepository) GetById(ctx context.Context, id uint) (*models.User, error) {
	return utils.GetById[models.User](ctx, repo.db, id)
}

// Save implements interfaces.UserRepository.
func (repo *userRepository) Save(ctx context.Context, entity *models.User) (*models.User, error) {
	return utils.SaveQuery[models.User](ctx, repo.db, entity, enums.Create)
}

// Update implements interfaces.UserRepository.
func (repo *userRepository) Update(ctx context.Context, entity *models.User) (*models.User, error) {
	return utils.SaveQuery[models.User](ctx, repo.db, entity, enums.Update)
}

var _ interfaces.UserRepository = &userRepository{}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}
