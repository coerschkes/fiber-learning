package internal

import "github.com/gofiber/fiber/v2"

func NewJsonResponse(status uint, msg string, data interface{}) fiber.Map {
	return fiber.Map{"status": status, "message": msg, "data": data}
}
