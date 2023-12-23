package routes

import (
	"net/http"
	airplanehandlers "ww2analytic/modules/airplane/http_handlers"
	shiphandlers "ww2analytic/modules/ship/http_handlers"

	"github.com/labstack/echo"
)

func InitV1() *echo.Echo {
	e := echo.New()

	e.GET("api/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to WW2 Analytic")
	})

	// =============== Public routes ===============

	// Airplane
	e.GET("api/v1/aircraft/byrole/:ord", airplanehandlers.GetTotalAirplaneRole)
	e.GET("api/v1/aircraft/bymanufacturer/:ord", airplanehandlers.GetTotalAirplaneManufacturer)
	e.GET("api/v1/aircraft/bycountry/:ord", airplanehandlers.GetTotalAirplaneCountry)

	// Ships
	e.GET("api/v1/ships/byclass/:ord", shiphandlers.GetTotalShipClass)
	e.GET("api/v1/ships/bycountry/:ord", shiphandlers.GetTotalShipCountry)
	e.GET("api/v1/ships/bylaunchyear/:ord", shiphandlers.GetTotalShipLaunchYear)

	// =============== Private routes ===============

	return e
}
