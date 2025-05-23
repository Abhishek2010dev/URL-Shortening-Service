package main

import (
	"log"

	"github.com/Abhishek2010dev/URL-Shortening-Service/database"
	"github.com/Abhishek2010dev/URL-Shortening-Service/handler"
	"github.com/Abhishek2010dev/URL-Shortening-Service/repository"
	"github.com/Abhishek2010dev/URL-Shortening-Service/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: utils.ErrorHandler,
	})
	app.Use(logger.New())

	shortenRepo := repository.NewShorten(database.Connect())
	shortenHandler := handler.NewShorten(shortenRepo)

	app.Post("/shorten", shortenHandler.Create)
	app.Get("/shorten/:short_code", shortenHandler.GetByShortCode)
	app.Get("/shorten/:short_code/stats", shortenHandler.GetURLStatistics)
	app.Delete("/shorten/:short_code", shortenHandler.Delete)
	app.Patch("/shorten/:short_code", shortenHandler.Update)
	app.Get("/:short_code", shortenHandler.Redirect)

	log.Fatal(app.Listen(":3000"))
}
