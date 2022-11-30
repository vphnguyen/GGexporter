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
    "strings"
	"strconv"
	"github.com/relvacode/iso8601"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)
type ExternalAPIResponse struct {
	Timestamp *iso8601.Time
}
const collector = "GoldenGate"

var config model.Config

type GoldenGateCollector struct {
	metricStatus     *prometheus.Desc
	metricTrailRba *prometheus.Desc
	metricTrailSeq *prometheus.Desc
	//-- w
	metricTrailIoWriteCount *prometheus.Desc
	metricTrailIoWriteByte *prometheus.Desc
	//-- r
	metricTrailIoReadCount *prometheus.Desc
	metricTrailIoReadByte *prometheus.Desc
	//== EXT w_ghi vao trail
	metricTrailMaxBytes *prometheus.Desc
	//== PUMP r_doc trail &  w_ghi vao trail
	//== REP r_doc trail
	metricStatistics *prometheus.Desc
	metricRepLag *prometheus.Desc
}

// ===== Khai bao cac describe o ben duoi
func (collector *GoldenGateCollector) Describe(ch chan<- *prometheus.Desc) {
	// == STATUS & RBA + SEQ
	ch <- collector.metricStatus
	ch <- collector.metricTrailRba
	ch <- collector.metricTrailSeq

	//== EXTRACT
	ch <- collector.metricTrailIoWriteCount
	ch <- collector.metricTrailIoWriteByte
	ch <- collector.metricTrailMaxBytes

	//== PUMP
	ch <- collector.metricTrailIoReadCount
	ch <- collector.metricTrailIoReadByte
	ch <- collector.metricStatistics
	ch <- collector.metricRepLag
}

// ===== Chi tiet hon cac decribe
func NewGoldenGateCollector(c model.Config) *GoldenGateCollector {
	config = c
	return &GoldenGateCollector{
		// === STATUS & RBA + SEQ
		metricStatus: prometheus.NewDesc(
			prometheus.BuildFQName(collector, "", "status"),
			"Shows status of golden gate instances _type 2:Capture:EXTRACT 4:pump:EXTRACT 3:Delivery:REPLICAT 14:PMSRVR 1:MANAGER _status 3:running 6:stopped 8:append 1:Registered never executed",
			[]string{"mgr_host", "group_name", "type"}, nil,
		),
		// ==
		metricTrailSeq: prometheus.NewDesc(
			prometheus.BuildFQName(collector, "", "trail_seq"),
			"Trail Output _ trail_seq _ rotate times",
			[]string{"trail_name", "trail_path", "hostname", "group_name"}, nil,
		),
		metricTrailRba: prometheus.NewDesc(
			prometheus.BuildFQName(collector, "", "trail_rba"),
			"Trail Output _ trail_rba _ current bytes size of trail",
			[]string{"trail_name", "trail_path", "hostname", "group_name"}, nil,
		),
		// == WRITE
		metricTrailIoWriteCount: prometheus.NewDesc(
			prometheus.BuildFQName(collector, "", "io_write_count"),
			"Trail Output _ io write count",
			[]string{"trail_name", "trail_path", "hostname", "group_name"}, nil,
		),
		metricTrailIoWriteByte: prometheus.NewDesc(
			prometheus.BuildFQName(collector, "", "io_write_bytes"),
			"Trail Output _ io write bytes",
			[]string{"trail_name", "trail_path", "hostname", "group_name"}, nil,
		),
		// == READ
		metricTrailIoReadCount: prometheus.NewDesc(
			prometheus.BuildFQName(collector, "", "io_read_count"),
			"PUMP Trail Output _ io read count",
			[]string{"trail_name", "trail_path", "hostname", "group_name"}, nil,
		),
		metricTrailIoReadByte: prometheus.NewDesc(
			prometheus.BuildFQName(collector, "", "io_read_bytes"),
			"Trail Output _ io read bytes",
			[]string{"trail_name", "trail_path", "hostname", "group_name"}, nil,
		),
		//==== EXTRACT
		metricTrailMaxBytes: prometheus.NewDesc(
			prometheus.BuildFQName(collector, "", "extract_trail_max_bytes"),
			"Trail Output _ extract_trail_max_bytes _ Max size of a trail can be reach before rotate",
			[]string{"trail_name", "trail_path", "hostname", "group_name"}, nil,
		),
		//==== PUMP
		metricStatistics: prometheus.NewDesc(
			prometheus.BuildFQName(collector, "", "statistics"),
			"metricStatistics HELP",
			[]string{"hostname", "group_name", "mapped"}, nil,
		),
		metricRepLag: prometheus.NewDesc(
			prometheus.BuildFQName(collector, "", "replicat_lag"),
			"last_operation_ts_Metric sub HELP",
			[]string{ "group_name"}, nil,
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
			anExtract, er := storage.GetExtract(config.RootURL, aGroup.URL)
			if er != nil {
				log.Warnf("%s",er)
				log.Warnf("Service - Could be an InitLoad Extract - Skipped ")
				continue
			}
			listOfExtract = append(listOfExtract, anExtract)
			continue
		}
		if aGroup.IsPump() {
			aPump, er := storage.GetPump(config.RootURL, aGroup.URL)
			if er != nil {
				log.Warnf("%s",er)
				continue
			}
			listOfPump = append(listOfPump, *aPump)
			continue
		}
		if aGroup.IsManager() {
			var er error
			manager, _ = storage.GetManager(config.RootURL, aGroup.URL)
			if er != nil {
				log.Warnf("%s",er)
				continue
			}
			continue
		}
		if aGroup.IsPerformanceServer() {
			var er error
			performanceServer, _ = storage.GetPerformanceServer(config.RootURL, aGroup.URL)
			if er != nil {
				log.Warnf("%s",er)
				continue
			}
			continue
		}
		if aGroup.IsReplicat() {
			aReplicat, er := storage.GetReplicat(config.RootURL, aGroup.URL)
			if er != nil {
				log.Warnf("%s",er)
				log.Warnf("Service - Could be an InitLoad Replicat - Skipped ")
				continue
			}
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
	ch <- prometheus.MustNewConstMetric(collector.metricStatus,
		prometheus.GaugeValue,
		toFloat64(manager.Process.Status),
		[]string{config.MgrHost, manager.Process.Name, typeToString(manager.Process.Type)}...)
	// ===== Extract    =======
	for _, extract := range *listOfExtract {
		ch <- prometheus.MustNewConstMetric(collector.metricStatus,
			prometheus.GaugeValue,
			toFloat64(extract.Process.Status),
			[]string{config.MgrHost, extract.Process.Name, typeToString(extract.Process.Type)}...)

		ch <- prometheus.MustNewConstMetric(collector.metricRepLag,
			prometheus.GaugeValue,
			getLagTime(extract.Process.PositionEr.LastCheckpointTs,extract.Process.PositionEr.OutputCheckpoint,"extract"),
			extract.Process.Name)
		for _, trail := range extract.Process.TrailOutput {
			//========== io_write_count     "trail_name","trail_path","hostname","group_name"
			ch <- prometheus.MustNewConstMetric(collector.metricTrailIoWriteCount,
				prometheus.GaugeValue,
				toFloat64(trail.IoWriteCount),
				[]string{trail.TrailName, trail.TrailPath, trail.Hostname, extract.Process.Name}...)
			//========== io_write_bytes
			ch <- prometheus.MustNewConstMetric(collector.metricTrailIoWriteByte,
				prometheus.GaugeValue,
				toFloat64(trail.IoWriteBytes),
				[]string{trail.TrailName, trail.TrailPath, trail.Hostname, extract.Process.Name}...)
			//========== metricTrailRba
			ch <- prometheus.MustNewConstMetric(collector.metricTrailRba,
				prometheus.GaugeValue,
				toFloat64(trail.TrailRba),
				[]string{trail.TrailName, trail.TrailPath, trail.Hostname, extract.Process.Name}...)
			//========== metricTrailSeq
			ch <- prometheus.MustNewConstMetric(collector.metricTrailSeq,
				prometheus.GaugeValue,
				toFloat64(trail.TrailSeq),
				[]string{trail.TrailName, trail.TrailPath, trail.Hostname, extract.Process.Name}...)
			//========== extract_metricTrailMaxBytes
			ch <- prometheus.MustNewConstMetric(collector.metricTrailMaxBytes,
				prometheus.GaugeValue,
				toFloat64(trail.TrailMaxBytes),
				[]string{trail.TrailName, trail.TrailPath, trail.Hostname, extract.Process.Name}...)
		}
		a := reflect.ValueOf(&extract.Process.StatisticsExtract).Elem()
		for i := 0; i < (a.NumField()); i++ {
			if a.Type().Field(i).Name != "Text" {
				ch <- prometheus.MustNewConstMetric(collector.metricStatistics,
					prometheus.GaugeValue,
					toFloat64(fmt.Sprintf("%s", a.Field(i).Interface())),
					[]string{config.MgrHost, extract.Process.Name, a.Type().Field(i).Name}...)
			}
		}
	}

	// ===== PUMP       =======
	for _, pump := range *listOfPump {
		ch <- prometheus.MustNewConstMetric(collector.metricStatus,
			prometheus.GaugeValue,
			toFloat64(pump.Process.Status),
			[]string{config.MgrHost, pump.Process.Name, typeToString(pump.Process.Type)}...)

		ch <- prometheus.MustNewConstMetric(collector.metricRepLag,
			prometheus.GaugeValue,
			getLagTime4P(	pump.Process.PositionEr.LastOperationTs,
							pump.Process.PositionEr.InputCheckpoint,
							pump.Process.PositionEr.LastCheckpointTs,
							pump.Process.PositionEr.OutputCheckpoint ), pump.Process.Name)

		// === Trail in
		// -- REad
		ch <- prometheus.MustNewConstMetric(collector.metricTrailIoReadCount,
			prometheus.GaugeValue,
			toFloat64(pump.Process.TrailInput.IoReadCount),
			[]string{pump.Process.TrailInput.TrailName, pump.Process.TrailInput.TrailPath, config.MgrHost, pump.Process.Name}...)
		ch <- prometheus.MustNewConstMetric(collector.metricTrailIoReadByte,
			prometheus.GaugeValue,
			toFloat64(pump.Process.TrailInput.IoReadBytes),
			[]string{pump.Process.TrailInput.TrailName, pump.Process.TrailInput.TrailPath, config.MgrHost, pump.Process.Name}...)
		// -- RBA - SEQ
		ch <- prometheus.MustNewConstMetric(collector.metricTrailRba,
			prometheus.GaugeValue,
			toFloat64(pump.Process.TrailInput.TrailRba),
			[]string{pump.Process.TrailInput.TrailName, pump.Process.TrailInput.TrailPath, config.MgrHost, pump.Process.Name}...)

		ch <- prometheus.MustNewConstMetric(collector.metricTrailSeq,
			prometheus.GaugeValue,
			toFloat64(pump.Process.TrailInput.TrailSeq),
			[]string{pump.Process.TrailInput.TrailName, pump.Process.TrailInput.TrailPath, config.MgrHost, pump.Process.Name}...)
		// === Trail out (s)
		for _, trailout := range pump.Process.TrailOutput {
			// -- WRITE
			ch <- prometheus.MustNewConstMetric(collector.metricTrailIoWriteCount,
				prometheus.GaugeValue,
				toFloat64(trailout.IoWriteCount),
				[]string{trailout.TrailName, trailout.TrailPath, trailout.Hostname, pump.Process.Name}...)
			ch <- prometheus.MustNewConstMetric(collector.metricTrailIoWriteByte,
				prometheus.GaugeValue,
				toFloat64(trailout.IoWriteBytes),
				[]string{trailout.TrailName, trailout.TrailPath, trailout.Hostname, pump.Process.Name}...)
			// -- RBA + SEQ
			ch <- prometheus.MustNewConstMetric(collector.metricTrailRba,
				prometheus.GaugeValue,
				toFloat64(trailout.TrailRba),
				[]string{trailout.TrailName, trailout.TrailPath, trailout.Hostname, pump.Process.Name}...)
			ch <- prometheus.MustNewConstMetric(collector.metricTrailSeq,
				prometheus.GaugeValue,
				toFloat64(trailout.TrailSeq),
				[]string{trailout.TrailName, trailout.TrailPath, trailout.Hostname, pump.Process.Name}...)
			//========== metricTrailMaxBytes
			ch <- prometheus.MustNewConstMetric(collector.metricTrailMaxBytes,
				prometheus.GaugeValue,
				toFloat64(trailout.TrailMaxBytes),
				[]string{trailout.TrailName, trailout.TrailPath, trailout.Hostname, pump.Process.Name}...)
		}
	}

	// ===== PMSRVR     =======
	ch <- prometheus.MustNewConstMetric(collector.metricStatus,
		prometheus.GaugeValue,
		toFloat64(performanceServer.Process.Status),
		[]string{config.MgrHost, performanceServer.Process.Name, typeToString(performanceServer.Process.Type)}...)

	// ===== REPLICAT   =======
	for _, replicat := range *listOfReplicat {
		ch <- prometheus.MustNewConstMetric(collector.metricStatus,
			prometheus.GaugeValue,
			toFloat64(replicat.Process.Status),
			[]string{config.MgrHost, replicat.Process.Name, typeToString(replicat.Process.Type)}...)
		ch <- prometheus.MustNewConstMetric(collector.metricRepLag,
			prometheus.GaugeValue,
			getLagTime(replicat.Process.PositionEr.LastOperationTs,replicat.Process.PositionEr.InputCheckpoint,"replicat"),
			replicat.Process.Name)


		for _, trailin := range replicat.Process.TrailInput {
			// -- Read
			ch <- prometheus.MustNewConstMetric(collector.metricTrailIoReadCount,
				prometheus.GaugeValue,
				toFloat64(trailin.IoReadCount),
				[]string{trailin.TrailName, trailin.TrailPath, config.MgrHost, replicat.Process.Name}...)
			ch <- prometheus.MustNewConstMetric(collector.metricTrailIoReadByte,
				prometheus.GaugeValue,
				toFloat64(trailin.IoReadBytes),
				[]string{trailin.TrailName, trailin.TrailPath, config.MgrHost, replicat.Process.Name}...)
			// -- RBA + SEQ
			ch <- prometheus.MustNewConstMetric(collector.metricTrailRba,
				prometheus.GaugeValue,
				toFloat64(trailin.TrailRba),
				[]string{trailin.TrailName, trailin.TrailPath, config.MgrHost, replicat.Process.Name}...)
			ch <- prometheus.MustNewConstMetric(collector.metricTrailSeq,
				prometheus.GaugeValue,
				toFloat64(trailin.TrailSeq),
				[]string{trailin.TrailName, trailin.TrailPath, config.MgrHost, replicat.Process.Name}...)
		}
		// Dem so luong field trong Statistics sau do chuyen thanh Lable
		a := reflect.ValueOf(&replicat.Process.StatisticsReplicat).Elem()
		for i := 0; i < (a.NumField()); i++ {
			if a.Type().Field(i).Name != "Text" {
				ch <- prometheus.MustNewConstMetric(collector.metricStatistics,
					prometheus.GaugeValue,
					toFloat64(fmt.Sprintf("%s", a.Field(i).Interface())),
					[]string{config.MgrHost, replicat.Process.Name, a.Type().Field(i).Name}...)
			}
		}

	}
}

//------ Chuyen tu string trong object thanh float64 phu hop voi metric gauge
func toFloat64(input string) float64 {
	metric, er := strconv.ParseFloat(input, 64)
	if er != nil {
		log.Warnf("Services.toFloat64. Noi dung dau vao (%s) khong phu hop", input)
	}
	return metric
}

func getLagTime4P(input string, input2 string ,input3 string ,input4 string) float64 {

		a1, _ := iso8601.ParseString(input)
		a2, er := iso8601.ParseString(strings.Replace(strings.Trim(strings.Split(input2,"\n")[3] ,"Timestamp: ")," ","T",1))
		if er != nil {
			log.Warnf("Service.Collector.getLagTime(%s) khong phu hop", input)
		}
		log.Warnf("1 1 (%s) ", a1)
		log.Warnf("1 2 (%s) ", a2)
		v1 := float64(a1.Sub(a2).Microseconds())
		log.Warnf("= (%f)\n ", v1)

		l, _ := iso8601.ParseString(input3)
		o, er := iso8601.ParseString(strings.Replace(strings.Trim(strings.Split(input4,"\n")[3] ,"Timestamp: ")," ","T",1))
		if er != nil {
			log.Warnf("Service.Collector.getLagTime(%s) khong phu hop", input3)
		}


		log.Warnf("2 1 (%s) ", l)
		log.Warnf("2 2 (%s) ", o)
		v2 := float64(l.Sub(o).Microseconds())
		log.Warnf("= (%f)\n======== ", v2)
		return v1+v2	
}

func getLagTime(input string, input2 string ,tp string) float64 {
	if tp == "extract"{
		last, _ := iso8601.ParseString(input)
		op, er := iso8601.ParseString(strings.Replace(strings.Trim(strings.Split(input2,"\n")[3] ,"Timestamp: ")," ","T",1))
		if er != nil {
			log.Warnf("Service.Collector.getLagTime(%s) khong phu hop", input)
		}
		metric := float64(last.Sub(op).Microseconds())
		return metric	
	} else{

		last, _ := iso8601.ParseString(input)
		op, er := iso8601.ParseString(strings.Replace(strings.Trim(strings.Split(input2,"\n")[3] ,"Timestamp: ")," ","T",1))

		if er != nil {
			log.Warnf("Service.Collector.getLagTime(%s) khong phu hop", input)
		}
		metric := float64(last.Sub(op).Microseconds())
		return metric	
	}
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
	log.Warnf("Service.Collector.Status.Khong the chuyen type %s thanh string", inputString)
	return "Unknown"
}
