package model

type User struct {
	ID       int `gorm:"primary_key"`
	Username string
	Password string
	Email    string
	Gender   string
	Birthday string
	City     string
	Address  string
}
