package handler


import (
	"log"
	"strconv"
   "net/http"
   "github.com/prometheus/client_golang/prometheus"
   "github.com/prometheus/client_golang/prometheus/promhttp"

)
func ActivieHTTPHandler(port int, registry *prometheus.Registry){

	handler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{}) 
	http.Handle("/metrics", handler)
    log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))

}

