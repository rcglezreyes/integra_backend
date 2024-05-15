package utils

import (
	cfg "integra_backend/internal/config"
	resp "integra_backend/internal/entity"
	"integra_backend/internal/message"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleSuccess(c echo.Context, data interface{}) error {
	res := resp.ResponseGeneric{
		Status:  cfg.STATUS_OK,
		Message: message.MsgStatusOkSuccess,
		Data:    data,
	}
	return c.JSON(http.StatusOK, res)
}

func HandleError(c echo.Context, status int, message string) error {
	res := resp.ResponseGeneric{
		Status:  cfg.STATUS_FAILED,
		Message: message,
	}
	return c.JSON(status, res)
}
