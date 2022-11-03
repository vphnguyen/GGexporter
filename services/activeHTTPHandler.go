package services


import (
	"log"
   "net/http"
   "github.com/prometheus/client_golang/prometheus"
   "github.com/prometheus/client_golang/prometheus/promhttp"

)
func ActivieHTTPHandler(registry *prometheus.Registry){

	handler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{}) 
	http.Handle("/metrics", handler)
    log.Fatal(http.ListenAndServe(":9101", nil))

}

