package entity

import "time"

type Child struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	UserID      uint      `gorm:"column:user_id" json:"userId"`
	FullName    string    `json:"fullName"`
	NickName    string    `json:"nickName"`
	NumberPlate string    `json:"numberPlate"`
	DOB         time.Time `json:"dob"`
	DOC         string    `json:"doc"`
	Team        string    `json:"team"`
	Community   string    `json:"community"`
	CreatedTime time.Time `json:"createdTime"`
	IsDeleted   bool      `json:"isDeleted"`
	User        User      `gorm:"foreignKey:UserID"`
}

func (Child) TableName() string {
	return "child"
}
