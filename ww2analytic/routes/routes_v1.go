package routes

import (
	"net/http"
	airplanehandlers "ww2analytic/modules/airplane/http_handlers"

	"github.com/labstack/echo"
)

func InitV1() *echo.Echo {
	e := echo.New()

	e.GET("api/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to WW2 Analytic")
	})

	// =============== Public routes ===============

	// Dictionary
	e.GET("api/v1/airplane/role", airplanehandlers.GetTotalAirplaneRole)
	e.GET("api/v1/airplane/manufacturer", airplanehandlers.GetTotalAirplaneManufacturer)
	e.GET("api/v1/airplane/country", airplanehandlers.GetTotalAirplaneCountry)

	// =============== Private routes ===============

	return e
}
