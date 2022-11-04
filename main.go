package main

import (

    "GGexporter/services"

    "GGexporter/handler"
    "github.com/prometheus/client_golang/prometheus"
)


func main() {
 




    //collector := services.NewFooCollector()
    registry := prometheus.NewRegistry() 
    registry.MustRegister(services.NewGoldenGateCollector())


    handler.ActivieHTTPHandler(9101,registry)


}