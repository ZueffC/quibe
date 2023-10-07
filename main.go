package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars/v2"
	"github.com/joho/godotenv"
	"github.com/zueffc/quibe/app"
)

func main() {
	//set up the handlebars templates engine
	hb := handlebars.New("./templates", ".html")
	hb.Reload(true)
	hb.AddFunc("bbcode", app.BBCodeFinder)

	//load environment variables
	err := godotenv.Load(".env")

	//set up the web framework
	fiber_app := fiber.New(
		fiber.Config{
			Views: hb,
		},
	)
	fiber_app.Static("/static", "./static")

	if err != nil {
		println(err)
	}

	app.Start(fiber_app)
}
