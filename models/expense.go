package models

import "gorm.io/gorm"

type Expense struct {
	gorm.Model
	GroupID uint
	PaidBy  uint
	Amount  float64 `gorm:"type:decimal(10,2)"` // FIXED
	Note    string
}
