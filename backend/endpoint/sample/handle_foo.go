package sample

import (
	"github.com/bsthun/gut"
	"github.com/gofiber/fiber/v2"
	"project1/type/payload"
)

func HandleFoo(c *fiber.Ctx) error {
	return c.JSON(&payload.Foo{
		Bar: gut.Ptr(123),
	})
}
