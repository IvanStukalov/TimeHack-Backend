package models

// easyjson -all ./internal/models/user.go

type User struct {
	ID           int    `json:"ID"`
	Login        string `json:"Login"`
	Username     string `json:"Username"`
	PasswordHash string `json:"PasswordHash"`
}
