package httphandlers

import (
	"net/http"
	"ww2analytic/modules/systems/repositories"

	"github.com/labstack/echo"
)

func GetTotalHistoriesByType(c echo.Context) error {
	ord := c.Param("ord")
	view := "history_type"
	table := "histories"

	result, err := repositories.GetTotalStats("api/v1/histories/bytype:"+ord, ord, view, table)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
