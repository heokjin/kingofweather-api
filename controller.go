package main

import (
	"github.com/labstack/echo"
	"heokjin/kingofweather-api/model"

	// "fmt"
	"net/http"
)

func Check(c echo.Context) error {
	return c.String(http.StatusOK, "OK")

}

//중기날씨예보
func GetWeatherMidWater(c echo.Context) error {
	tmFc := c.QueryParam("tmFc")
	regId := c.QueryParam("regId")

	res, err := model.GetMidLand(tmFc, regId)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, res)
}
