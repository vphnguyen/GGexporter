package main

import (
	"GGexporter/entities"
	"GGexporter/handler"
	"GGexporter/services"
	"flag"
	"github.com/prometheus/client_golang/prometheus"
)

var config entities.Config

func getFlag(config *entities.Config){

    flag.StringVar(&config.Bind, "eP", "0.0.0.0:9101", "Exporter Port")
    flag.StringVar(&config.MgrHost, "mH", "10.0.0.201", "Manager Host")
    flag.StringVar(&config.MgrPort, "mP", "1616", "Manager Port")
    flag.StringVar(&config.RootURL, "rU", "http://gg-svmgr.io", "RootURL")
    flag.Parse()

}

func main() {
	// Gan cac flag vao config
    getFlag(&config)
	
	// registry thu thap metric => xem services
	registry := prometheus.NewRegistry()
	registry.MustRegister(services.NewGoldenGateCollector(config))

	// host web & khai bao flags => xem handler
	handler.ActiveHTTPHandler(registry, config.Bind)

}

