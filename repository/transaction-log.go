package repository

import (
	"context"
	"github.com/iman-khaeruddin/isport-auth/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TransactionLog struct {
	db *gorm.DB
}

func NewTransactionLog(iSportDB *gorm.DB) TransactionLog {
	return TransactionLog{
		db: iSportDB,
	}
}

type TransactionLogInterface interface {
	Create(ctx context.Context, transactionLog *entity.TransactionLog) (*entity.TransactionLog, error)
	UpdateSelectedFields(
		ctx context.Context,
		transactionLog *entity.TransactionLog,
		fields ...string,
	) (*entity.TransactionLog, error)
	FindByRefID(ctx context.Context, refID uint) (entity.TransactionLog, error)
}

func (repo TransactionLog) Create(ctx context.Context, transactionLog *entity.TransactionLog) (*entity.TransactionLog, error) {
	err := repo.db.WithContext(ctx).Model(&entity.TransactionLog{}).Omit(clause.Associations).Create(transactionLog).Error
	return transactionLog, err
}

func (repo TransactionLog) UpdateSelectedFields(
	ctx context.Context,
	transactionLog *entity.TransactionLog,
	fields ...string,
) (*entity.TransactionLog, error) {
	err := repo.db.WithContext(ctx).Model(transactionLog).Select(fields).Updates(*transactionLog).Error
	return transactionLog, err
}

func (repo TransactionLog) FindByRefID(ctx context.Context, refID uint) (entity.TransactionLog, error) {
	var transactionLog entity.TransactionLog
	err := repo.db.WithContext(ctx).
		Where("ref_id = ?", refID).
		First(&transactionLog).
		Error
	return transactionLog, err
}
