package models

import "time"

const UnsetValue = "__UNSET__"

type UserCreateDto struct {
	Id         int       `json:"id"`
	FirstName  string    `json:"first_name"`
	SecondName string    `json:"second_name"`
	Email      string    `json:"email" validate:"required,email"`
	Password   string    `json:"password" validate:"required,min=8"`
	Role       Role      `json:"role_id"`
	CreatedAt  time.Time `json:"created_at"`
}

type UserUpdateDto struct {
	Id         int       `json:"id"`
	FirstName  string    `json:"first_name"`
	SecondName string    `json:"second_name"`
	Email      string    `json:"email" validate:"required,email"`
	Password   string    `json:"password"`
	Role       Role      `json:"role_id"`
	CreatedAt  time.Time `json:"created_at"`
}

func (u UserCreateDto) ToUserUpdateDto() UserUpdateDto {
	return UserUpdateDto{
		Id:         u.Id,
		FirstName:  u.FirstName,
		SecondName: u.SecondName,
		Email:      u.Email,
		Role:       u.Role,
		CreatedAt:  u.CreatedAt,
	}
}
