package random

import (
	"math/rand"
	"time"

	"github.com/DanielTitkov/anomaly-detection-service/internal/domain"
)

// Service provides random datasets for given metrics
type Service struct{}

// NewService returns random dataset service
func NewService() *Service {
	return &Service{}
}

func (s *Service) Fetch(job *domain.DetectionJob) (*domain.Dataset, error) {
	rand.Seed(time.Now().UnixNano())

	var data []domain.DataItem
	for i := 0; i <= 30; i++ {
		data = append(data, domain.DataItem{
			Value:     rand.NormFloat64(),
			Timestamp: time.Now().Add(24 * time.Duration(i-30) * time.Hour),
		})
	}

	return &domain.Dataset{
		SiteID:    job.SiteID,
		Metric:    job.Metric,
		Attribute: job.Attribute,
		StartDate: time.Now().Add(-30 * time.Hour * 24), // for the demo purpose
		EndDate:   time.Now(),
		Data:      data,
	}, nil
}
