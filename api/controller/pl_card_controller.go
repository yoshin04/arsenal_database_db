package controller

import (
	usecase "app/usecase/commands"
	queryService "app/usecase/queries"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IPlCardController interface {
	FindAll(c echo.Context) error
	FindOneById(c echo.Context) error
	ImportCsv(c echo.Context) error
}

type plCardController struct {
	queryService           queryService.IPlCardQueryService
	importPlCardCsvUsecase usecase.IImportPlCardCsvUsecase
}

func NewPlCardController(qs queryService.IPlCardQueryService, importPlCardCsvUc usecase.IImportPlCardCsvUsecase) IPlCardController {
	return &plCardController{
		queryService:           qs,
		importPlCardCsvUsecase: importPlCardCsvUc,
	}
}

func (pc *plCardController) FindAll(c echo.Context) error {
	plCards, err := pc.queryService.FindAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve cards")
	}
	return c.JSON(http.StatusOK, plCards)
}

func (pc *plCardController) FindOneById(c echo.Context) error {
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

func (pc *plCardController) ImportCsv(c echo.Context) error {
	log.Println("Running PlCardController.ImportCsv")
	file, err := c.FormFile("file")
	if err != nil {
		log.Printf("Error encountered: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to retrieve file")
	}

	src, err := file.Open()
	if err != nil {
		log.Printf("Error encountered: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to open file")
	}
	defer src.Close()

	_, err = pc.importPlCardCsvUsecase.Run(src)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to import cards")
	}

	return c.JSON(http.StatusCreated, "Cards imported successfully")
}
