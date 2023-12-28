package httphandlers

import (
	"net/http"
	"ww2analytic/modules/facilities/repositories"

	"github.com/labstack/echo"
)

func GetTotalFacilitiesByType(c echo.Context) error {
	ord := c.Param("ord")
	view := "type"
	table := "facilities"

	result, err := repositories.GetTotalStats("api/v1/facilities/bytype:"+ord, ord, view, table)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalFacilitiesByCountry(c echo.Context) error {
	ord := c.Param("ord")
	view := "country"
	table := "facilities"

	result, err := repositories.GetTotalStats("api/v1/facilities/bycountry:"+ord, ord, view, table)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
