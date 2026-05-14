package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	router := echo.New()

	router.Use(middleware.Logger())

	router.GET("/health-check", healthCheck)

	port := getEnv("MARADONA_HTTP_PORT", "8080")

	if err := router.Start(fmt.Sprintf(":%s", port)); err != nil {
		panic(err)
	}
}

func healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"message": "everything is good",
	})
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return fallback
}
