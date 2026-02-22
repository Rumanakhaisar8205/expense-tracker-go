package models

import "gorm.io/gorm"

type ExpenseShare struct {
	gorm.Model
	ExpenseID uint
	UserID    uint
	Share     float64 `gorm:"type:decimal(10,2)"`
}
