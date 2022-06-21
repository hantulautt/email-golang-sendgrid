package repository

import (
	"email/entity"
	"gorm.io/gorm"
)

type emailRepositoryImpl struct {
	Db *gorm.DB
}

func (email emailRepositoryImpl) Update(uuid string, err error) {
	update := entity.Email{}
	if err != nil {
		email.Db.Model(&update).Where("uuid=?", uuid).Update("message", err.Error()).Unscoped()
	} else {
		email.Db.Model(&update).Where("uuid=?", uuid).Update("status", 1).Unscoped()
	}
}

func NewEmailRepository(database *gorm.DB) EmailRepository {
	return &emailRepositoryImpl{
		Db: database,
	}
}

func (email emailRepositoryImpl) GetAll() (response []entity.Email) {
	email.Db.Where("status", 0).Find(&response)
	return response
}

func (email emailRepositoryImpl) GetEmail(uuid string) (response entity.Email) {
	email.Db.Where("uuid=?", uuid).First(&response)
	return response
}
