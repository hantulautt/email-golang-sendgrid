package service

type EmailService interface {
	Send()
	Resend(uuid string) (err error)
}
