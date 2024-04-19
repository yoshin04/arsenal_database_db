package controller

import (
	queryService "app/usecase/queries"
	"log"

	"github.com/labstack/echo/v4"
)

type IIncludeCodeController interface {
	FindMany(echo.Context) error
}

type includeCodeController struct {
	includeCodeQueryService queryService.IIncludeCodeQueryService
}

func NewIncludeCodeController(includeCodeQs queryService.IIncludeCodeQueryService) IIncludeCodeController {
	return &includeCodeController{
		includeCodeQueryService: includeCodeQs,
	}
}

func (icc *includeCodeController) FindMany(c echo.Context) error {
	log.Println("Running IncludeCodeController.FindMany")
	result, err := icc.includeCodeQueryService.FindMany()
	if err != nil {
		return echo.NewHTTPError(500, "Error fetching include codes")
	}
	return c.JSON(200, result)
}
