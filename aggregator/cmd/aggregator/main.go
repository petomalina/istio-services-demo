package main

import (
	"github.com/labstack/echo"
	"github.com/petomalina/istio-services-demo/aggregator/pkg/metrics"
	log "github.com/sirupsen/logrus"
)

func main() {
	e := echo.New()

	svc := &metrics.Service{
		Snapshot: make(map[string]int64),
	}

	e.POST("/observe", svc.Observe)
	e.GET("/query", svc.Query)

	log.Panic(e.Start(":80"))
}
