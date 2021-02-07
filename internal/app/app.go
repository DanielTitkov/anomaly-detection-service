package app

import (
	"github.com/DanielTitkov/anomaly-detection-service/internal/configs"
	"github.com/DanielTitkov/anomaly-detection-service/internal/domain"
	"github.com/DanielTitkov/anomaly-detection-service/internal/logger"
)

type (
	App struct {
		cfg                 configs.Config
		logger              *logger.Logger
		repo                Repository
		datasetService      DatasetFetcher
		notificationService Notifier
	}
	Repository interface {
		// anomalies
		CreateAnomaly(*domain.Anomaly) (*domain.Anomaly, error)
		FilterAnomalies(*domain.FilterAnomaliesArgs) ([]*domain.Anomaly, error)
		SetAnomalyStatus(anomalyID int, processed bool) error

		// detection jobs
		CreateDetectionJob(*domain.DetectionJob) (*domain.DetectionJob, error)
		DeleteDetectionJobByID(int) error
		FilterDetectionJobs(*domain.FilterDetectionJobsArgs) ([]*domain.DetectionJob, error)
		CreateDetectionInstanceJob(*domain.DetectionJobInstance) (*domain.DetectionJobInstance, error)
	}
	DatasetFetcher interface {
		Fetch(*domain.DetectionJob) (*domain.Dataset, error)
	}
	Notifier interface {
		Notify(*domain.Notification) error
	}
	Analyzer interface {
		FindOutliers(*domain.Dataset, *domain.DetectionJob) ([]*domain.Anomaly, error)
	}
)

func NewApp(
	cfg configs.Config,
	logger *logger.Logger,
	repo Repository,
	notificationService Notifier,
) *App {
	return &App{
		cfg:                 cfg,
		logger:              logger,
		repo:                repo,
		notificationService: notificationService,
	}
}
