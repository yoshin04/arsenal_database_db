package controller

import (
	queryService "app/usecase/queries"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IPlCardController interface {
	FindAll(c echo.Context) error
	FindOneById(c echo.Context) error
}

type PlCardController struct {
	queryService queryService.IPlCardQueryService
}

func NewPlCardController(qs queryService.IPlCardQueryService) IPlCardController {
	return &PlCardController{
		queryService: qs,
	}
}

func (pc *PlCardController) FindAll(c echo.Context) error {
	plCards, err := pc.queryService.FindAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve cards")
	}
	return c.JSON(http.StatusOK, plCards)
}

func (pc *PlCardController) FindOneById(c echo.Context) error {
	id := c.Param("id")
	plCard, err := pc.queryService.FindOneById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve card")
	}

	if plCard == nil {
		return echo.NewHTTPError(http.StatusNotFound, "Card not found")
	}

	return c.JSON(http.StatusOK, plCard)
}
