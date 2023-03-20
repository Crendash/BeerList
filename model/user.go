package model

import (
	"Beer-BackendV1/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name    string  `gorm:"size:255,not null" json:"name"`
	Surname string  `gorm:"size:255,not null" json:"surname"`
	Balance float64 `gorm:"not null" json:"balance"`
}

func (user *User) Save() (*User, error) {
	err := database.Database.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func ReturnAllUsers() ([]User, error) {
	var users []User
	result := database.Database.Find(&users)
	if err := result.Error; err != nil {
		return nil, err
	}
	return users, nil
}

func Delete(user User) {
	database.Database.Delete(user)
}

func FindUserByName(name, surname string) (User, error) {
	var user User
	err := database.Database.Where("name=? AND surname=?", name, surname).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func UpdateBalance(user User, balance float64) {
	database.Database.Model(&User{}).Where("name=? AND surname=?", user.Name, user.Surname).Update("balance", balance)
}
