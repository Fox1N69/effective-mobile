package models

import "time"

type User struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	Surname        string    `json:"surname"`
	Patronymic     string    `json:"patronymic"`
	PassportNumber string    `json:"passportNumber"`
	Address        string    `json:"text"`
	CreatedAt      time.Time `json:"createdAt,omitempty"`
	UpdatedAt      time.Time `json:"updatedAt,omitempty"`
}
