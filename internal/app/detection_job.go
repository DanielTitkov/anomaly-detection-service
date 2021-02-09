package app

import (
	"fmt"
	"time"

	"github.com/DanielTitkov/anomaly-detection-service/internal/domain"
)

func (a *App) CreateDetectionJob(job *domain.DetectionJob) (*domain.DetectionJob, error) {
	job, err := a.repo.CreateDetectionJob(job)
	if err != nil {
		return nil, err
	}

	if job.Schedule != "" {
		err = a.cron.AddFunc(job.Schedule, func() { a.RunBackgroundJob(*job) })
		if err != nil {
			return nil, err
		}
	}

	return job, nil
}

func (a *App) DeleteDetectionJob(jobID int) error {
	return a.repo.DeleteDetectionJobByID(jobID)
}

func (a *App) ListDetectionJobs(args *domain.FilterDetectionJobsArgs) ([]*domain.DetectionJob, error) {
	return a.repo.FilterDetectionJobs(args)
}

func (a *App) RunDetectionJob(job *domain.DetectionJob) ([]*domain.Anomaly, *domain.DetectionJobInstance, error) {
	startedAt := time.Now()

	dataset, err := a.dataset.Fetch(job)
	if err != nil {
		return []*domain.Anomaly{}, nil, err
	}

	anomalies, err := a.analyzer.FindOutliers(dataset, job)
	if err != nil {
		return []*domain.Anomaly{}, nil, err
	}

	finishedAt := time.Now()
	instance, err := a.repo.CreateDetectionInstanceJob(&domain.DetectionJobInstance{
		DetectionJobID: job.ID,
		StartedAt:      startedAt,
		FinishedAt:     finishedAt,
	})
	if err != nil {
		return []*domain.Anomaly{}, nil, err
	}

	for _, anom := range anomalies { // TODO: add bulk method
		anom.DetectionJobInstanceID = instance.ID
		anom, err = a.repo.CreateAnomaly(anom)
		if err != nil {
			a.logger.Error("failed to create anomaly", err)
		}
	}

	return anomalies, instance, nil
}

func (a *App) RunBackgroundJob(job domain.DetectionJob) {
	jobStr := fmt.Sprintf("%+v", job) // TODO: maybe add domain model method
	a.logger.Info("running detection job", jobStr)
	_, _, err := a.RunDetectionJob(&job)
	if err != nil {
		a.logger.Error("detection job failed", err)
	}
	a.logger.Info("detection job finished", "")
}

func (a *App) GetPendingJobs() ([]*domain.DetectionJob, error) {
	return []*domain.DetectionJob{}, nil
}

func (a *App) ScheduleJobs() error {
	jobs, err := a.repo.FilterDetectionJobs(
		&domain.FilterDetectionJobsArgs{
			Scheduled: true,
		},
	)
	if err != nil {
		return err
	}

	for _, job := range jobs {
		scheduledJob := *job
		err = a.cron.AddFunc(job.Schedule, func() { a.RunBackgroundJob(scheduledJob) })
		if err != nil {
			return err
		}
		a.logger.Info("scheduled job", job.Schedule)
	}

	return nil
}
