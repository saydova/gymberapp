package models

import "gorm.io/gorm"

type Class struct {
	gorm.Model
	ClassName   string  `json:"class_name"`
	Description string  `json:"description"`
	Schedule    string  `json:"schedule"`
	TrainerID   uint    `json:"trainer_id"`
	Trainer     Trainer `gorm:"foreignKey:TrainerID"`
}
