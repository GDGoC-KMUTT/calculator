package sample

import (
	"github.com/bsthun/gut"
	"github.com/gofiber/fiber/v2"
	"net/url"
	"project1/common"
)

func HandleRedirect(c *fiber.Ctx) error {
	targetUrl, err := url.JoinPath(*common.Config.FrontendPath, "foo")
	if err != nil {
		return gut.Err(false, "Failed to join path", err)
	}

	return c.Redirect(targetUrl)
}
