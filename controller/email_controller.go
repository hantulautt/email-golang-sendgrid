package controller

import (
	"email/model"
	"email/service"
	"github.com/gofiber/fiber/v2"
)

type EmailController struct {
	emailService service.EmailService
}

func NewEmailController(service *service.EmailService) EmailController {
	return EmailController{
		emailService: *service,
	}
}

func (controller *EmailController) Route(app *fiber.App) {
	app.Get("/send-email", controller.Send)
	app.Post("/resend-email", controller.Resend)
}

func (controller *EmailController) Send(ctx *fiber.Ctx) error {
	controller.emailService.Send()
	return ctx.JSON(model.WebResponse{
		Code:    fiber.StatusOK,
		Status:  true,
		Message: "Ok",
	})
}

func (controller *EmailController) Resend(ctx *fiber.Ctx) error {
	var request model.EmailRequest
	var message string
	errRequest := ctx.BodyParser(&request)
	if errRequest != nil {
		message = errRequest.Error()
	}
	response := controller.emailService.Resend(request.Uuid)
	if response != nil {
		message = response.Error()
	}
	return ctx.JSON(model.WebResponse{
		Code:    fiber.StatusOK,
		Status:  true,
		Message: message,
	})
}
