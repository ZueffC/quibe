package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()
	err := godotenv.Load(".env")

	if err != nil {
		println(err)
	}

	app.Get("/", func(ctx *fiber.Ctx) error {
		return nil
	})

	app.Get("/:board", func(ctx *fiber.Ctx) error {
		return nil
	})

	app.Get("/:board/topic/:topic_id", func(ctx *fiber.Ctx) error {
		return nil
	})

	app.Listen(":" + os.Getenv("APP_PORT"))
}
