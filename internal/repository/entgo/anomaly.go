package entgo

import (
	"context"

	"github.com/DanielTitkov/anomaly-detection-service/internal/domain"
)

func (r *EntgoRepository) CreateAnomaly(a *domain.Anomaly) (*domain.Anomaly, error) {
	anom, err := r.client.Anomaly.
		Create().
		SetDetectionJobInstanceID(a.DetectionJobID).
		SetProcessed(a.Processed).
		SetType(a.Type).
		SetValue(a.Value).
		SetPeriodStart(a.PeriodStart).
		SetPeriodEnd(a.PeriodEnd).
		Save(context.TODO())

	if err != nil {
		return nil, err
	}

	a.ID = anom.ID
	return a, nil
}

func (r *EntgoRepository) FilterAnomalies(args *domain.FilterAnomaliesArgs) ([]*domain.Anomaly, error) {
	return []*domain.Anomaly{}, nil
}

func (r *EntgoRepository) SetAnomalyStatus(anomalyID int, processed bool) error {
	_, err := r.client.Anomaly.
		UpdateOneID(anomalyID).
		SetProcessed(processed).
		Save(context.TODO())

	return err
}
