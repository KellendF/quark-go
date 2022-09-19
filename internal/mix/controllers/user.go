package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type User struct{}

// Index
func (p *User) Index(c *fiber.Ctx) error {

	return c.SendString("Auth Page")
}
