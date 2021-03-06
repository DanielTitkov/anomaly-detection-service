package handler

import (
	"net/http"

	"github.com/DanielTitkov/anomaly-detection-service/internal/api/model"
	"github.com/labstack/echo"
	"github.com/robfig/cron"
)

func (h *Handler) ListJobsHandler(c echo.Context) error {
	request := new(model.ListJobsRequest)
	if err := c.Bind(request); err != nil {
		return err
	}

	jobs, err := h.app.ListDetectionJobs(&request.Filter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: "failed to filter jobs",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.ListJobsResponse{
		Jobs: jobs,
	})
}

func (h *Handler) AddJobHandler(c echo.Context) error {
	request := new(model.AddJobRequest)
	if err := c.Bind(request); err != nil {
		return err
	}

	if request.Job.Schedule != "" {
		_, err := cron.Parse(request.Job.Schedule)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: "failed to parse cron schedule",
				Error:   err.Error(),
			})
		}
	}

	// TODO: add futher validation
	job, err := h.app.CreateDetectionJob(&request.Job)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: "failed to create job",
			Error:   err.Error(),
		})
	}

	if request.Sync {
		anomalies, instance, err := h.app.RunDetectionJob(job)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: "failed to run sync job",
				Error:   err.Error(),
			})
		}
		return c.JSON(http.StatusOK, model.AddJobResponse{
			Job:      job,
			Instance: instance,
			Result:   anomalies,
		})
	}

	// start async job
	go h.app.RunBackgroundJob(*job)

	return c.JSON(http.StatusOK, model.OKResponse{
		Status:  "ok",
		Message: "job created",
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
