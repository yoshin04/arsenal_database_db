package controller

import (
	queryService "app/usecase/queries"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ILinkAbilityController interface {
	FindMany(c echo.Context) error
}

type linkAbilityController struct {
	linkAbilityQueryService queryService.ILinkAbilityQueryService
}

func NewLinkAbilityController(linkAbilityQs queryService.ILinkAbilityQueryService) ILinkAbilityController {
	return &linkAbilityController{
		linkAbilityQueryService: linkAbilityQs,
	}
}

func (lac *linkAbilityController) FindMany(c echo.Context) error {
	log.Println("Running LinkAbilityController.FindMany")
	result, err := lac.linkAbilityQueryService.FindMany()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve link abilities")
	}
	return c.JSON(http.StatusOK, result)
}
