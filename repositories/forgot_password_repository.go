package repositories

import (
	"assignment_4/entities/models"
	"assignment_4/enums"
	"assignment_4/interfaces"
	"assignment_4/utils"
	"context"

	"gorm.io/gorm"
)

type forgotPasswordRepository struct {
	db *gorm.DB
}

// BeginTrx implements interfaces.ForgotPasswordRepository.
func (repo *forgotPasswordRepository) BeginTrx(ctx context.Context) interfaces.ForgotPasswordRepository {
	panic("unimplemented")
}

// CommitOrRollback implements interfaces.ForgotPasswordRepository.
func (repo *forgotPasswordRepository) CommitOrRollback(ctx context.Context) {
	panic("unimplemented")
}

// FindAll implements interfaces.ForgotPasswordRepository.
func (repo *forgotPasswordRepository) FindAll(ctx context.Context) ([]models.ForgotPasswordToken, error) {
	panic("unimplemented")
}

// GetById implements interfaces.ForgotPasswordRepository.
func (repo *forgotPasswordRepository) GetById(ctx context.Context, id uint) (*models.ForgotPasswordToken, error) {
	panic("unimplemented")
}

// Save implements interfaces.ForgotPasswordRepository.
func (repo *forgotPasswordRepository) Save(ctx context.Context, entity *models.ForgotPasswordToken) (*models.ForgotPasswordToken, error) {
	return utils.SaveQuery[models.ForgotPasswordToken](ctx, repo.db, entity, enums.Create)
}

// Update implements interfaces.ForgotPasswordRepository.
func (repo *forgotPasswordRepository) Update(ctx context.Context, entity *models.ForgotPasswordToken) (*models.ForgotPasswordToken, error) {
	panic("unimplemented")
}

func NewForgotPasswordRepository(db *gorm.DB) *forgotPasswordRepository {
	return &forgotPasswordRepository{db: db}
}

var _ interfaces.ForgotPasswordRepository = &forgotPasswordRepository{}
