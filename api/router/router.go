package router

import (
	"app/controller"
	middleware "app/midleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewRouter(plc controller.IPlCardController, msc controller.IMsCardController, tsc controller.ITacticalCardController) *echo.Echo {
	e := echo.New()
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "hello",
		})
	})
	e.GET("/pl-cards", plc.FindAll)
	e.GET("/pl-cards/:id", plc.FindOneById)
	e.POST("/pl-cards/import-csv", plc.ImportCsv, middleware.CSVImportAuthMiddleware)

	e.GET("/ms-cards", msc.FindMany)
	e.POST("/ms-cards/import-csv", msc.ImportCsv, middleware.CSVImportAuthMiddleware)

	e.POST("/tactical-cards/import-csv", tsc.ImportCsv, middleware.CSVImportAuthMiddleware)

	return e
}
