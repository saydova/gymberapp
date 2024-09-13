package models

import "gorm.io/gorm"

type Enrollment struct {
	gorm.Model
	MemberID       uint   `json:"member_id"`
	ClassID        uint   `json:"class_id"`
	EnrollmentDate string `json:"enrollment_date"`
	Member         Member `gorm:"foreignKey:MemberID"`
	Class          Class  `gorm:"foreignKey:ClassID"`
}
