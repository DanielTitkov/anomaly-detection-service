package mock

import (
	"fmt"

	"github.com/DanielTitkov/anomaly-detection-service/internal/domain"
)

// Service provides mock notification processing
type Service struct{}

// NewService returns mock notification service
func NewService() *Service {
	return &Service{}
}

// Notify prints notification to stdout
func (s *Service) Notify(n *domain.Notification) error {
	fmt.Println(n)
	return nil
}
