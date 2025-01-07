package sample

import (
	"github.com/bsthun/gut"
	"github.com/gofiber/fiber/v2"
	"project1/type/payload"
)

func HandleCalculatorPlus(c *fiber.Ctx) error {
	body := new(payload.PlusBody)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	if err := gut.Validate(body); err != nil {
		return err
	}

	result := *body.No1 + *body.No2

	return c.JSON(&payload.CalculateResult{
		Result: &result,
	})
}
