package fiber

import (
	"fmt"
	"github.com/bsthun/gut"
	"github.com/gofiber/fiber/v2"
	"project1/type/response"
)

func HandleNotFound(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
		Success: gut.Ptr(false),
		Message: gut.Ptr(fmt.Sprintf("%s %s not found", c.Method(), c.Path())),
		Error:   gut.Ptr("NOT_FOUND"),
	})
}
