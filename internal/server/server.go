package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Execute() {
	e := echo.New()
	e.GET("/healthcheck", healthCheck)

	e.Static("/coelho", "output")
	e.Static("/thumbs", "thumbs")

	e.Logger.Fatal(e.Start(":" + "8000"))
}

func healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "WORKING")
}
