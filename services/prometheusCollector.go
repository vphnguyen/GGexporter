package services

import (
    "time"
    log "github.com/sirupsen/logrus"
    //"strconv"
    //"GGexporter/storage"
    "GGexporter/entities"
    "github.com/prometheus/client_golang/prometheus"
    //"math/rand"
    "fmt"
)
const  collector = "GoldenGate"
type GoldenGateCollector struct{}

// =============================
// Config config structure
// =============================
type Metric struct {
        Name        string
        Type        string
        Description string
        Labels      []string
        Value       string
        metricDesc  *prometheus.Desc
}


var (
    mgroups entities.MGroups
    mpointsofextract []entities.MpointsOfExtract
    mpointsofmgr entities.MpointsOfMGR
    mpointsofpmsrvr entities.MpointsOfPMSRVR
    
    //==   Metrics

    MetricList = []Metric { 
        Metric{ Name: "Name1", Type:"Gauge" , Description:"Des1" , Labels:[]string{"a", "b"}, Value:"11" },        
        Metric{ Name: "Name2", Type:"Gauge" , Description:"Des2" , Labels:[]string{"a", "b"}, Value:"12" },   
    }

)

func (e *GoldenGateCollector) Describe(ch chan<- *prometheus.Desc) { 

    for i, metric := range MetricList {
        MetricList[i].metricDesc = prometheus.NewDesc(
                prometheus.BuildFQName(collector, "",  metric.Name),
                metric.Description,
                metric.Labels, nil,
        )
        fmt.Println(metric.Name)
        fmt.Println(&metric.metricDesc)
        log.Infof("metric description for \"%s\" registerd",  metric.Name )
    }
}


func (e *GoldenGateCollector) Collect(ch chan<- prometheus.Metric) {
        
    ch <- prometheus.MustNewConstMetric(MetricList[1].metricDesc, prometheus.CounterValue, float64(time.Now().Second()), []string{"a2","a3"}...)
    ch <- prometheus.MustNewConstMetric(MetricList[0].metricDesc, prometheus.CounterValue, float64(time.Now().Second()), []string{"a2","a3"}...)
 
}


















// func NewFooCollector() *fooCollector {
//     return &fooCollector{
//         fooMetric: prometheus.NewDesc(
//             prometheus.BuildFQName(collector, "", "1_This_is_namew"),
//             "2 me description",
//             []string{"a","b"}, nil,
//         ),
//         barMetric: prometheus.NewDesc("bar_metric",
//             "Shows whether a bar has occurred in our cluster",
//             nil, nil,
//         ),
//         statusMetric: prometheus.NewDesc("gg_status",
//             "Shows status of golden gate running instance",
//             nil, nil,
//         ),
//     }
// }

// func (collector *fooCollector) Describe(ch chan<- *prometheus.Desc) {

//     //Update this section with the each metric you create for a given collector
//     ch <- collector.fooMetric
//     ch <- collector.barMetric
//     ch <- collector.statusMetric
// }

// //Collect implements required collect function for all promehteus collectors
// func (collector *fooCollector) Collect(ch chan<- prometheus.Metric) {

//     //Implement logic here to determine proper metric value to return to prometheus
//     //for each descriptor or call other functions that do so.
//     var metricValue float64




//     storage.GetGGRunningInstances(&mgroups)
//     // ifff => for 
//     //=====
//     // == check empty - clean 
//     storage.GetGGRunningExtractInstances(&mgroups,&mpointsofextract)    
//     storage.GetGGRunningMGRInstances(&mgroups,&mpointsofmgr)    
//     storage.GetGGRunningPMSRVRInstances(&mgroups,&mpointsofpmsrvr) 

//     if 1 == 1 {
//         metricValue += rand.Float64()
//     }


//     //Write latest value for each metric in the prometheus metric channel.
//     //Note that you can pass CounterValue, GaugeValue, or UntypedValue types here.
//     m1 := prometheus.MustNewConstMetric(collector.fooMetric, prometheus.GaugeValue, metricValue,[]string{"a","b"}...)
//     m2 := prometheus.MustNewConstMetric(collector.barMetric, prometheus.GaugeValue, metricValue)

//     var metric float64
//     metric, _= strconv.ParseFloat( mpointsofextract[0].Process.Status, 64)




//     ch <- m1
//     ch <- m2
//     ch <- prometheus.MustNewConstMetric(collector.statusMetric, prometheus.GaugeValue, metric)
// }