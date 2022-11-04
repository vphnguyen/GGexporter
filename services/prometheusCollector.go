package services

import (
    // "time"
    // log "github.com/sirupsen/logrus"
    "strconv"
    "GGexporter/storage"
    "GGexporter/entities"
    "github.com/prometheus/client_golang/prometheus"
    //"math/rand"
    // "fmt"
)
const  collector = "GoldenGate"
type GoldenGateCollector struct{

    fooMetric       *prometheus.Desc
    barMetric       *prometheus.Desc
    statusMetric    *prometheus.Desc
}




var (
    mgroups entities.MGroups
    mpointsofextract []entities.MpointsOfExtract
    mpointsofmgr entities.MpointsOfMGR
    mpointsofpmsrvr entities.MpointsOfPMSRVR
    
)




func NewGoldenGateCollector() *GoldenGateCollector {
    return &GoldenGateCollector{
        fooMetric: prometheus.NewDesc(
            prometheus.BuildFQName(collector, "a", "1_This_is_namew"),
            "2 me description",
            []string{"a","b"}, nil,
        ),
        barMetric: prometheus.NewDesc("bar_metric",
            "Shows whether a bar has occurred in our cluster",
            nil, nil,
        ),
        statusMetric: prometheus.NewDesc("gg_status",
            "Shows status of golden gate running instance",
            nil, nil,
        ),
    }
}

func (collector *GoldenGateCollector) Describe(ch chan<- *prometheus.Desc) {

    ch <- collector.fooMetric
    ch <- collector.barMetric
    ch <- collector.statusMetric
}

func (collector *GoldenGateCollector) Collect(ch chan<- prometheus.Metric) {


    var metricValue float64

    storage.GetGGRunningInstances(&mgroups)
    // ifff => for 
    //=====
    // == check empty - clean 
    storage.GetGGRunningExtractInstances(&mgroups,&mpointsofextract)    
    storage.GetGGRunningMGRInstances(&mgroups,&mpointsofmgr)    
    storage.GetGGRunningPMSRVRInstances(&mgroups,&mpointsofpmsrvr) 

  


    //Write latest value for each metric in the prometheus metric channel.
    //Note that you can pass CounterValue, GaugeValue, or UntypedValue types here.
    m1 := prometheus.MustNewConstMetric(collector.fooMetric, prometheus.GaugeValue, metricValue,[]string{"a","b"}...)
    m2 := prometheus.MustNewConstMetric(collector.barMetric, prometheus.GaugeValue, metricValue)

    var metric float64
    metric, _= strconv.ParseFloat( mpointsofextract[0].Process.Status, 64)




    ch <- m1
    ch <- m2
    ch <- prometheus.MustNewConstMetric(collector.statusMetric, prometheus.GaugeValue, metric)
}