package app

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func Start(fiber_app *fiber.App) error {
	migrator := Migrator{
		Models: []interface{}{&Boards{}, &Messages{}},
	}

	db := migrator.ConnectToDB().Migrate().db

	fiber_app.Use(func(ctx *fiber.Ctx) error {
		boards := GetAllBoards(db)

		ctx.Bind(fiber.Map{
			"SiteName": os.Getenv("APP_NAME"),
			"boards":   boards,
		})

		return ctx.Next()
	})

	Route(fiber_app, db)
	return fiber_app.Listen(":" + os.Getenv("APP_PORT"))
}
