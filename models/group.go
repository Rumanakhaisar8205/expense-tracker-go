package models

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name  string `json:"name"`
	Users []User `gorm:"many2many:group_users;"`
}
