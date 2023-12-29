package httphandlers

import (
	"net/http"
	"ww2analytic/modules/stories/repositories"

	"github.com/labstack/echo"
)

func GetTotalStoriesByType(c echo.Context) error {
	ord := c.Param("ord")
	view := "story_type"
	table := "stories"

	result, err := repositories.GetTotalStats("api/v1/stories/bytype:"+ord, ord, view, table)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalStoriesByLocation(c echo.Context) error {
	ord := c.Param("ord")
	view := "story_location"
	table := "stories"

	result, err := repositories.GetTotalStats("api/v1/stories/bylocation:"+ord, ord, view, table)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalStoriesByResult(c echo.Context) error {
	ord := c.Param("ord")
	view := "story_result"
	table := "stories"

	result, err := repositories.GetTotalStats("api/v1/stories/byresult:"+ord, ord, view, table)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
