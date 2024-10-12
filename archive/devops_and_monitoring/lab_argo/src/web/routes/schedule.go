package routes

import (
	"github.com/gofiber/fiber/v2"
	"lab-argo-app/pkg/schedule"
	"lab-argo-app/web/handlers"
)

func ScheduleRouter(app fiber.Router, service schedule.Service) {
	app.Get("/", handlers.GetSchedule(service))
	app.Get("/lessons/:day", handlers.GetDay(service))
}
