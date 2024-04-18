package controller

import (
	usecase "app/usecase/commands"
	queryService "app/usecase/queries"
	"app/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IMsCardController interface {
	ImportCsv(c echo.Context) error
	FindMany(c echo.Context) error
	FindOneById(c echo.Context) error
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
	queryParams := c.QueryParams()
	for key, value := range queryParams {
		log.Printf("Query Param [%s]: %v\n", key, value)
	}
	offsetParam := c.QueryParam("offset")
	limitParam := c.QueryParam("limit")

	offset, _ := strconv.Atoi(offsetParam)
	limit, _ := strconv.Atoi(limitParam)

	costParams := c.QueryParams()["costs[]"]
	costs := make([]uint8, 0, len(costParams))
	for _, cp := range costParams {
		if cost, err := strconv.Atoi(cp); err == nil {
			costs = append(costs, uint8(cost))
		} else {
			log.Printf("Error parsing cost '%s': %v", cp, err)
		}
	}

	rarities := c.QueryParams()["rarities[]"]
	msTypes := c.QueryParams()["msTypes[]"]
	groundSuitabilities := c.QueryParams()["groundSuitabilities[]"]
	spaceSuitabilities := c.QueryParams()["spaceSuitabilities[]"]
	desertSuitabilities := c.QueryParams()["desertSuitabilities[]"]
	underwaterSuitabilities := c.QueryParams()["underwaterSuitabilities[]"]

	includedCode := utils.OptionalString(c.QueryParam("includedCode"))
	linkAbilityID := utils.OptionalString(c.QueryParam("linkAbilityID"))
	keyword := utils.OptionalString(c.QueryParam("keyword"))

	input := queryService.MsCardFindManyInput{
		Offset:                  offset,
		Limit:                   limit,
		Costs:                   costs,
		Rarities:                rarities,
		MsTypes:                 msTypes,
		GroundSuitabilities:     groundSuitabilities,
		SpaceSuitabilities:      spaceSuitabilities,
		DesertSuitabilities:     desertSuitabilities,
		UnderwaterSuitabilities: underwaterSuitabilities,
		IncludedCode:            includedCode,
		LinkAbilityID:           linkAbilityID,
		Keyword:                 keyword,
	}

	msCards, err := mc.msCardQueryService.FindMany(input)
	if err != nil {
		log.Printf("Error while retrieving cards: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve cards")
	}

	return c.JSON(http.StatusOK, msCards)
}

func (mc *msCardController) FindOneById(c echo.Context) error {
	log.Println("Running MsCardController.FindOneById")
	id := c.Param("id")
	msCard, err := mc.msCardQueryService.FindOneById(id)
	if err != nil {
		log.Printf("Error while retrieving card: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve card")
	}

	if msCard == nil {
		return echo.NewHTTPError(http.StatusNotFound, "Card not found")
	}

	return c.JSON(http.StatusOK, msCard)
}
