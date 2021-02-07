package job

import (
	"time"

	"github.com/DanielTitkov/anomaly-detection-service/internal/app"
	"github.com/DanielTitkov/anomaly-detection-service/internal/configs"
	"github.com/DanielTitkov/anomaly-detection-service/internal/logger"
)

type Service struct {
	cfg    configs.Config
	logger *logger.Logger
	app    *app.App
}

func NewService(
	cfg configs.Config,
	logger *logger.Logger,
	app *app.App,
) *Service {
	return &Service{
		cfg:    cfg,
		logger: logger,
		app:    app,
	}
}

func (s *Service) GatherSystemSummary() {
	go func() {
		for {
			// s.app.CreateSystemSummary()
			// s.logger.Info("calculated system summary", "")
			time.Sleep(time.Duration(s.cfg.Job.SystemSummaryPeriod) * time.Second)
		}
	}()
}
