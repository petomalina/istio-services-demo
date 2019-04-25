package metrics

import (
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"sync"
)

// StorageURL is a URL used to connect to the storage service
const StorageURL = "storage"
const ContentTypeJSON = "application/json"

// Service encapsulates the metric controller logic for Observing and Querying data
type Service struct {
	Snapshot      map[string]int64
	SnapshotGuard sync.Mutex
}

// Observe passes a metric information into the storage and updates a local snapshot
// of data used by the querying system
func (s *Service) Observe(c echo.Context) error {
	valuesURL := StorageURL + "?metric=" + c.QueryParam("metric") + "&value=" + c.QueryParam("value")
	_, err := http.Post(valuesURL, ContentTypeJSON, nil)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	s.SnapshotGuard.Lock()
	defer s.SnapshotGuard.Unlock()

	value, err := strconv.Atoi(c.QueryParam("value"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// update the snapshot value
	s.Snapshot[c.QueryParam("metric")] += int64(value)

	return c.String(http.StatusCreated, "OK")
}

type queryResponse struct {
	AggregatedValue int64
}

// Query returns a snapshot of data in case it's present. Otherwise, the snapshot is created
// from events stored in the storage service and synchronized with the local cache
func (s *Service) Query(c echo.Context) error {
	s.SnapshotGuard.Lock()
	defer s.SnapshotGuard.Unlock()

	metric := c.QueryParam("metric")

	if ss, ok := s.Snapshot[metric]; !ok {
		res, err := http.Get(StorageURL + "?metric=" + metric)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		metrics := []int64{}
		err = json.NewDecoder(res.Body).Decode(metrics)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusOK, queryResponse{
		AggregatedValue: s.Snapshot[metric],
	})
}
