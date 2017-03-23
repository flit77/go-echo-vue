package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

func main() {
	// Create a new instance of Echo
	e := echo.New()

	e.GET("/tasks", func(c echo.Context) error { return c.JSON(200, "GET Tasks") })
	e.PUT("/tasks", func(c echo.Context) error { return c.JSON(200, "PUT Tasks") })
	e.DELETE("/tasks/:id", func(c echo.Context) error { return c.JSON(200, "DELETE Task "+c.Param("id")) })

	// Start as a web server
	e.Run(standard.New(":8000"))
}
