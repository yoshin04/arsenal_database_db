package controller

import (
	usecase "app/usecase/commands"
	queryService "app/usecase/queries"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IMsCardController interface {
	ImportCsv(c echo.Context) error
	FindMany(c echo.Context) error
}

type msCardController struct {
	importMsCardCsvUsecase usecase.IImportMsCardCsvUsecase
	msCardQueryService     queryService.IMsCardQueryService
}

func NewMsCardController(importMsCardUc usecase.IImportMsCardCsvUsecase, msCardQs queryService.IMsCardQueryService) IMsCardController {
	return &msCardController{
		importMsCardCsvUsecase: importMsCardUc,
		msCardQueryService:     msCardQs,
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

func (mc *msCardController) FindMany(c echo.Context) error {
	log.Println("Running MsCardController.FindMany")
	offsetParam := c.QueryParam("offset")
	limitParam := c.QueryParam("limit")

	offset, _ := strconv.Atoi(offsetParam)
	limit, _ := strconv.Atoi(limitParam)

	msCards, err := mc.msCardQueryService.FindMany(queryService.MsCardFindManyInput{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve cards")
	}

	return c.JSON(http.StatusOK, msCards)
}
