package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	EmailAddress string `json:"emailAddress" gorm:"<-"`
	Password     string `json:"password"`
}
