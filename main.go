package main

import (

    "GGexporter/services"

    "GGexporter/handler"
    "github.com/prometheus/client_golang/prometheus"
)
var mgrhost string
var mgrport string

func main() {
 
    mgrhost = "10.0.0.201"
    mgrport = "1616"

    registry := prometheus.NewRegistry() 
    registry.MustRegister(services.NewGoldenGateCollector())


    handler.ActivieHTTPHandler(9101,registry)


}