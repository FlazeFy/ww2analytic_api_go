package routes

import (
	"net/http"
	airplanehandlers "ww2analytic/modules/airplane/http_handlers"
	fcthandlers "ww2analytic/modules/facilities/http_handlers"
	shiphandlers "ww2analytic/modules/ship/http_handlers"
	strhandlers "ww2analytic/modules/stories/http_handlers"
	vechandlers "ww2analytic/modules/vehicle/http_handlers"

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

	// Vehicles
	e.GET("api/v1/vehicles/byrole/:ord", vechandlers.GetTotalVehicleRole)
	e.GET("api/v1/vehicles/bycountry/:ord", vechandlers.GetTotalVehicleCountry)

	// Facilities
	e.GET("api/v1/facilities/bytype/:ord", fcthandlers.GetTotalFacilitiesByType)
	e.GET("api/v1/facilities/bycountry/:ord", fcthandlers.GetTotalFacilitiesByCountry)

	// Stories
	e.GET("api/v1/stories/bytype/:ord", strhandlers.GetTotalStoriesByType)
	e.GET("api/v1/stories/bylocation/:ord", strhandlers.GetTotalStoriesByLocation)
	e.GET("api/v1/stories/byresult/:ord", strhandlers.GetTotalStoriesByResult)

	// =============== Private routes ===============

	return e
}
