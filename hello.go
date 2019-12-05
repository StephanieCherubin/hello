package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func AuthRequired(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello")
}

func main() {
	fmt.Println("Server running")

	e := echo.New()

	//Debug mode
	e.Debug = true

	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}
