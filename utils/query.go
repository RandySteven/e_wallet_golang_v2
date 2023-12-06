package utils

import (
	"context"
	"fmt"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/enums"

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

func CountTotalItems[T any](ctx context.Context, db *gorm.DB, entity *T) (uint, error) {
	var res int64 = 0
	err := db.WithContext(ctx).Model(&entity).Count(&res).Error
	out := uint(res)
	if err != nil {
		return out, err
	}
	return out, nil
}

func CountTotalItemsCondition[T any](ctx context.Context, db *gorm.DB, entity *T, field string, value string) (uint, error) {
	var res int64 = 0
	condition := fmt.Sprintf("%s = ?", field)
	err := db.WithContext(ctx).Model(&entity).Count(&res).Where(condition, value).Error
	out := uint(res)
	if err != nil {
		return out, err
	}
	return out, nil
}
