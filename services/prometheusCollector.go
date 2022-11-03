package services

import (

    
    "github.com/prometheus/client_golang/prometheus"
)

type fooCollector struct {
    fooMetric *prometheus.Desc
    barMetric *prometheus.Desc
}





//You must create a constructor for you collector that
//initializes every descriptor and returns a pointer to the collector
func newFooCollector() *fooCollector {
    return &fooCollector{
        fooMetric: prometheus.NewDesc("foo_metric",
            "Shows whether a foo has occurred in our cluster",
            nil, nil,
        ),
        barMetric: prometheus.NewDesc("bar_metric",
            "Shows whether a bar has occurred in our cluster",
            nil, nil,
        ),
    }
}

//Each and every collector must implement the Describe function.
//It essentially writes all descriptors to the prometheus desc channel.
func (collector *fooCollector) Describe(ch chan<- *prometheus.Desc) {

    //Update this section with the each metric you create for a given collector
    ch <- collector.fooMetric
    ch <- collector.barMetric
}

//Collect implements required collect function for all promehteus collectors
func (collector *fooCollector) Collect(ch chan<- prometheus.Metric) {

    //Implement logic here to determine proper metric value to return to prometheus
    //for each descriptor or call other functions that do so.
    var metricValue float64
    if 1 == 1 {
        metricValue += rand.Float64()
    }

    //Write latest value for each metric in the prometheus metric channel.
    //Note that you can pass CounterValue, GaugeValue, or UntypedValue types here.
    m1 := prometheus.MustNewConstMetric(collector.fooMetric, prometheus.GaugeValue, metricValue)
    m2 := prometheus.MustNewConstMetric(collector.barMetric, prometheus.GaugeValue, metricValue)
    m1 = prometheus.NewMetricWithTimestamp(time.Now().Add(-time.Hour), m1)
    m2 = prometheus.NewMetricWithTimestamp(time.Now(), m2)
    ch <- m1
    ch <- m2
}