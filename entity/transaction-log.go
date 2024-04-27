package entity

import "time"

type TransactionLog struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	TransactionID  string    `json:"transactionId"`
	RefID          string    `json:"refId"`
	ParticipantID  uint      `gorm:"column:participant_id" json:"participantId"`
	Amount         string    `json:"amount"`
	PaymentStatus  string    `json:"paymentStatus"`
	PaymentMethod  string    `json:"paymentMethod"`
	VirtualAccount string    `json:"virtualAccount"`
	ExpiredTime    time.Time `json:"expiredTime"`
	PaymentLink    string    `json:"paymentLink"`
	LastResponse   string    `json:"LastResponse"`
	Response       string    `json:"response"`
	CreatedTime    time.Time `json:"createdTime"`
	UpdatedTime    time.Time `json:"updatedTime"`
}

func (TransactionLog) TableName() string {
	return "transaction_log"
}
