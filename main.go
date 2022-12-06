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
package main

import (
	"GGexporter/handler"
	"GGexporter/model"
	"GGexporter/services"
	"flag"

	"github.com/prometheus/client_golang/prometheus"
)

// Clean returns the shortest path name equivalent to path
var config model.Config

// code
func getFlag(config *model.Config) {
	flag.StringVar(&config.Bind, "eP", "0.0.0.0:9101", "Exporter Port")
	flag.StringVar(&config.MgrHost, "mH", "10.0.0.201", "Manager Host 's address")
	flag.StringVar(&config.MgrPort, "mP", "1616", "Manager's Port")
	flag.StringVar(&config.RootURL, "rU", "http://gg-svmgr.io", "GG's Metrics Web Address")
	flag.Parse()
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
