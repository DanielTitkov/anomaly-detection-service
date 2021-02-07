package entgo

import "github.com/DanielTitkov/anomaly-detection-service/internal/domain"

func (r *EntgoRepository) CreateDetectionJob(j *domain.DetectionJob) (*domain.DetectionJob, error) {
	return &domain.DetectionJob{}, nil
}

func (r *EntgoRepository) DeleteDetectionJobByID(id int) error {
	return nil
}

func (r *EntgoRepository) FilterDetectionJobs(args *domain.FilterDetectionJobsArgs) ([]*domain.DetectionJob, error) {
	return []*domain.DetectionJob{}, nil
}
