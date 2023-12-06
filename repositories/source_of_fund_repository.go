package repositories

import (
	"assignment_4/entities/models"
	"assignment_4/interfaces"
	"assignment_4/utils"
	"context"

	"gorm.io/gorm"
)

type sourceOfFundRepo struct {
	db *gorm.DB
}

// Count implements interfaces.SourceOfFundRepository.
func (repo *sourceOfFundRepo) Count(ctx context.Context) (uint, error) {
	return utils.CountTotalItems[models.SourceOfFund](ctx, repo.db, &models.SourceOfFund{})
}

// BeginTrx implements interfaces.SourceOfFundRepository.
func (repo *sourceOfFundRepo) BeginTrx(ctx context.Context) interfaces.SourceOfFundRepository {
	panic("unimplemented")
}

// CommitOrRollback implements interfaces.SourceOfFundRepository.
func (repo *sourceOfFundRepo) CommitOrRollback(ctx context.Context) {
	panic("unimplemented")
}

// FindAll implements interfaces.SourceOfFundRepository.
func (repo *sourceOfFundRepo) FindAll(ctx context.Context) ([]models.SourceOfFund, error) {
	panic("unimplemented")
}

// GetById implements interfaces.SourceOfFundRepository.
func (repo *sourceOfFundRepo) GetById(ctx context.Context, id uint) (*models.SourceOfFund, error) {
	panic("unimplemented")
}

// GetSourceOfFundBySource implements interfaces.SourceOfFundRepository.
func (repo *sourceOfFundRepo) GetSourceOfFundBySource(ctx context.Context, source string) (*models.SourceOfFund, error) {
	var sourceFund *models.SourceOfFund
	err := repo.db.WithContext(ctx).Model(&models.SourceOfFund{}).
		Where("source = ?", source).
		Scan(&sourceFund).
		Error
	if err != nil {
		return nil, err
	}
	return sourceFund, nil
}

// Save implements interfaces.SourceOfFundRepository.
func (repo *sourceOfFundRepo) Save(ctx context.Context, entity *models.SourceOfFund) (*models.SourceOfFund, error) {
	panic("unimplemented")
}

// Update implements interfaces.SourceOfFundRepository.
func (repo *sourceOfFundRepo) Update(ctx context.Context, entity *models.SourceOfFund) (*models.SourceOfFund, error) {
	panic("unimplemented")
}

func NewSourceOfFundRepo(db *gorm.DB) *sourceOfFundRepo {
	return &sourceOfFundRepo{db: db}
}

var _ interfaces.SourceOfFundRepository = &sourceOfFundRepo{}
