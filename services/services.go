package services

/*
*
*
*   Chiu trach nhiem khai bao va khoi tao cac collector
*   Thu thap metric o phan collect
*   Goi ham chuyen xml thanh Object
*   Goi ham chuyen Object thanh Metric
*
*
*
*
 */

import (
	"GGexporter/model"
	"GGexporter/storage"
	"fmt"
	"reflect"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

const collector = "GoldenGate"

var config model.Config

type GoldenGateCollector struct {
	statusMetric     *prometheus.Desc
	trail_rba_Metric *prometheus.Desc
	trail_seq_Metric *prometheus.Desc
	//-- w
	trail_iowc_Metric *prometheus.Desc
	trail_iowb_Metric *prometheus.Desc
	//-- r
	trail_iorc_Metric *prometheus.Desc
	trail_iorb_Metric *prometheus.Desc
	//== EXT w_ghi vao trail
	trail_max_bytes_Metric *prometheus.Desc
	//== PUMP r_doc trail &  w_ghi vao trail
	//== REP r_doc trail
	statistics_Metric *prometheus.Desc
}

// ===== Khai bao cac describe o ben duoi
func (collector *GoldenGateCollector) Describe(ch chan<- *prometheus.Desc) {
	// == STATUS & RBA + SEQ
	ch <- collector.statusMetric
	ch <- collector.trail_rba_Metric
	ch <- collector.trail_seq_Metric

	//== EXTRACT
	ch <- collector.trail_iowc_Metric
	ch <- collector.trail_iowb_Metric
	ch <- collector.trail_max_bytes_Metric

	//== PUMP
	ch <- collector.trail_iorc_Metric
	ch <- collector.trail_iorb_Metric
	ch <- collector.statistics_Metric
}

// ===== Chi tiet hon cac decribe
func NewGoldenGateCollector(c model.Config) *GoldenGateCollector {
	config = c
	return &GoldenGateCollector{
		// === STATUS & RBA + SEQ
		statusMetric: prometheus.NewDesc(
			prometheus.BuildFQName(collector, "", "status"),
			"Shows status of golden gate instances _type 2:Capture:EXTRACT 4:pump:EXTRACT 3:Delivery:REPLICAT 14:PMSRVR 1:MANAGER _status 3:running 6:stopped 8:append 1:Registered never executed",
			[]string{"mgr_host", "group_name", "type"}, nil,
		),
		// ==
		trail_seq_Metric: prometheus.NewDesc(
			prometheus.BuildFQName(collector, "", "trail_seq"),
			"Trail Output _ trail_seq _ rotate times",
			[]string{"trail_name", "trail_path", "hostname", "group_name"}, nil,
		),
		trail_rba_Metric: prometheus.NewDesc(
			prometheus.BuildFQName(collector, "", "trail_rba"),
			"Trail Output _ trail_rba _ current bytes size of trail",
			[]string{"trail_name", "trail_path", "hostname", "group_name"}, nil,
		),
		// == WRITE
		trail_iowc_Metric: prometheus.NewDesc(
			prometheus.BuildFQName(collector, "", "io_write_count"),
			"Trail Output _ io write count",
			[]string{"trail_name", "trail_path", "hostname", "group_name"}, nil,
		),
		trail_iowb_Metric: prometheus.NewDesc(
			prometheus.BuildFQName(collector, "", "io_write_bytes"),
			"Trail Output _ io write bytes",
			[]string{"trail_name", "trail_path", "hostname", "group_name"}, nil,
		),
		// == READ
		trail_iorc_Metric: prometheus.NewDesc(
			prometheus.BuildFQName(collector, "", "io_read_count"),
			"PUMP Trail Output _ io read count",
			[]string{"trail_name", "trail_path", "hostname", "group_name"}, nil,
		),
		trail_iorb_Metric: prometheus.NewDesc(
			prometheus.BuildFQName(collector, "", "io_read_bytes"),
			"Trail Output _ io read bytes",
			[]string{"trail_name", "trail_path", "hostname", "group_name"}, nil,
		),
		//==== EXTRACT
		trail_max_bytes_Metric: prometheus.NewDesc(
			prometheus.BuildFQName(collector, "", "extract_trail_max_bytes"),
			"Trail Output _ extract_trail_max_bytes _ Max size of a trail can be reach before rotate",
			[]string{"trail_name", "trail_path", "hostname", "group_name"}, nil,
		),
		//==== PUMP
		statistics_Metric: prometheus.NewDesc(
			prometheus.BuildFQName(collector, "", "statistics"),
			"statistics_Metric HELP",
			[]string{"hostname", "group_name", "mapped"}, nil,
		),
	}
}

//------ Khi request se bat dau lay metric (Neu khong lay duoc se delay & timeout)
func (collector *GoldenGateCollector) Collect(ch chan<- prometheus.Metric) {

	/*
	 *   Luu cac chuoi byte thanh cac object local
	 *   Moi vong lap request metric se reset cac object nay
	 */
	var (
		manager           model.ManagerModel
		performanceServer model.PerformanceServerModel
		listOfExtract     []model.ExtractModel
		listOfPump        []model.PumpModel
		listOfReplicat    []model.ReplicatModel
	)
	mgroups, err := storage.GetGroups(config.RootURL)
	if err != nil {
		fmt.Println(err)
		//panic("Service - khong the parser Object - groups")
	}
	for _, aGroup := range mgroups.GroupRefs {
		if aGroup.IsExtract() {
			anExtract, _ := storage.GetExtract(config.RootURL, aGroup.URL)
			listOfExtract = append(listOfExtract, anExtract)
			continue
		}
		if aGroup.IsPump() {
			aPump, _ := storage.GetPump(config.RootURL, aGroup.URL)
			listOfPump = append(listOfPump, *aPump)
			continue
		}
		if aGroup.IsManager() {
			manager, _ = storage.GetManager(config.RootURL, aGroup.URL)
			continue
		}
		if aGroup.IsPerformanceServer() {
			performanceServer, _ = storage.GetPerformanceServer(config.RootURL, aGroup.URL)
			continue
		}
		if aGroup.IsReplicat() {
			aReplicat, _ := storage.GetReplicat(config.RootURL, aGroup.URL)
			listOfReplicat = append(listOfReplicat, aReplicat)
			continue
		}

	}
	getMetrics(ch, collector, &manager, &performanceServer, &listOfExtract, &listOfPump, &listOfReplicat)

}

// ===== Ham chiu trach nhiem render cac metric tu object va truyen vao cac collector
func getMetrics(ch chan<- prometheus.Metric, collector *GoldenGateCollector,
	manager *model.ManagerModel,
	performanceServer *model.PerformanceServerModel,
	listOfExtract *[]model.ExtractModel,
	listOfPump *[]model.PumpModel,
	listOfReplicat *[]model.ReplicatModel) {

	// ===== MGR        =======
	ch <- prometheus.MustNewConstMetric(collector.statusMetric,
		prometheus.GaugeValue,
		getMetricValue(manager.Process.Status),
		[]string{config.MgrHost, manager.Process.Name, typeToString(manager.Process.Type)}...)
	// ===== Extract    =======
	for _, extract := range *listOfExtract {
		ch <- prometheus.MustNewConstMetric(collector.statusMetric,
			prometheus.GaugeValue,
			getMetricValue(extract.Process.Status),
			[]string{config.MgrHost, extract.Process.Name, typeToString(extract.Process.Type)}...)
		for _, trail := range extract.Process.TrailOutput {
			//========== io_write_count     "trail_name","trail_path","hostname","group_name"
			ch <- prometheus.MustNewConstMetric(collector.trail_iowc_Metric,
				prometheus.GaugeValue,
				getMetricValue(trail.IoWriteCount),
				[]string{trail.TrailName, trail.TrailPath, trail.Hostname, extract.Process.Name}...)
			//========== io_write_bytes
			ch <- prometheus.MustNewConstMetric(collector.trail_iowb_Metric,
				prometheus.GaugeValue,
				getMetricValue(trail.IoWriteBytes),
				[]string{trail.TrailName, trail.TrailPath, trail.Hostname, extract.Process.Name}...)
			//========== trail_rba_Metric
			ch <- prometheus.MustNewConstMetric(collector.trail_rba_Metric,
				prometheus.GaugeValue,
				getMetricValue(trail.TrailRba),
				[]string{trail.TrailName, trail.TrailPath, trail.Hostname, extract.Process.Name}...)
			//========== trail_seq_Metric
			ch <- prometheus.MustNewConstMetric(collector.trail_seq_Metric,
				prometheus.GaugeValue,
				getMetricValue(trail.TrailSeq),
				[]string{trail.TrailName, trail.TrailPath, trail.Hostname, extract.Process.Name}...)
			//========== extract_trail_max_bytes_Metric
			ch <- prometheus.MustNewConstMetric(collector.trail_max_bytes_Metric,
				prometheus.GaugeValue,
				getMetricValue(trail.TrailMaxBytes),
				[]string{trail.TrailName, trail.TrailPath, trail.Hostname, extract.Process.Name}...)
		}
		a := reflect.ValueOf(&extract.Process.StatisticsExtract).Elem()
		for i := 0; i < (a.NumField()); i++ {
			if a.Type().Field(i).Name != "Text" {
				ch <- prometheus.MustNewConstMetric(collector.statistics_Metric,
					prometheus.GaugeValue,
					getMetricValue(fmt.Sprintf("%s", a.Field(i).Interface())),
					[]string{config.MgrHost, extract.Process.Name, a.Type().Field(i).Name}...)
			}
		}
	}

	// ===== PUMP       =======
	for _, pump := range *listOfPump {
		ch <- prometheus.MustNewConstMetric(collector.statusMetric,
			prometheus.GaugeValue,
			getMetricValue(pump.Process.Status),
			[]string{config.MgrHost, pump.Process.Name, typeToString(pump.Process.Type)}...)
		// === Trail in
		// -- REad
		ch <- prometheus.MustNewConstMetric(collector.trail_iorc_Metric,
			prometheus.GaugeValue,
			getMetricValue(pump.Process.TrailInput.IoReadCount),
			[]string{pump.Process.TrailInput.TrailName, pump.Process.TrailInput.TrailPath, config.MgrHost, pump.Process.Name}...)
		ch <- prometheus.MustNewConstMetric(collector.trail_iorb_Metric,
			prometheus.GaugeValue,
			getMetricValue(pump.Process.TrailInput.IoReadBytes),
			[]string{pump.Process.TrailInput.TrailName, pump.Process.TrailInput.TrailPath, config.MgrHost, pump.Process.Name}...)
		// -- RBA - SEQ
		ch <- prometheus.MustNewConstMetric(collector.trail_rba_Metric,
			prometheus.GaugeValue,
			getMetricValue(pump.Process.TrailInput.TrailRba),
			[]string{pump.Process.TrailInput.TrailName, pump.Process.TrailInput.TrailPath, config.MgrHost, pump.Process.Name}...)

		ch <- prometheus.MustNewConstMetric(collector.trail_seq_Metric,
			prometheus.GaugeValue,
			getMetricValue(pump.Process.TrailInput.TrailSeq),
			[]string{pump.Process.TrailInput.TrailName, pump.Process.TrailInput.TrailPath, config.MgrHost, pump.Process.Name}...)
		// === Trail out (s)
		for _, trailout := range pump.Process.TrailOutput {
			// -- WRITE
			ch <- prometheus.MustNewConstMetric(collector.trail_iowc_Metric,
				prometheus.GaugeValue,
				getMetricValue(trailout.IoWriteCount),
				[]string{trailout.TrailName, trailout.TrailPath, trailout.Hostname, pump.Process.Name}...)
			ch <- prometheus.MustNewConstMetric(collector.trail_iowb_Metric,
				prometheus.GaugeValue,
				getMetricValue(trailout.IoWriteBytes),
				[]string{trailout.TrailName, trailout.TrailPath, trailout.Hostname, pump.Process.Name}...)
			// -- RBA + SEQ
			ch <- prometheus.MustNewConstMetric(collector.trail_rba_Metric,
				prometheus.GaugeValue,
				getMetricValue(trailout.TrailRba),
				[]string{trailout.TrailName, trailout.TrailPath, trailout.Hostname, pump.Process.Name}...)
			ch <- prometheus.MustNewConstMetric(collector.trail_seq_Metric,
				prometheus.GaugeValue,
				getMetricValue(trailout.TrailSeq),
				[]string{trailout.TrailName, trailout.TrailPath, trailout.Hostname, pump.Process.Name}...)
			//========== trail_max_bytes_Metric
			ch <- prometheus.MustNewConstMetric(collector.trail_max_bytes_Metric,
				prometheus.GaugeValue,
				getMetricValue(trailout.TrailMaxBytes),
				[]string{trailout.TrailName, trailout.TrailPath, trailout.Hostname, pump.Process.Name}...)
		}
	}

	// ===== PMSRVR     =======
	ch <- prometheus.MustNewConstMetric(collector.statusMetric,
		prometheus.GaugeValue,
		getMetricValue(performanceServer.Process.Status),
		[]string{config.MgrHost, performanceServer.Process.Name, typeToString(performanceServer.Process.Type)}...)

	// ===== REPLICAT   =======
	for _, replicat := range *listOfReplicat {
		ch <- prometheus.MustNewConstMetric(collector.statusMetric,
			prometheus.GaugeValue,
			getMetricValue(replicat.Process.Status),
			[]string{config.MgrHost, replicat.Process.Name, typeToString(replicat.Process.Type)}...)

		for _, trailin := range replicat.Process.TrailInput {
			// -- Read
			ch <- prometheus.MustNewConstMetric(collector.trail_iorc_Metric,
				prometheus.GaugeValue,
				getMetricValue(trailin.IoReadCount),
				[]string{trailin.TrailName, trailin.TrailPath, config.MgrHost, replicat.Process.Name}...)
			ch <- prometheus.MustNewConstMetric(collector.trail_iorb_Metric,
				prometheus.GaugeValue,
				getMetricValue(trailin.IoReadBytes),
				[]string{trailin.TrailName, trailin.TrailPath, config.MgrHost, replicat.Process.Name}...)
			// -- RBA + SEQ
			ch <- prometheus.MustNewConstMetric(collector.trail_rba_Metric,
				prometheus.GaugeValue,
				getMetricValue(trailin.TrailRba),
				[]string{trailin.TrailName, trailin.TrailPath, config.MgrHost, replicat.Process.Name}...)
			ch <- prometheus.MustNewConstMetric(collector.trail_seq_Metric,
				prometheus.GaugeValue,
				getMetricValue(trailin.TrailSeq),
				[]string{trailin.TrailName, trailin.TrailPath, config.MgrHost, replicat.Process.Name}...)
		}
		// Dem so luong field trong Statistics sau do chuyen thanh Lable
		a := reflect.ValueOf(&replicat.Process.StatisticsReplicat).Elem()
		for i := 0; i < (a.NumField()); i++ {
			if a.Type().Field(i).Name != "Text" {
				ch <- prometheus.MustNewConstMetric(collector.statistics_Metric,
					prometheus.GaugeValue,
					getMetricValue(fmt.Sprintf("%s", a.Field(i).Interface())),
					[]string{config.MgrHost, replicat.Process.Name, a.Type().Field(i).Name}...)
			}
		}

	}
}

//------ Chuyen tu string trong object thanh float64 phu hop voi metric gauge
func getMetricValue(input string) float64 {
	metric, er := strconv.ParseFloat(input, 64)
	if er != nil {
		log.Errorf("Services.getMetricValue. Noi dung dau vao (%s) khong phu hop", input)
	}
	return metric
}

//------ Chuyen tu string type trong object thanh cac string day du, de hieu
func typeToString(inputString string) string {
	if inputString == model.TYPE_PMSRVR {
		return "Performance_Metrics_Server"
	}
	if inputString == model.TYPE_MGR {
		return "Manager"
	}
	if inputString == model.TYPE_EXTRACT {
		return "Extract_Capture"
	}
	if inputString == model.TYPE_PUMP {
		return "Extract_Pump"
	}
	if inputString == model.TYPE_REPLICAT {
		return "Replicat_Delivery"
	}
	log.Warnf("Collector.Khong the chuyen type %s thanh string", inputString)
	return "Unknown"

}
