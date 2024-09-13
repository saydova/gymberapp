package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	MemberID      uint    `json:"member_id"`
	PaymentDate   string  `json:"payment_date"`
	Amount        float64 `json:"amount"`
	PaymentMethod string  `json:"payment_method"`
	Status        string  `json:"status"`
	Member        Member  `gorm:"foreignKey:MemberID"`
}
