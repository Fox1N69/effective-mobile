package models

type User struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	Surname        string `json:"surname"`
	Patronymic     string `json:"patronymic"`
	PassportNumber string `json:"passportNumber"`
	Address        string `json:"text"`
}
