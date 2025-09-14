package main

import (
	"example/server/config"
	"example/server/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Static("/uploads", "./uploads")

	config.InitDB()

	router.InitRouter(app)

	config.FiberRun(app, "3000")
}
