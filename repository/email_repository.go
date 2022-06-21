package repository

import "email/entity"

type EmailRepository interface {
	GetAll() (response []entity.Email)
	Update(uuid string, err error)
	GetEmail(uuid string) (response entity.Email)
}
