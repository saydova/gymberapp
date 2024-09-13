package models

import "gorm.io/gorm"

type Trainer struct {
	gorm.Model
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Specialization string `json:"specialization"`
	PhoneNumber    string `json:"phone_number"`
	Email          string `json:"email" gorm:"unique"`
	HireDate       string `json:"hire_date"`
}
