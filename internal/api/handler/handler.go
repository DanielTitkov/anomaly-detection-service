package handler

import (
	"github.com/DanielTitkov/anomaly-detection-service/internal/app"
	"github.com/DanielTitkov/anomaly-detection-service/internal/configs"
	"github.com/DanielTitkov/anomaly-detection-service/internal/logger"
	"github.com/labstack/echo"
)

type Handler struct {
	cfg    configs.Config
	logger *logger.Logger
	app    *app.App
}

func NewHandler(
	e *echo.Echo,
	cfg configs.Config,
	logger *logger.Logger,
	app *app.App,
) *Handler {
	h := &Handler{
		cfg:    cfg,
		logger: logger,
		app:    app,
	}
	h.link(e)
	return h
}

func (h *Handler) link(e *echo.Echo) {
	e.POST("/listJobs", h.ListJobsHandler)
	e.POST("/addJob", h.AddJobHandler)
	e.POST("/deleteJob", h.DeleteJobHandler)
	e.POST("/listAnomalies", h.ListAnomaliesHandler)
	e.POST("/setAnomalyStatus", h.SetAnomalyStatusHandler)
}
