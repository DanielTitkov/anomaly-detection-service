package fake

import (
	"fmt"

	"github.com/DanielTitkov/anomaly-detection-service/internal/domain"
)

// Service provides fake notification processing
type Service struct{}

// NewService returns fake notification service
func NewService() *Service {
	return &Service{}
}

// Notify prints notification to stdout
func (s *Service) Notify(n *domain.Notification) error {
	fmt.Println(n)
	return nil
}
