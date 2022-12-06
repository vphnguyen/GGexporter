// Handler prometheus webhook
package handler

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

// Handler prometheus webhook at  bind [0.0.0.0:9101]
func ActiveHTTPHandler(registry *prometheus.Registry, bind string) {
	//log.Infof("MGR: http://%s:%s  OR (%s)=> bind %s", entities.config.MgrHost,entities.config.MgrPort, entities.config.RootURL, entities.config.bind )

	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		h := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
		h.ServeHTTP(w, r)
	})

	// start server
	log.Infof("Starting http server at: %s", bind)
	if err := http.ListenAndServe(bind, nil); err != nil {
		log.Errorf("Failed to start http server: %s", err)
	}
}
