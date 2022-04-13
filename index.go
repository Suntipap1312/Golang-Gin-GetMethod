package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/profile", func(c *gin.Context) {
		c.String(http.StatusOK, "My Profile")
	})

	r.GET("/login", func(c *gin.Context) {
		username, password := c.Query("username"), c.Query("password")
		c.JSON(http.StatusOK, gin.H{"status": "OK", "username": username, "password": password})
	})

	r.Run("localhost:80")
}
