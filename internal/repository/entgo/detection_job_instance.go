package entgo

import "github.com/DanielTitkov/anomaly-detection-service/internal/domain"

func (r *EntgoRepository) CreateDetectionInstanceJob(i *domain.DetectionJobInstance) (*domain.DetectionJobInstance, error) {
	return &domain.DetectionJobInstance{}, nil
}
