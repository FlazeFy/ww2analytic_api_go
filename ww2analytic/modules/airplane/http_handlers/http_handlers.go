package httphandlers

import (
	"net/http"
	"ww2analytic/modules/airplane/repositories"

	"github.com/labstack/echo"
)

func GetTotalAirplaneRole(c echo.Context) error {
	ord := c.Param("ord")
	view := "aircraft_role"
	table := "aircraft"

	result, err := repositories.GetTotalStats("api/v1/airplane/role:"+ord, ord, view, table)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalAirplaneManufacturer(c echo.Context) error {
	ord := c.Param("ord")
	view := "aircraft_manufacturer"
	table := "aircraft"

	result, err := repositories.GetTotalStats("api/v1/airplane/manufacturer:"+ord, ord, view, table)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalAirplaneCountry(c echo.Context) error {
	ord := c.Param("ord")
	view := "country_code"
	table := "aircraft"

	result, err := repositories.GetTotalStats("api/v1/airplane/country:"+ord, ord, view, table)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
