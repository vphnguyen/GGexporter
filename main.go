package main

import (
    "GGexporter/services"
    "GGexporter/handler"
    "github.com/prometheus/client_golang/prometheus"
)
var exporterport int

func main() {
    registry := prometheus.NewRegistry() 
    registry.MustRegister(services.NewGoldenGateCollector())
    handler.ActivieHTTPHandler(registry)

}