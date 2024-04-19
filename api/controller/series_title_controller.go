package controller

import (
	queryService "app/usecase/queries"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ISeriesTitleController interface {
	FindMany(echo.Context) error
}

type seriesTitleController struct {
	seriesTitleQueryService queryService.ISeriesTitleQueryService
}

func NewSeriesTitleController(seriesTitleQs queryService.ISeriesTitleQueryService) ISeriesTitleController {
	return &seriesTitleController{
		seriesTitleQueryService: seriesTitleQs,
	}
}

func (stc *seriesTitleController) FindMany(c echo.Context) error {
	log.Println("Running SeriesTitleController.FindMany")
	result, err := stc.seriesTitleQueryService.FindAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error fetching series titles")
	}
	return c.JSON(http.StatusOK, result)
}
