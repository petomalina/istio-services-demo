package metrics

import (
	"github.com/labstack/echo"
	"sync"
)

// Service encapsulates the metric controller logic for Observing and Querying data
type Service struct {
	Snapshot      map[string]int64
	SnapshotGuard sync.Mutex
}

// Observe passes a metric information into the storage and updates a local snapshot
// of data used by the querying system
func (s *Service) Observe(c echo.Context) error {
	s.SnapshotGuard.Lock()
	defer s.SnapshotGuard.Unlock()

	return nil
}

// Query returns a snapshot of data in case it's present. Otherwise, the snapshot is created
// from events stored in the storage service and synchronized with the local cache
func (s *Service) Query(c echo.Context) error {
	s.SnapshotGuard.Lock()
	defer s.SnapshotGuard.Unlock()

	return nil
}
