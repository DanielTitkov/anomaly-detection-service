package app

import "github.com/DanielTitkov/anomaly-detection-service/internal/domain"

func (a *App) CreateDetectionJob(*domain.DetectionJob) error {
	return nil
}

func (a *App) DeleteDetectionJob(jobID int) error {
	return nil
}

func (a *App) ListDetectionJobs(args *domain.FilterDetectionJobsArgs) ([]*domain.DetectionJob, error) {
	return []*domain.DetectionJob{}, nil
}

func (a *App) RunDetectionJob(jobID int) error {
	return nil
}

func (a *App) GetPendingJobs() ([]*domain.DetectionJob, error) {
	return []*domain.DetectionJob{}, nil
}

func (a *App) RunJobs() error {
	return nil
}
