package middleware

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func CSVImportAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		clientKey := c.Request().Header.Get("x-csv-import-key")

		serverKey := os.Getenv("CSV_IMPORT_KEY")

		if clientKey != serverKey {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid authentication key")
		}

		return next(c)
	}
}
