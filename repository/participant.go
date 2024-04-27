package repository

import (
	"auth/entity"
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Participant struct {
	db *gorm.DB
}

func NewParticipant(iSportDB *gorm.DB) Participant {
	return Participant{
		db: iSportDB,
	}
}

type ParticipantInterface interface {
	Create(ctx context.Context, participant *entity.Participant) (*entity.Participant, error)
	UpdateSelectedFields(
		ctx context.Context,
		participant *entity.Participant,
		fields ...string,
	) (*entity.Participant, error)
	FindByID(ctx context.Context, id uint) (entity.Participant, error)
}

func (repo Participant) Create(ctx context.Context, participant *entity.Participant) (*entity.Participant, error) {
	err := repo.db.WithContext(ctx).Model(&entity.Participant{}).Omit(clause.Associations).Create(participant).Error
	return participant, err
}

func (repo Participant) UpdateSelectedFields(
	ctx context.Context,
	participant *entity.Participant,
	fields ...string,
) (*entity.Participant, error) {
	err := repo.db.WithContext(ctx).Model(participant).Select(fields).Updates(*participant).Error
	return participant, err
}

func (repo Participant) FindByID(ctx context.Context, id uint) (entity.Participant, error) {
	var participant entity.Participant
	err := repo.db.WithContext(ctx).
		First(&participant, id).
		Error
	return participant, err
}
