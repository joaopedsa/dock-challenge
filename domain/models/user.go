package models

import "time"

type User struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name" validate:"required,min=3,max=64"`
	CPF       string    `json:"cpf" validate:"required,min=11,max=11"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
