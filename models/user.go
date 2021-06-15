package models

import "time"

type User struct {
	ID        uint `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Email     string
	Password  string
	Gender    string
	Enabled   bool
	CreatedAt time.Time
}
