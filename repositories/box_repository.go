package repositories

import (
	"assignment_4/entities/models"
	"assignment_4/enums"
	"assignment_4/interfaces"
	"assignment_4/utils"
	"context"

	"gorm.io/gorm"
)

type boxRepository struct {
	db *gorm.DB
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
