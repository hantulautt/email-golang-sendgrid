package service

import (
	"email/entity"
	"email/helper"
	"email/repository"
	"fmt"
	"github.com/pkg/errors"
)

type emailServiceImpl struct {
	EmailRepository repository.EmailRepository
}

func NewEmailService(repository *repository.EmailRepository) EmailService {
	return &emailServiceImpl{
		EmailRepository: *repository,
	}
}

func (email emailServiceImpl) Send() {
	var response []entity.Email
	response = email.EmailRepository.GetAll()
	for _, item := range response {
		err := helper.Send(item)
		fmt.Println(err)
		fmt.Println(item.Uuid)
		email.EmailRepository.Update(item.Uuid, err)
	}
}

func (email emailServiceImpl) Resend(uuid string) (err error) {
	var item entity.Email
	item = email.EmailRepository.GetEmail(uuid)
	if item.Uuid == "" {
		return errors.New("Record not found")
	}
	err = helper.Send(item)
	email.EmailRepository.Update(item.Uuid, err)
	return err
}
