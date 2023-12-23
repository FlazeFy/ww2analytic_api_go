package httphandlers

import (
	"net/http"
	"ww2analytic/modules/vehicle/repositories"

	"github.com/labstack/echo"
)

func GetTotalVehicleRole(c echo.Context) error {
	ord := c.Param("ord")
	view := "primary_role"
	table := "vehicles"

	result, err := repositories.GetTotalStats("api/v1/vehicles/role:"+ord, ord, view, table)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalVehicleCountry(c echo.Context) error {
	ord := c.Param("ord")
	view := "country"
	table := "vehicles"

	result, err := repositories.GetTotalStats("api/v1/vehicles/:"+ord, ord, view, table)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
