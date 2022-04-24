package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", holaMundo)
	e.GET("/mexico", holaMexico)
	err := e.Start(":8080")
	if err != nil {
		fmt.Printf("No pude subir el servidor %v", err)
	}
}

func holaMundo(c echo.Context) error {
	return c.String(http.StatusOK, "Hola mundo")
}

func holaMexico(c echo.Context) error {
	return c.String(http.StatusOK, "Hola mexico")
}
