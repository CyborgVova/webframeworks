package main

import (
	"fmt"
	"log"

	fiber "github.com/gofiber/fiber/v2"
)

func Hello(f *fiber.Ctx) error {
	f.SendString("Hello, World !!!")
	return nil
}

func main() {
	app := fiber.New(fiber.Config{
		ServerHeader: "Fiber Framework",
	})

	app.Get("/hello", Hello)

	fmt.Println("Starting server on localhost:8080 ...")
	log.Fatal(app.Listen("localhost:8080"))
}
