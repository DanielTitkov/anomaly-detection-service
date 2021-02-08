package app

import (
	"time"

	"github.com/DanielTitkov/anomaly-detection-service/internal/domain"
)

func (a *App) CreateDetectionJob(job *domain.DetectionJob) error {
	_, err := a.repo.CreateDetectionJob(job)
	return err
}

func (a *App) DeleteDetectionJob(jobID int) error {
	return a.repo.DeleteDetectionJobByID(jobID)
}

func (a *App) ListDetectionJobs(args *domain.FilterDetectionJobsArgs) ([]*domain.DetectionJob, error) {
	return a.repo.FilterDetectionJobs(args)
}

func (a *App) RunDetectionJob(job *domain.DetectionJob) error {
	startedAt := time.Now()

	dataset, err := a.dataset.Fetch(job)
	if err != nil {
		return err
	}

	anomalies, err := a.analyzer.FindOutliers(dataset, job)
	if err != nil {
		return err
	}

	finishedAt := time.Now()
	_, err = a.repo.CreateDetectionInstanceJob(&domain.DetectionJobInstance{
		DetectionJobID: job.ID,
		StartedAt:      startedAt,
		FinishedAt:     finishedAt,
	})
	if err != nil {
		return err
	}

	for _, anom := range anomalies { // TODO: add bulk method
		_, err := a.repo.CreateAnomaly(anom)
		if err != nil {
			a.logger.Error("failed to create anomaly", err)
		}
	}

	return nil
}

func (a *App) GetPendingJobs() ([]*domain.DetectionJob, error) {
	return []*domain.DetectionJob{}, nil
}

func (a *App) RunJobs() error {
	return nil
}
