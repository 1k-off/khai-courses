package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"lab-argo-app/pkg/entities"
	"lab-argo-app/pkg/schedule"
	"lab-argo-app/web/presenter"
	"strings"
)

var (
	errUnsuppDay = fmt.Errorf("unsupported day")
)

func GetSchedule(service schedule.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		lessons, err := service.GetSchedule()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(lessons)
	}
}

func GetDay(service schedule.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		day := strings.ToLower(c.Params("day"))
		ok := entities.ValidateDay(day)
		if !ok {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.ErrorResponse(errUnsuppDay))
		}
		lessons, err := service.GetDay(day)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.ErrorResponse(err))
		}
		return c.JSON(presenter.SuccessResponse(lessons))
	}
}
