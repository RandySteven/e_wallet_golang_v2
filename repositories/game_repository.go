package repositories

import (
	"assignment_4/entities/models"
	"assignment_4/enums"
	"assignment_4/interfaces"
	"assignment_4/utils"
	"context"

	"gorm.io/gorm"
)

type gameRepository struct {
	db *gorm.DB
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
func (*gameRepository) GetById(ctx context.Context, id uint) (*models.Game, error) {
	panic("unimplemented")
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
