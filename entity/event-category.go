package entity

type EventCategory struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	EventID      uint   `gorm:"column:event_id" json:"eventId"`
	CategoryName string `json:"categoryName"`
	IsDeleted    string `json:"isDeleted"`
}

func (EventCategory) TableName() string {
	return "event_category"
}
