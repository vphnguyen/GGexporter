package handler


import (

	"flag"
	log "github.com/sirupsen/logrus"
	"strconv"
	"net/http"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

)
func ActivieHTTPHandler(port int, registry *prometheus.Registry){
	var bind string
    flag.StringVar(&bind, "bind", "0.0.0.0:"+strconv.Itoa(port), "bind")
    flag.Parse()

    http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
        h := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
        h.ServeHTTP(w, r)
    })
	// handler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{}) 
	// http.Handle("/metrics", handler)
 //    log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))

    // start server
    log.Infof("Starting http server - %s", bind)
    if err := http.ListenAndServe(bind, nil); err != nil {
        log.Errorf("Failed to start http server: %s", err)
    }
}

