package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"

	mw "webframeworks/middleware"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/basicauth"
)

func Logging(f *fiber.Ctx) error {
	fmt.Printf("method=%s, address=%s:%s, path=%s\n", f.Method(), f.IP(), f.Port(), f.Path())
	return f.Next()
}

func GetBasicAuth(f *fiber.Ctx) (string, string, bool) {
	auth := f.Get("Authorization")
	if len(auth) <= 6 {
		return "", "", false
	}
	b, err := base64.StdEncoding.DecodeString(auth[6:])
	if err != nil {
		return "", "", false
	}

	cred := string(b)
	index := strings.Index(cred, ":")
	return cred[:index], cred[index+1:], true
}

func Authorization(f *fiber.Ctx) error {
	username, password, ok := GetBasicAuth(f)
	if !ok {
		f.SendStatus(http.StatusUnauthorized)
	}
	if mw.Auth[username] != password {
		f.SendStatus(http.StatusUnauthorized)
		return nil
	}
	return f.Next()
}

func Hello(f *fiber.Ctx) error {
	switch f.Route().Method {
	case http.MethodGet:
		if name := f.Query("name"); name != "" {
			f.SendString(fmt.Sprintf("Hello, %s\n", name))
			return nil
		}
	case http.MethodPost:
		username, _, _ := GetBasicAuth(f)
		f.SendString(fmt.Sprintf("Hello, %s\n", username))
		return nil
	}
	f.SendString("Hello, Stranger\n")
	return nil
}

func main() {
	app := fiber.New()

	app.Use(Logging)

	authorized := app.Group("/")
	authorized.Use(Authorization)
	// authorized.Use(basicauth.New(basicauth.Config{Users: mw.Auth}))
	authorized.Post("/auth", Hello)

	app.Get("/hello", Hello)
	// app.Post("/auth", Authorization, Hello)

	log.Fatal(app.Listen("localhost:8080"))
}
