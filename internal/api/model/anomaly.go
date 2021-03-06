package model

import "github.com/DanielTitkov/anomaly-detection-service/internal/domain"

// Domain models are used for simplicity.
// In production-grade system there would be separate models.

type (
	ListAnomaliesRequest struct {
		Filter domain.FilterAnomaliesArgs
	}
	ListAnomaliesResponse struct {
		Anomalies []*domain.Anomaly
	}
	SetAnomalyStatusRequest struct {
	}
	SetAnomalyStatusResponse struct {
	}
)
