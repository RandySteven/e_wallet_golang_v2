package utils

import (
	"assignment_4/enums"
	"context"

	"gorm.io/gorm"
)

func SelectQuery[T any](ctx context.Context, db *gorm.DB) ([]T, error) {
	var entities []T
	err := db.WithContext(ctx).Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return entities, nil
}

func SaveQuery[T any](ctx context.Context, db *gorm.DB, entity *T, action string) (*T, error) {
	query := db.WithContext(ctx)
	switch action {
	case enums.Create:
		query.Create(&entity)
	case enums.Update:
		query.Updates(&entity)
	}
	err := query.Error
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func GetById[T any](ctx context.Context, db *gorm.DB, id uint) (entity *T, err error) {
	err = db.WithContext(ctx).Model(&entity).Where("id = ?", id).Scan(&entity).Error
	if err != nil {
		return nil, err
	}
	return entity, nil
}
