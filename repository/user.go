package repository

import (
	"auth/entity"
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User struct {
	db *gorm.DB
}

func NewUser(iSportDB *gorm.DB) User {
	return User{
		db: iSportDB,
	}
}

type UserInterface interface {
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
	UpdateSelectedFields(
		ctx context.Context,
		user *entity.User,
		fields ...string,
	) (*entity.User, error)
	FindByEmail(ctx context.Context, email string) (entity.User, error)
}

func (repo User) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	err := repo.db.WithContext(ctx).Model(&entity.User{}).Omit(clause.Associations).Create(user).Error
	return user, err
}

func (repo User) UpdateSelectedFields(
	ctx context.Context,
	user *entity.User,
	fields ...string,
) (*entity.User, error) {
	err := repo.db.WithContext(ctx).Model(user).Select(fields).Updates(*user).Error
	return user, err
}

func (repo User) FindByEmail(ctx context.Context, email string) (entity.User, error) {
	var user entity.User
	err := repo.db.WithContext(ctx).
		Where("email = ?", email).
		First(&user).
		Error
	return user, err
}
