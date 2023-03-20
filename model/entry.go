package model

import (
	"Beer-BackendV1/database"
	"gorm.io/gorm"
)

type Entry struct {
	gorm.Model
	Type    string `gorm:"type:text" json:"type"`
	Amount  int64  `json:"amount"`
	Name    string `gorm:"size:255,not null" json:"name"`
	Surname string `gorm:"size:255,not null" json:"surname"`
}

func (entry *Entry) Save() (*Entry, error) {
	err := database.Database.Create(&entry).Error
	if err != nil {
		return &Entry{}, err
	}
	return entry, nil
}

func ReturnAllEntries() ([]Entry, error) {
	var entries []Entry
	result := database.Database.Find(&entries)
	if err := result.Error; err != nil {
		return nil, err
	}
	return entries, nil
}

func ReturnAllEntriesOfUser(userName UserName) ([]Entry, error) {
	var entries []Entry
	result := database.Database.Where("name = ? AND surname = ?", userName.Name, userName.Surname).Find(&entries)
	if err := result.Error; err != nil {
		return nil, err
	}
	return entries, nil
}
