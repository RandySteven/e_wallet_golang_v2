package repositories

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/entities/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/enums"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/interfaces"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/utils"

	"gorm.io/gorm"
)

type boxRepository struct {
	db *gorm.DB
}

// Count implements interfaces.BoxRepository.
func (repo *boxRepository) Count(ctx context.Context) (uint, error) {
	return utils.CountTotalItems[models.Box](ctx, repo.db, &models.Box{})
}

// GetNineRandomBoxes implements interfaces.BoxRepository.
func (repo *boxRepository) GetNineRandomBoxes(ctx context.Context) ([]models.Box, error) {
	var boxes []models.Box
	err := repo.db.WithContext(ctx).Order("random()").Limit(9).Find(&boxes).Error
	if err != nil {
		return nil, err
	}
	return boxes, nil
}

// BeginTrx implements interfaces.BoxRepository.
func (repo *boxRepository) BeginTrx(ctx context.Context) interfaces.BoxRepository {
	panic("unimplemented")
}

// CommitOrRollback implements interfaces.BoxRepository.
func (repo *boxRepository) CommitOrRollback(ctx context.Context) {
	panic("unimplemented")
}

// FindAll implements interfaces.BoxRepository.
func (repo *boxRepository) FindAll(ctx context.Context) ([]models.Box, error) {
	panic("unimplemented")
}

// GetById implements interfaces.BoxRepository.
func (repo *boxRepository) GetById(ctx context.Context, id uint) (*models.Box, error) {
	return utils.GetById[models.Box](ctx, repo.db, id)
}

// Save implements interfaces.BoxRepository.
func (repo *boxRepository) Save(ctx context.Context, entity *models.Box) (*models.Box, error) {
	return utils.SaveQuery[models.Box](ctx, repo.db, entity, enums.Create)
}

// Update implements interfaces.BoxRepository.
func (repo *boxRepository) Update(ctx context.Context, entity *models.Box) (*models.Box, error) {
	panic("unimplemented")
}

func NewBoxRepository(db *gorm.DB) *boxRepository {
	return &boxRepository{db: db}
}

var _ interfaces.BoxRepository = &boxRepository{}
