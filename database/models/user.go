package models

type User struct {
	EmailAddress string `json:"emailAddress" gorm:"<-"`
	Password     string `json:"password"`
}
