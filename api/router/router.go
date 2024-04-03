package router

import (
	"app/controller"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewRouter(plc controller.IPlCardController) *echo.Echo {
	e := echo.New()
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "hello",
		})
	})
	e.GET("/pl-cards", plc.FindAll)
	e.GET("/pl-cards/:id", plc.FindOneById)
	e.POST("/pl-cards/import-csv", plc.ImportCsv)

	return e
}
