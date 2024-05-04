package router

import (
	"app/controller"
	middleware "app/midleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewRouter(plc controller.IPlCardController, msc controller.IMsCardController, tsc controller.ITacticalCardController, lac controller.ILinkAbilityController, stc controller.ISeriesTitleController, icc controller.IIncludeCodeController, ac controller.IAuthController) *echo.Echo {
	e := echo.New()
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "hello",
		})
	})

	e.POST("/signup", ac.Register)

	e.GET("/pl-cards", plc.FindMany)
	e.GET("/pl-cards/:id", plc.FindOneById)
	e.POST("/pl-cards/import-csv", plc.ImportCsv, middleware.CSVImportAuthMiddleware)

	e.GET("/ms-cards", msc.FindMany)
	e.GET("/ms-cards/:id", msc.FindOneById)
	e.POST("/ms-cards/import-csv", msc.ImportCsv, middleware.CSVImportAuthMiddleware)

	e.GET("/tactical-cards", tsc.FindAll)
	e.POST("/tactical-cards/import-csv", tsc.ImportCsv, middleware.CSVImportAuthMiddleware)

	e.GET("/series-titles", stc.FindMany)

	e.GET("/link-abilities", lac.FindMany)

	e.GET("/include-codes", icc.FindMany)
	return e
}
