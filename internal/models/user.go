package models

import "time"

type User struct {
	Id         int       `json:"id" validate:"required"`
	FirstName  string    `json:"first_name"  validate:"required"`
	SecondName string    `json:"second_name"`
	Email      string    `json:"email" validate:"required,email"`
	Password   string    `json:"password" validate:"required,min=8"`
	Role       Role      `json:"role_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
