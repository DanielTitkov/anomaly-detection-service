package app

import (
	"github.com/DanielTitkov/anomaly-detection-service/internal/configs"
	"github.com/DanielTitkov/anomaly-detection-service/internal/domain"
	"github.com/DanielTitkov/anomaly-detection-service/internal/logger"
	"github.com/robfig/cron"
)

type (
	// App combines services and holds business logic
	App struct {
		cfg          configs.Config
		logger       *logger.Logger
		repo         Repository
		dataset      DatasetFetcher
		notification Notifier
		analyzer     Analyzer
		cron         *cron.Cron
	}
	// Repository stores data
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
	// DatasetFetcher fetches dataset for a given job
	DatasetFetcher interface {
		Fetch(*domain.DetectionJob) (*domain.Dataset, error)
	}
	// Notifier provides notification service
	Notifier interface {
		Notify(*domain.Notification) error
	}
	// Analyzer provides outlier detection service
	Analyzer interface {
		FindOutliers(*domain.Dataset, *domain.DetectionJob) ([]*domain.Anomaly, error)
	}
)

func NewApp(
	cfg configs.Config,
	logger *logger.Logger,
	repo Repository,
	dataset DatasetFetcher,
	notification Notifier,
	analyzer Analyzer,
) *App {
	c := cron.New()
	c.Start()

	return &App{
		cfg:          cfg,
		logger:       logger,
		repo:         repo,
		dataset:      dataset,
		notification: notification,
		analyzer:     analyzer,
		cron:         c,
	}
}
