package controller

import (
	usecase "app/usecase/commands"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IMsCardController interface {
	ImportCsv(c echo.Context) error
}

type msCardController struct {
	importMsCardCsvUsecase usecase.IImportMsCardCsvUsecase
}

func NewMsCardController(importMsCardUc usecase.IImportMsCardCsvUsecase) IMsCardController {
	return &msCardController{
		importMsCardCsvUsecase: importMsCardUc,
	}
}

func (mc *msCardController) ImportCsv(c echo.Context) error {
	log.Println("Running MsCardController.ImportCsv")
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

	_, err = mc.importMsCardCsvUsecase.Run(src)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to import cards")
	}

	return c.JSON(http.StatusCreated, "Cards imported successfully")
}
