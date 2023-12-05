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

// UpdateUserPassword implements interfaces.ForgotPasswordRepository.
func (repo *forgotPasswordRepository) UpdateUserPassword(ctx context.Context, token *models.ForgotPasswordToken, password string) (*models.User, error) {
	var user *models.User
	err := repo.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&models.ForgotPasswordToken{}).
			Where("id = ?", token.ID).
			Update("is_valid", false).Error
		if err != nil {
			return err
		}

		user, err := utils.GetById[models.User](ctx, repo.db, token.UserID)
		if err != nil {
			return err
		}

		user.Password = password
		_, err = utils.SaveQuery[models.User](ctx, repo.db, user, enums.Update)
		if err != nil {
			return err
		}

		return nil
	})

	return user, err
}

// GetPasswordTokenByToken implements interfaces.ForgotPasswordRepository.
func (repo *forgotPasswordRepository) GetPasswordTokenByToken(ctx context.Context, resetToken string) (*models.ForgotPasswordToken, error) {
	var token *models.ForgotPasswordToken
	err := repo.db.WithContext(ctx).Where("reset_token = ?", resetToken).Find(&token).Error
	if err != nil {
		return nil, err
	}
	return token, nil
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
