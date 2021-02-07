package entgo

import "github.com/DanielTitkov/anomaly-detection-service/internal/domain"

func (r *EntgoRepository) CreateAnomaly(a *domain.Anomaly) (*domain.Anomaly, error) {
	return &domain.Anomaly{}, nil
}

func (r *EntgoRepository) FilterAnomalies(args *domain.FilterAnomaliesArgs) ([]*domain.Anomaly, error) {
	return []*domain.Anomaly{}, nil
}

func (r *EntgoRepository) SetAnomalyStatus(anomalyID int, processed bool, status string) error {
	return nil
}
