package main

import (
	"fmt"
	"log"
	"os"

	"github.com/anthdm/bs/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env.local"); err != nil {
		log.Fatal(err)
	}

	app := fiber.New(fiber.Config{
		ErrorHandler:          handlers.ErrorHandler,
		DisableStartupMessage: true,
		PassLocalsToViews:     true,
		Views:                 createEngine(),
	})

	initRoutes(app)

	listenAddr := os.Getenv("HTTP_LISTEN_ADDR")

	fmt.Printf("app listening on: http://127.0.0.1:%s\n", listenAddr)

	log.Fatal(app.Listen(listenAddr))
}

func initRoutes(app *fiber.App) {
	app.Static("/public", "./public")
	app.Get("/", handlers.HandleHome)
}
