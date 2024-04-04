package main

import (
	"fmt"
	"log"
	"net/http"
	"webframeworks/fiber/auth"
	mw "webframeworks/fiber/middleware"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/basicauth"
)

func Hello(f *fiber.Ctx) error {
	switch f.Route().Method {
	case http.MethodGet:
		if name := f.Query("name"); name != "" {
			f.SendString(fmt.Sprintf("Hello, %s\n", name))
			return nil
		}
	case http.MethodPost:
		username, _, _ := auth.GetBasicAuth(f)
		f.SendString(fmt.Sprintf("Hello, %s\n", username))
		return nil
	}
	f.SendString("Hello, Stranger\n")
	return nil
}

func main() {
	app := fiber.New()

	app.Use(mw.Logging)

	authorized := app.Group("/")
	authorized.Use(mw.Authorization)
	// authorized.Use(basicauth.New(basicauth.Config{Users: mw.Auth}))
	authorized.Post("/auth", Hello)

	app.Get("/hello", Hello)
	// app.Post("/auth", Authorization, Hello)

	log.Fatal(app.Listen("localhost:8080"))
}
