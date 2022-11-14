package main

import (
	"GGexporter/handler"
	"GGexporter/model"
	"GGexporter/services"
	"flag"

	"github.com/prometheus/client_golang/prometheus"
)

var config model.Config

func getFlag(config *model.Config) {
	flag.StringVar(&config.Bind, "eP", "0.0.0.0:9101", "Exporter Port")
	flag.StringVar(&config.MgrHost, "mH", "10.0.0.201", "Manager Host 's address")
	flag.StringVar(&config.MgrPort, "mP", "1616", "Manager's Port")
	flag.StringVar(&config.RootURL, "rU", "http://gg-svmgr.io", "GG's Metrics Web Address")
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
