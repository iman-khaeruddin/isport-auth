package entity

import "time"

type Event struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	EventName string    `json:"eventName"`
	EventDate time.Time `json:"eventDate"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	Poster    string    `json:"poster"`
	IsDeleted bool      `json:"isDeleted"`
}

func (Event) TableName() string {
	return "event"
}
