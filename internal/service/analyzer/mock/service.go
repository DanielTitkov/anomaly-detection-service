package mock

import (
	"github.com/DanielTitkov/anomaly-detection-service/internal/domain"
)

// Service provides mock analyzer
type Service struct{}

// NewService returns mock analyzer service
func NewService() *Service {
	return &Service{}
}

func (s *Service) FindOutliers(dataset *domain.Dataset, job *domain.DetectionJob) ([]*domain.Anomaly, error) {
	var outliers []*domain.Anomaly
	for i, data := range dataset.Data {
		if i == 2 || i == 15 {
			outliers = append(outliers, &domain.Anomaly{
				DetectionJobID: job.ID,
				Type:           domain.AnomalyTypeWarning,
				Value:          data.Value,
				PeriodStart:    data.Timestamp,
				PeriodEnd:      data.Timestamp,
				Processed:      false,
			})
		}
	}

	return outliers, nil
}
