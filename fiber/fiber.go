package main

import (
	"log"

	fiber "github.com/gofiber/fiber/v2"
)

func Hello(f *fiber.Ctx) error {
	f.SendString("Hello, World !!!")
	return nil
}

func main() {
	app := fiber.New()

	app.Get("/hello", Hello)

	log.Fatal(app.Listen("localhost:8080"))
}
