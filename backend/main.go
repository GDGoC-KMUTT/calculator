package main

import (
	"project1/common/config"
	"project1/common/fiber"
)

func main() {
	config.Init()
	fiber.Init()
}
