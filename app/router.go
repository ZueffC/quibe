package app

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Route(fiber_app *fiber.App, db *gorm.DB) {
	fiber_app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Render("index", fiber.Map{"PageName": "Index"}, "layouts/default")
	})

	fiber_app.Get("/:board_name", func(ctx *fiber.Ctx) error {
		board_name := ctx.Params("board_name")
		messages := GetAllMessFromBoard(board_name, db)

		return ctx.Render("board", fiber.Map{
			"PageName": "Board",
			"BN":       board_name,
			"Messages": messages,
		}, "layouts/default")
	})

	fiber_app.Get("/:board_name/topic/:topic_id", func(ctx *fiber.Ctx) error {
		board_name, topic_id := ctx.Params("board_name"), ctx.Params("topic_id")
		b_id := GetBoardByName(board_name, db).ID
		t_id, _ := strconv.Atoi(topic_id)

		topic_start := GetMessById(t_id, db)
		msgs := GetMessByBoardAndTopic(b_id, t_id, db)

		return ctx.Render("reply", fiber.Map{
			"Messages": msgs,
			"TStart":   topic_start,
		}, "layouts/default")
	})

	fiber_app.Post("/:board_name", func(ctx *fiber.Ctx) error {
		board := ctx.Params("board_name")
		data := new(MessageForm)
		data.IsReply = false
		data.Board = board

		file, ferr := ctx.FormFile("file_input")

		if ferr == nil {
			path := fmt.Sprintf("./static/imgs/%s", file.Filename)
			data.File = path
			ctx.SaveFile(file, path)
		}

		if err := ctx.BodyParser(data); err != nil {
			return err
		}

		if err := AddNewMessage(data, 0, db); err != nil {
			return err
		}

		return ctx.Redirect("/" + board)
	})

	fiber_app.Post("/:board_name/topic/:topic_id", func(ctx *fiber.Ctx) error {
		board_name, topic_id := ctx.Params("board_name"), ctx.Params("topic_id")
		t_id, _ := strconv.Atoi(topic_id)

		data := new(MessageForm)
		data.IsReply = true
		data.Board = board_name

		file, ferr := ctx.FormFile("file_input")

		if ferr == nil {
			path := fmt.Sprintf("./static/imgs/%s", file.Filename)
			data.File = path
			ctx.SaveFile(file, path)
		}

		if err := ctx.BodyParser(data); err != nil {
			return err
		}

		if err := AddNewMessage(data, t_id, db); err != nil {
			return err
		}

		return ctx.Redirect("/")
	})

}
