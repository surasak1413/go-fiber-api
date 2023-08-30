package handler

import (
	"loly/request"
	"loly/response"
	"loly/services"

	"github.com/gofiber/fiber/v2"
)

func RegisterHandler(c *fiber.Ctx) error {
	tag := "Register Handler"

	req := &request.AuthRequest{}
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Status:  fiber.StatusBadRequest,
			Tag:     tag,
			Message: "error",
			Error:   err.Error(),
		})
	}

	if err := req.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Status:  fiber.StatusBadRequest,
			Tag:     tag,
			Message: "validation error",
			Error:   err.Error(),
		})
	}

	sv := services.NewService()
	err := sv.Register(*req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Status:  fiber.StatusBadRequest,
			Tag:     tag,
			Message: "error",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.SuccessResponse{
		Status:  fiber.StatusOK,
		Tag:     tag,
		Message: "success",
		Data:    "Create user success",
	})
}
