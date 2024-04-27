package repository

import (
	"auth/entity"
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type EventCategory struct {
	db *gorm.DB
}

func NewEventCategory(iSportDB *gorm.DB) EventCategory {
	return EventCategory{
		db: iSportDB,
	}
}

type EventCategoryInterface interface {
	Create(ctx context.Context, eventCategory *entity.EventCategory) (*entity.EventCategory, error)
	UpdateSelectedFields(
		ctx context.Context,
		eventCategory *entity.EventCategory,
		fields ...string,
	) (*entity.Event, error)
	FindByEventID(ctx context.Context, id uint) (entity.EventCategory, error)
}

func (repo EventCategory) Create(ctx context.Context, eventCategory *entity.EventCategory) (*entity.EventCategory, error) {
	err := repo.db.WithContext(ctx).Model(&entity.Event{}).Omit(clause.Associations).Create(eventCategory).Error
	return eventCategory, err
}

func (repo EventCategory) UpdateSelectedFields(
	ctx context.Context,
	eventCategory *entity.EventCategory,
	fields ...string,
) (*entity.EventCategory, error) {
	err := repo.db.WithContext(ctx).Model(eventCategory).Select(fields).Updates(*eventCategory).Error
	return eventCategory, err
}

func (repo EventCategory) FindByEventID(ctx context.Context, eventID uint) (entity.EventCategory, error) {
	var event entity.EventCategory
	err := repo.db.WithContext(ctx).
		Where("event_id = ?", eventID).
		First(&event).
		Error
	return event, err
}
