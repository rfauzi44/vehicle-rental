package model

import (
	"time"
)

type User struct {
	UserID      string     `gorm:"type:uuid; primaryKey; default:uuid_generate_v4()" json:"user_id,omitempty" valid:"-"`
	FullName    string     `json:"full_name" valid:"type(string)"`
	Email       string     `json:"email" valid:"email"`
	Password    string     `json:"password,omitempty" valid:"type(string)"`
	Gender      string     `json:"gender" valid:"type(string)"`
	Address     string     `json:"address,omitempty" valid:"type(string)"`
	Birthday    *time.Time `json:"birthday,omitempty"`
	PhoneNumber string     `json:"phone_number" valid:"type(string)"`
	CreatedAt   time.Time  `json:"created_at" valid:"-"`
	UpdatedAt   time.Time  `json:"updated_at,omitempty" valid:"-" `
}

type Users []User
