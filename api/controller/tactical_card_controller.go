package controller

import (
	usecase "app/usecase/commands"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ITacticalCardController interface {
	ImportCsv(c echo.Context) error
}

type tacticalCardController struct {
	importTacticalCardCsvUsecase usecase.IImportTacticalCardCsvUsecase
}

func NewTacticalCardController(importTacticalCardUc usecase.IImportTacticalCardCsvUsecase) ITacticalCardController {
	return &tacticalCardController{
		importTacticalCardCsvUsecase: importTacticalCardUc,
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
