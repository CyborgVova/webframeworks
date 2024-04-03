package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	c.Header("Content-type", "application/json")
	name, ok := c.GetQuery("name")
	if !ok {
		c.Status(http.StatusBadRequest)
		fmt.Fprintln(c.Writer, "Hello, Stranger")
		return
	}
	fmt.Fprintf(c.Writer, "Hello, %s\n", name)
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Handle("GET", "/hello", Hello)
	router.POST("/auth", Hello)

	fmt.Println("Starting server on localhost:8080 ...")
	log.Fatal(router.Run("localhost:8080"))
}
