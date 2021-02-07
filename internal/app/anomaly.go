package app

import "github.com/DanielTitkov/anomaly-detection-service/internal/domain"

func (a *App) CreateAnomaly(*domain.Anomaly) error {
	return nil
}

func (a *App) ListAnomalies(*domain.FilterAnomaliesArgs) ([]*domain.Anomaly, error) {
	return []*domain.Anomaly{}, nil
}

func (a *App) SetAnomalyStatus(anomalyID int, processed bool, status string) error {
	return nil
}
