package storage

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
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

	metric := c.QueryParam("metric")
	value, _ := strconv.Atoi(c.QueryParam("value"))

	if _, ok := s.Metrics[metric]; !ok {
		s.Metrics[metric] = make([]int64, 10)
	}

	s.Metrics[metric] = append(s.Metrics[metric], int64(value))

	return c.String(http.StatusCreated, "OK")
}

// Query returns a snapshot of data in case it's present. Otherwise, the snapshot is created
// from events stored in the storage service and synchronized with the local cache
func (s *Service) Get(c echo.Context) error {
	s.MetricsGuard.Lock()
	defer s.MetricsGuard.Unlock()

	return c.JSON(http.StatusOK, s.Metrics[c.QueryParam("metric")])
}
