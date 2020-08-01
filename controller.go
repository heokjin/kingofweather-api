package main

import (
	// "fmt"
	"net/http"

	"github.com/labstack/echo"
)


func Check(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
