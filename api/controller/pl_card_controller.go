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

type IPlCardController interface {
	FindMany(c echo.Context) error
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

func (pc *plCardController) FindMany(c echo.Context) error {
	log.Println("Running PlCardController.FindMany")
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
	plTypes := c.QueryParams()["plTypes[]"]

	includedCode := utils.OptionalString(c.QueryParam("includedCode"))
	linkAbilityID := utils.OptionalString(c.QueryParam("linkAbilityID"))
	keyword := utils.OptionalString(c.QueryParam("keyword"))

	input := queryService.PlCardFindManyInput{
		Offset:        offset,
		Limit:         limit,
		Costs:         costs,
		Rarities:      rarities,
		PlTypes:       plTypes,
		IncludedCode:  includedCode,
		LinkAbilityID: linkAbilityID,
		Keyword:       keyword,
	}

	plCards, err := pc.queryService.FindMany(input)
	if err != nil {
		log.Printf("Error while retrieving cards: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve cards")
	}
	return c.JSON(http.StatusOK, plCards)
}

func (pc *plCardController) FindOneById(c echo.Context) error {
	log.Println("Running PlCardController.FindOneById")
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
