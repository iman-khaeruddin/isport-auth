package repository

import (
	"context"
	"github.com/iman-khaeruddin/isport-auth/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Event struct {
	db *gorm.DB
}

func NewEvent(iSportDB *gorm.DB) Event {
	return Event{
		db: iSportDB,
	}
}

type EventInterface interface {
	Create(ctx context.Context, event *entity.Event) (*entity.Event, error)
	UpdateSelectedFields(
		ctx context.Context,
		event *entity.Event,
		fields ...string,
	) (*entity.Event, error)
	FindByID(ctx context.Context, id uint) (entity.Event, error)
}

func (repo Event) Create(ctx context.Context, event *entity.Event) (*entity.Event, error) {
	err := repo.db.WithContext(ctx).Model(&entity.Event{}).Omit(clause.Associations).Create(event).Error
	return event, err
}

func (repo Event) UpdateSelectedFields(
	ctx context.Context,
	event *entity.Event,
	fields ...string,
) (*entity.Event, error) {
	err := repo.db.WithContext(ctx).Model(event).Select(fields).Updates(*event).Error
	return event, err
}

func (repo Event) FindByID(ctx context.Context, id uint) (entity.Event, error) {
	var event entity.Event
	err := repo.db.WithContext(ctx).
		First(&event, id).
		Error
	return event, err
}
