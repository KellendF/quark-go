package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type File struct{}

func (p *File) Show(c *fiber.Ctx) error {

	return c.JSON("")
}
