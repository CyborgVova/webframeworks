package middleware

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"webframeworks/fiber/auth"
)

func Logging(f *fiber.Ctx) error {
	fmt.Printf("method=%s, address=%s:%s, path=%s\n", f.Method(), f.IP(), f.Port(), f.Path())
	return f.Next()
}

func Authorization(f *fiber.Ctx) error {
	username, password, ok := auth.GetBasicAuth(f)
	if !ok {
		f.SendStatus(http.StatusUnauthorized)
	}
	if auth.Auth[username] != password {
		f.SendStatus(http.StatusUnauthorized)
		return nil
	}
	return f.Next()
}
