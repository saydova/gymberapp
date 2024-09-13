package models

import "gorm.io/gorm"

type Member struct {
	gorm.Model
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Dob            string `json:"dob"`
	Email          string `json:"email" gorm:"unique"`
	PhoneNumber    string `json:"phone_number"`
	JoinDate       string `json:"join_date"`
	MembershipType string `json:"membership_type"`
	Status         string `json:"status"`
}
