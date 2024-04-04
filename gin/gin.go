package main

import (
	"fmt"
	"log"
	"net/http"

	"webframeworks/storage"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	c.Header("Content-type", "application/json")
	switch c.Request.Method {
	case http.MethodGet:
		name, ok := c.GetQuery("name")
		if !ok {
			c.Status(http.StatusBadRequest)
			fmt.Fprintln(c.Writer, "Hello, Stranger")
			return
		}
		fmt.Fprintf(c.Writer, "Hello, %s\n", name)
	case http.MethodPost:
		fmt.Fprintf(c.Writer, "Hello, %s\n", c.MustGet(gin.AuthUserKey).(string))
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	// router := gin.New()

	router := gin.Default()
	router.POST("/auth", gin.BasicAuth(storage.Auth), Hello)

	router.Handle("GET", "/hello", Hello)

	// authorized := router.Group("/")
	// authorized.Use(gin.BasicAuth(mw.Auth))
	// authorized.Use(gin.Logger())
	// authorized.Use(gin.Recovery())

	// authorized.POST("/auth", Hello)

	fmt.Println("Starting server on localhost:8080 ...")
	log.Fatal(router.Run("localhost:8080"))
}
