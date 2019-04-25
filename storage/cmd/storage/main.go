package main

import (
	"github.com/labstack/echo"
	"github.com/petomalina/istio-services-demo/storage/pkg/storage"
	log "github.com/sirupsen/logrus"
)

func main() {
	e := echo.New()

	svc := &storage.Service{
		Metrics: make(map[string][]int64),
	}

	e.POST("/", svc.Create)
	e.GET("/", svc.Get)

	log.Panic(e.Start(":80"))
}
