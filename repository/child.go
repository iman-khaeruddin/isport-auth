package repository

import (
	"auth/entity"
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Child struct {
	db *gorm.DB
}

func NewChild(iSportDB *gorm.DB) Child {
	return Child{
		db: iSportDB,
	}
}

type ChildInterface interface {
	Create(ctx context.Context, child *entity.Child) (*entity.Child, error)
	UpdateSelectedFields(
		ctx context.Context,
		child *entity.Child,
		fields ...string,
	) (*entity.Child, error)
	FindByParentID(ctx context.Context, userID uint) (entity.Child, error)
}

func (repo Child) Create(ctx context.Context, child *entity.Child) (*entity.Child, error) {
	err := repo.db.WithContext(ctx).Model(&entity.Child{}).Omit(clause.Associations).Create(child).Error
	return child, err
}

func (repo Child) UpdateSelectedFields(
	ctx context.Context,
	child *entity.Child,
	fields ...string,
) (*entity.Child, error) {
	err := repo.db.WithContext(ctx).Model(child).Select(fields).Updates(*child).Error
	return child, err
}

func (repo Child) FindByParentID(ctx context.Context, userID uint) (entity.Child, error) {
	var event entity.Child
	err := repo.db.WithContext(ctx).
		Where("user_id = ?", userID).
		First(&event).
		Error
	return event, err
}
