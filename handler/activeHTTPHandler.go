package handler


import (

	"flag"
	log "github.com/sirupsen/logrus"
	"net/http"
    "GGexporter/services"
    "GGexporter/entities"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

)
func ActivieHTTPHandler( registry *prometheus.Registry){


// === Chuyen tu port sang thanh flag bind
    services.MgrHost = "10.0.0.201"
    services.MgrPort = "1616"


	var bind string
    flag.StringVar(&bind, "eP", "0.0.0.0:9101", "Exporter Port")
    flag.StringVar(&services.MgrHost, "mH", "10.0.0.201" , "Manager Host")
    flag.StringVar(&services.MgrPort, "mP", "1616"       , "Manager Port")
    flag.StringVar(&entities.RootURL, "rU", "http://gg-svmgr.io"      , "RootURL")
    flag.Parse()


    log.Infof("MGR: http://%s:%s  OR (%s)=> bind %s", services.MgrHost,services.MgrPort, entities.RootURL,bind )


    http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
        h := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
        h.ServeHTTP(w, r)
    })

// start server
    log.Infof("Starting http server - %s", bind)
    if err := http.ListenAndServe(bind, nil); err != nil {
        log.Errorf("Failed to start http server: %s", err)
    }
}

