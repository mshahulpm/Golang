package controllers

import "github.com/gofiber/fiber/v2"

// @Summary Show a string message
// @Description get string message
// @Produce plain
// @Success 200 {string} string "Hello, World!"
// @Router /a [get]
func RoutAController(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

// @Summary Show a string message
// @Description get string message
// @Produce plain
// @Success 200 {string} string "Hello, World!"
// @Router /b [get]
func RoutBController(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

// @Summary Show a string message
// @Description get string message
// @Produce plain
// @Success 200 {string} string "Hello, World!"
// @Router /c [get]
func RoutCController(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

// @Summary Show a string message
// @Description get string message
// @Produce plain
// @Success 200 {string} string "Hello, World!"
// @Router /d [get]
func RoutDController(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
