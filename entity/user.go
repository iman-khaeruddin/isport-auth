package entity

import "time"

type User struct {
	ID                  uint      `gorm:"primaryKey" json:"id"`
	Email               string    `json:"email"`
	PhoneNumber         string    `json:"phoneNumber"`
	FirstName           string    `json:"firstName"`
	LastName            string    `json:"lastName"`
	Password            string    `json:"password"`
	LastLogin           time.Time `gorm:"default:null" json:"lastLogin"`
	TokenForgotPassword string    `gorm:"default:null" json:"tokenForgotPassword"`
	GoogleID            string    `gorm:"default:null" json:"googleId"`
	CreatedTime         time.Time `json:"createdTime"`
	IsDeleted           bool      `json:"isDeleted"`
}

func (User) TableName() string {
	return "user"
}
