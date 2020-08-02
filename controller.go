package main

import (
	"github.com/labstack/echo"
	// "fmt"
	"net/http"
)


func Check(c echo.Context) error {


	return c.String(http.StatusOK, "OK")

}

//중기날씨예보
func GetWeatherMidWater(c echo.Context) error {
	//tmFc := c.QueryParam("tmFc")
	//regId := c.QueryParam("regId")




	return c.String(http.StatusOK, "OK")
}
