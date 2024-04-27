package entity

import "time"

type Participant struct {
	ID            uint          `gorm:"primaryKey" json:"id"`
	UserID        uint          `gorm:"column:user_id" json:"userId"`
	EventID       uint          `gorm:"column:event_id" json:"eventId"`
	CategoryID    uint          `gorm:"column:category_id" json:"categoryId"`
	FullName      string        `json:"fullName"`
	NickName      string        `json:"nickName"`
	NumberPlate   string        `json:"numberPlate"`
	DOB           time.Time     `json:"dob"`
	DOC           string        `json:"doc"`
	Team          string        `json:"team"`
	Community     string        `json:"community"`
	PaymentStatus string        `json:"paymentStatus"`
	CreatedTime   time.Time     `json:"createdTime"`
	IsDeleted     bool          `json:"isDeleted"`
	User          User          `gorm:"foreignKey:UserID"`
	Event         Event         `gorm:"foreignKey:EventID"`
	EventCategory EventCategory `gorm:"foreignKey:CategoryID"`
}

func (Participant) TableName() string {
	return "participant"
}
