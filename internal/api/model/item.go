package model

import "github.com/DanielTitkov/anomaly-detection-service/internal/domain"

type (
	CreateItemRequest struct {
		Source string          `json:"source"`
		Task   string          `json:"task"`
		Data   domain.ItemData `json:"data"`
	}
)
