package main

import (
	"errors"
	"fmt"
	"net/http"

	mw "webframeworks/middleware"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Hello(c echo.Context) error {
	switch c.Request().Method {
	case http.MethodGet:
		name := c.QueryParam("name")
		if name != "" {
			c.String(http.StatusOK, fmt.Sprintf("Hello, %s", name))
			return nil
		}
		c.String(http.StatusOK, "Hello, Stranger")
	case http.MethodPost:
		name, _, _ := c.Request().BasicAuth()
		c.String(http.StatusOK, fmt.Sprintf("Hello, %s", name))
		return nil
	}

	return errors.New("bad request method")
}

func main() {
	e := echo.New()

	e.Use(echo.WrapMiddleware(mw.Logging))
	authorized := e.Group("/", echo.WrapMiddleware(mw.Authorization))
	// e.POST("/auth", Hello, echo.WrapMiddleware(mw.Authorization))
	// e.Use(echo.WrapMiddleware(mw.Authorization))

	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/hello", Hello)
	authorized.POST("auth", Hello)

	e.Logger.Fatal(e.Start("localhost:8080"))
}
