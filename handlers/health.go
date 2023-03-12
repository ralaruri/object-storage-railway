package handlers

import "github.com/gofiber/fiber/v2"

// @Summary Show the status of server.
// @Description get the status of server.
// @Tags health
// @Accept */*
// @Produce plain
// @Success 200 "OK"
// @Router /api/v1/health [get]
func HandlerHealthCheck(c *fiber.Ctx) error {
	return c.SendString("Everything is looking healthy, lets look from cheese!")
}
