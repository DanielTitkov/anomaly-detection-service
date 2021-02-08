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

	v1 := e.Group("/api/v1")
	v1.POST("/listJobs", h.ListJobsHandler)
	v1.POST("/addJob", h.AddJobHandler)
	v1.POST("/deleteJob", h.DeleteJobHandler)
	v1.POST("/listAnomalies", h.ListAnomaliesHandler)
	v1.POST("/setAnomalyStatus", h.SetAnomalyStatusHandler)
}
