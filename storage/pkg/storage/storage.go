package storage

import (
	"github.com/labstack/echo"
	"sync"
)

// Service encapsulates the metric controller logic for Observing and Querying data
type Service struct {
	Metrics      map[string][]int64
	MetricsGuard sync.Mutex
}

// Observe passes a metric information into the storage and updates a local snapshot
// of data used by the querying system
func (s *Service) Create(c echo.Context) error {
	s.MetricsGuard.Lock()
	defer s.MetricsGuard.Unlock()

	return nil
}

// Query returns a snapshot of data in case it's present. Otherwise, the snapshot is created
// from events stored in the storage service and synchronized with the local cache
func (s *Service) Get(c echo.Context) error {
	s.MetricsGuard.Lock()
	defer s.MetricsGuard.Unlock()

	return nil
}
