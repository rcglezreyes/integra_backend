package utils

import (
	"integra_backend/internal/responsegraph"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleSuccess(c echo.Context, data interface{}) error {
	res := responsegraph.ResponseGeneric{
		Status:  "Success",
		Message: "Posts Loaded",
		Data:    data,
	}
	return c.JSON(http.StatusOK, res)
}

func HandleError(c echo.Context, status int, message string) error {
	res := responsegraph.ResponseGeneric{
		Status:  "Failed",
		Message: message,
	}
	return c.JSON(status, res)
}
