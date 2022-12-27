// GGexporter.
// GGexporter được dùng để thu thập các metric từ Oracle Goldengate. Nhận các flags
// sau đó active handler.
//
//   - -eP "0.0.0.0:9101" |"Exporter Port" | Hay "Bind address" là địa chỉ mà ta sẽ lấy metrics bằng prometheus.
//
//   - -mH "10.0.0.201" | "Manager Host 's address" | đôi khi các địa chỉ trong label trả về thành localhost flag này sẽ chuyển thành giá  trị do người dùng định nghĩa.
//
//   - -mP "1616" | "Manager's Port" | Không cần thiết phải cấu hình có thể bỏ qua.
//
//   - -rU "http://gg-svmgr.io" | "GG's Metrics Web Address" IP:port | Nơi mà exporter này sẽ lấy số liệu, đây cũng là performance server.
//
// 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
package main

import (
	"GGexporter/handler"
	"GGexporter/model"
	"GGexporter/services"
	"flag"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

// Clean returns the shortest path name equivalent to path
var config model.Config

// Agr
func getFlag(config *model.Config) {
	flag.StringVar(&config.Bind, "eP", "0.0.0.0:9101", "Exporter Port")
	flag.StringVar(&config.MgrHost, "mH", "10.0.0.201", "Manager Host 's address")
	flag.StringVar(&config.MgrPort, "mP", "1616", "Manager's Port")
	flag.StringVar(&config.RootURL, "rU", "http://gg-svmgr.io", "GG's Metrics Web Address")

	debug := flag.Bool("debug", false, "Debug true to show all bebug log")
	warn := flag.Bool("warn", false, "Warn true to show form warning log ")
	errl := flag.Bool("error", true, "Error true to show error log ")

	flag.Parse()

	if *debug {
		*errl = false
		log.Infof("-Debug log: %t ", *debug)
		log.SetLevel(log.DebugLevel)
	}
	if *warn {
		*errl = false
		log.Infof("-Warning log: %t ", *warn)
		log.SetLevel(log.WarnLevel)
	}
	if *errl {
		log.Infof("-Error log: %t ", *errl)
		log.SetLevel(log.ErrorLevel)
	}
}

// Main
func main() {
	// Gan cac flag vao config
	getFlag(&config)

	// registry thu thap metric => xem services
	registry := prometheus.NewRegistry()
	registry.MustRegister(services.NewGoldenGateCollector(config))

	// host web & khai bao flags => xem handler
	handler.ActiveHTTPHandler(registry, config.Bind)

}
