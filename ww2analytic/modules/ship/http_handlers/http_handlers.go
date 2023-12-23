package httphandlers

import (
	"net/http"
	"ww2analytic/modules/ship/repositories"

	"github.com/labstack/echo"
)

func GetTotalShipClass(c echo.Context) error {
	ord := c.Param("ord")
	view := "class"
	table := "ships"

	result, err := repositories.GetTotalStats("api/v1/ships/role:"+ord, ord, view, table)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalShipCountry(c echo.Context) error {
	ord := c.Param("ord")
	view := "country"
	table := "ships"

	result, err := repositories.GetTotalStats("api/v1/ships/manufacturer:"+ord, ord, view, table)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalShipLaunchYear(c echo.Context) error {
	ord := c.Param("ord")
	view := "launch_year"
	table := "ships"

	result, err := repositories.GetTotalStats("api/v1/ships/country:"+ord, ord, view, table)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
