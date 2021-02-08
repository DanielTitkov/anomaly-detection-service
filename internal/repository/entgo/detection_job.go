package entgo

import (
	"context"

	"github.com/DanielTitkov/anomaly-detection-service/internal/domain"
	"github.com/DanielTitkov/anomaly-detection-service/internal/repository/entgo/ent/detectionjob"
)

func (r *EntgoRepository) CreateDetectionJob(j *domain.DetectionJob) (*domain.DetectionJob, error) {
	return &domain.DetectionJob{}, nil
}

func (r *EntgoRepository) DeleteDetectionJobByID(id int) error {
	return nil
}

func (r *EntgoRepository) GetDetectionJobByID(id int) (*domain.DetectionJob, error) {
	job, err := r.client.DetectionJob.Query().
		Where(detectionjob.IDEQ(id)).
		Only(context.TODO())

	if err != nil {
		return nil, err
	}

	var schedule string
	if job.Schedule != nil {
		schedule = *job.Schedule
	}
	return &domain.DetectionJob{
		ID:       job.ID,
		Schedule: schedule,
		Method:   job.Method,
		SiteID:   job.SiteID,
		// Metric:    job.Metric,
		Attribute:   job.Attribute,
		TimeAgo:     job.TimeAgo,
		TimeStep:    job.TimeStep,
		Description: job.Description,
	}, nil
}

func (r *EntgoRepository) FilterDetectionJobs(args *domain.FilterDetectionJobsArgs) ([]*domain.DetectionJob, error) {
	return []*domain.DetectionJob{}, nil
}
