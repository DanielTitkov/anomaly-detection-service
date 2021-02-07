package handler

import (
	"net/http"

	"github.com/DanielTitkov/anomaly-detection-service/internal/api/model"
	"github.com/labstack/echo"
)

func (h *Handler) ListJobsHandler(c echo.Context) error {
	request := new(model.ListJobsRequest)
	if err := c.Bind(request); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.OKResponse{
		Status:  "ok",
		Message: "",
	})
}

func (h *Handler) AddJobHandler(c echo.Context) error {
	request := new(model.AddJobRequest)
	if err := c.Bind(request); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.OKResponse{
		Status:  "ok",
		Message: "",
	})
}

func (h *Handler) DeleteJobHandler(c echo.Context) error {
	request := new(model.DeleteJobRequest)
	if err := c.Bind(request); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.OKResponse{
		Status:  "ok",
		Message: "",
	})
}
