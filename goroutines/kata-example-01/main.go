package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/task", handleTask)

	e.Start(":8080")
}

func handleTask(c echo.Context) error {
	// Simulate a time-consuming operation by introducing a delay
	time.Sleep(10 * time.Second)

	return c.String(http.StatusOK, "Task completed")
}
