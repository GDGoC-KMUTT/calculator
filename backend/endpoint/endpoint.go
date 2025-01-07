package endpoint

import (
	"github.com/gofiber/fiber/v2"
	"project1/endpoint/sample"
)

func Init(router fiber.Router) {
	// api endpoint
	api := router.Group("api")

	// public endpoint
	sampleGroup := api.Group("sample")
	sampleGroup.Get("foo", sample.HandleFoo)
	sampleGroup.Get("redirect", sample.HandleRedirect)
	sampleGroup.Post("calculator/plus", sample.HandleCalculatorPlus)
}
