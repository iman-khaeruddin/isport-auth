package entity

type UserEvent struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	UserID    uint   `gorm:"column:user_id" json:"userId"`
	EventID   uint   `gorm:"column:event_id" json:"eventId"`
	ChildID   uint   `gorm:"column:child_id" json:"childId"`
	EventName string `json:"eventName"`
}

func (UserEvent) TableName() string {
	return "user_event"
}
