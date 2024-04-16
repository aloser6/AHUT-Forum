package models

type User struct {
	ID       uint `gorm:"primaryKey"`
	Account  int
	Password string
	Username string
	Sex      string
	Grade    string
	College  string
	Major    string
}
