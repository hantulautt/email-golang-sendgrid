package exception

import (
	"email/model"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {

	_, ok := err.(ValidationError)
	if ok {
		return ctx.JSON(model.WebResponse{
			Code:    400,
			Status:  false,
			Message: "BAD_REQUEST",
			Data:    err.Error(),
		})
	}

	return ctx.JSON(model.WebResponse{
		Code:    500,
		Status:  false,
		Message: "INTERNAL_SERVER_ERROR",
		Data:    err.Error(),
	})
}
