package httphandlers

import (
	"net/http"
	"ww2analytic/modules/casualities/repositories"

	"github.com/labstack/echo"
)

func GetTotalCasualitiesByContinent(c echo.Context) error {
	ord := c.Param("ord")
	view := "continent"
	table := "casualities"

	result, err := repositories.GetTotalStats("api/v1/casualities/bycontinent:"+ord, ord, view, table)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
