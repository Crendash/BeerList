package model

type UserName struct {
	Name    string `json:"name" binding:"required"`
	Surname string `json:"surname" binding:"required"`
}
