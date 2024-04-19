package controller

import (
	usecase "app/usecase/commands"
	queryService "app/usecase/queries"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ITacticalCardController interface {
	ImportCsv(c echo.Context) error
	FindAll(c echo.Context) error
}

type tacticalCardController struct {
	importTacticalCardCsvUsecase usecase.IImportTacticalCardCsvUsecase
	tacticalCardQueryService     queryService.ITacticalCardQueryService
}

func NewTacticalCardController(importTacticalCardUc usecase.IImportTacticalCardCsvUsecase, tCardQs queryService.ITacticalCardQueryService) ITacticalCardController {
	return &tacticalCardController{
		importTacticalCardCsvUsecase: importTacticalCardUc,
		tacticalCardQueryService:     tCardQs,
	}
}

func (tc *tacticalCardController) ImportCsv(c echo.Context) error {
	log.Println("Running TacticalCardController.ImportCsv")
	file, err := c.FormFile("file")
	if err != nil {
		log.Printf("Error encountered: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve file")
	}

	src, err := file.Open()
	if err != nil {
		log.Printf("Error encountered: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to open file")
	}
	defer src.Close()

	_, err = tc.importTacticalCardCsvUsecase.Run(src)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to import cards")
	}

	return c.JSON(http.StatusCreated, "Cards imported successfully")
}

func (tc *tacticalCardController) FindAll(c echo.Context) error {
	log.Println("Running TacticalCardController.FindAll")
	queryParams := c.QueryParams()
	for key, value := range queryParams {
		log.Printf("Query Param [%s]: %v\n", key, value)
	}

	costParams := c.QueryParams()["costs[]"]
	costs := make([]uint8, len(costParams))
	for _, cp := range costParams {
		if cost, err := strconv.Atoi(cp); err == nil {
			costs = append(costs, uint8(cost))
		} else {
			log.Printf("Error parsing cost: %v", err)
		}
	}

	input := queryService.TacticalCardFindManyInput{
		Costs: costs,
	}

	tacticalCards, err := tc.tacticalCardQueryService.FindAll(input)
	if err != nil {
		log.Printf("Error encountered: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve cards")
	}

	return c.JSON(http.StatusOK, tacticalCards)
}
