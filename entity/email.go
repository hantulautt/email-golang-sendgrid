package entity

import "gorm.io/gorm"

type Email struct {
	Uuid         string `gorm:"primaryKey"`
	EmailTypeId  uint
	Key          string
	Subject      string
	EmailContent string
	Attachment   string
	FromEmail    string
	FromName     string
	ToEmail      string
	ToName       string
}

func (Email) TableName() string {
	return "emails"
}

func Unsent(db *gorm.DB) *gorm.DB {
	return db.Where("status", 0)
}
