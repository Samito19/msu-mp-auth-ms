package models

type User struct {
	EmailAddress string `json:"emailAddress"`
	Password     string `json:"password"`
}
