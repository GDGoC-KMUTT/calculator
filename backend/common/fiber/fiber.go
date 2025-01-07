package fiber

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"project1/common"
	"project1/common/fiber/middleware"
	"project1/endpoint"
	"project1/type/response"
)

func Init() {
	// initialize fiber instance
	app := fiber.New(fiber.Config{
		AppName:       "Project 1",
		ErrorHandler:  HandleError,
		Prefork:       false,
		StrictRouting: true,
		Network:       fiber.NetworkTCP,
	})

	// register middleware
	app.Use(middleware.Recover())
	app.Use(middleware.Cors())

	// register root endpoint
	app.All("/", func(c *fiber.Ctx) error {
		return c.JSON(response.Success("Pre-hackathon API ROOT"))
	})

	// register endpoint
	endpoint.Init(app)

	// register not found endpoint
	app.Use(HandleNotFound)

	// Startup
	err := app.Listen(*common.Config.Address)
	if err != nil {
		log.Fatal("unable to start fiber instance", err)
	}
}
