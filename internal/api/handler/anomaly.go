package handler

import (
	"net/http"

	"github.com/DanielTitkov/anomaly-detection-service/internal/api/model"
	"github.com/labstack/echo"
)

func (h *Handler) ListAnomaliesHandler(c echo.Context) error {
	request := new(model.ListAnomaliesRequest)
	if err := c.Bind(request); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.OKResponse{
		Status:  "ok",
		Message: "",
	})
}

func (h *Handler) SetAnomalyStatusHandler(c echo.Context) error {
	request := new(model.SetAnomalyStatusRequest)
	if err := c.Bind(request); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.OKResponse{
		Status:  "ok",
		Message: "",
	})
}
