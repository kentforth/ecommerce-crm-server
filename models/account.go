package models

type Account struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	DarkMode bool   `json:"dark_mode"`
}
