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

    statusMetric                        *prometheus.Desc
    extract_trail_iowc_Metric           *prometheus.Desc
    extract_trail_iowb_Metric           *prometheus.Desc
    extract_trail_max_bytes_Metric      *prometheus.Desc
    trail_rba_Metric                    *prometheus.Desc
    trail_seq_Metric                    *prometheus.Desc
    //== PUMP
    pump_trail_iorc_Metric           *prometheus.Desc
    pump_trail_iorb_Metric           *prometheus.Desc

}


func NewGoldenGateCollector() *GoldenGateCollector {
    return &GoldenGateCollector{
        statusMetric: prometheus.NewDesc(
            prometheus.BuildFQName(collector, "", "status"),
            "Shows status of golden gate instances _type 2:EXTRACT 4:PUMP 14:PMSRVR 1:MANAGER _status 3:running 6:stopped 8:append",
            []string{"mgr_host","group_name","type"}, nil,
        ),
        extract_trail_iowc_Metric: prometheus.NewDesc(
            prometheus.BuildFQName(collector, "", "extract_io_write_count"),
            "Extract Trail Output _ io write count",
            []string{"trail_name","trail_path","hostname","group_name"}, nil,
        ),
        extract_trail_iowb_Metric: prometheus.NewDesc(
            prometheus.BuildFQName(collector, "", "extract_io_write_bytes"),
            "Extract Trail Output _ io write bytes",
            []string{"trail_name","trail_path","hostname","group_name"}, nil,
        ),
        trail_rba_Metric: prometheus.NewDesc(
            prometheus.BuildFQName(collector, "", "trail_rba"),
            "Trail Output _ trail_rba _ current bytes size of trail",
            []string{"trail_name","trail_path","hostname","group_name"}, nil,
        ),
        trail_seq_Metric: prometheus.NewDesc(
            prometheus.BuildFQName(collector, "", "trail_seq"),
            "Trail Output _ trail_seq _ rotate times",
            []string{"trail_name","trail_path","hostname","group_name"}, nil,
        ),
        extract_trail_max_bytes_Metric: prometheus.NewDesc(
            prometheus.BuildFQName(collector, "", "extract_trail_max_bytes"),
            "Trail Output _ extract_trail_max_bytes _ Max size of a trail can be reach before rotate",
            []string{"trail_name","trail_path","hostname","group_name"}, nil,
        ),
        //==== PUMP 
        pump_trail_iorc_Metric: prometheus.NewDesc(
            prometheus.BuildFQName(collector, "", "pump_io_read_count"),
            "PUMP Trail Output _ io read count",
            []string{"trail_name","trail_path","hostname","group_name"}, nil,
        ),
        pump_trail_iorb_Metric: prometheus.NewDesc(
            prometheus.BuildFQName(collector, "", "pump_io_read_bytes"),
            "PUMP Trail Output _ io read bytes",
            []string{"trail_name","trail_path","hostname","group_name"}, nil,
        ),

    }
}

func (collector *GoldenGateCollector) Describe(ch chan<- *prometheus.Desc) {
    ch <- collector.statusMetric 
    ch <- collector.extract_trail_iowc_Metric 
    ch <- collector.extract_trail_iowb_Metric 
    ch <- collector.trail_rba_Metric 
    ch <- collector.trail_seq_Metric 
    ch <- collector.extract_trail_max_bytes_Metric 
    //== PUMP
    ch <- collector.pump_trail_iorc_Metric 
    ch <- collector.pump_trail_iorb_Metric 
}

func (collector *GoldenGateCollector) Collect(ch chan<- prometheus.Metric) {

    var (
        mgroups entities.MGroups
        mpointsofextract []entities.MpointsOfExtract
        mpointsofmgr entities.MpointsOfMGR
        mpointsofpmsrvr entities.MpointsOfPMSRVR 
        mpointsofpump []entities.MpointsOfPump      
    )

    storage.GetGGRunningInstances(&mgroups)
    //=====
    storage.GetGGRunningExtractInstances(&mgroups,&mpointsofextract)   
    storage.GetGGRunningPumpInstances(&mgroups,&mpointsofpump)     
    storage.GetGGRunningMGRInstances(&mgroups,&mpointsofmgr)    
    storage.GetGGRunningPMSRVRInstances(&mgroups,&mpointsofpmsrvr) 

    getstatus(ch,collector,  &mpointsofmgr, &mpointsofextract ,&mpointsofpmsrvr,&mpointsofpump)

}

func getMetricValue(input string) float64 {
    var metric float64
    metric, _= strconv.ParseFloat( input, 64)
    return metric 
}


func getstatus( ch chan<- prometheus.Metric, collector *GoldenGateCollector,   
                mpointsofmgr *entities.MpointsOfMGR , 
                mpointsofextract *[]entities.MpointsOfExtract ,
                mpointsofpmsrvr *entities.MpointsOfPMSRVR ,
                mpointsofpump *[]entities.MpointsOfPump   ){

    // ===== MGR =======
    ch <- prometheus.MustNewConstMetric(collector.statusMetric, 
                                        prometheus.GaugeValue, 
                                        getMetricValue( mpointsofmgr.Process.Status),
                                        []string{ "GGserver1" ,  mpointsofmgr.Process.Name  , mpointsofmgr.Process.Type }...)
    // ===== Extract =======    
    for _,extract := range (*mpointsofextract){
        ch <- prometheus.MustNewConstMetric(collector.statusMetric, 
                                                prometheus.GaugeValue, 
                                                getMetricValue( extract.Process.Status),
                                                []string{ "GGserver1" ,  extract.Process.Name  , extract.Process.Type }...)
        for _,trail := range (extract.Process.TrailOutput){
            //========== extract_io_write_count     "trail_name","trail_path","hostname","group_name"
            ch <- prometheus.MustNewConstMetric(collector.extract_trail_iowc_Metric, 
                                                prometheus.GaugeValue, 
                                                getMetricValue( trail.IoWriteCount ),
                                                []string{  trail.TrailName , trail.TrailPath , trail.Hostname , extract.Process.Name  }...)
            //========== extract_io_write_bytes     
            ch <- prometheus.MustNewConstMetric(collector.extract_trail_iowb_Metric, 
                                                prometheus.GaugeValue, 
                                                getMetricValue( trail.IoWriteBytes ),
                                                []string{  trail.TrailName , trail.TrailPath , trail.Hostname , extract.Process.Name  }...)
            //========== trail_rba_Metric     
            ch <- prometheus.MustNewConstMetric(collector.trail_rba_Metric, 
                                                prometheus.GaugeValue, 
                                                getMetricValue( trail.TrailRba ),
                                                []string{  trail.TrailName , trail.TrailPath , trail.Hostname , extract.Process.Name  }...)
            //========== trail_seq_Metric     
            ch <- prometheus.MustNewConstMetric(collector.trail_seq_Metric, 
                                                prometheus.GaugeValue, 
                                                getMetricValue( trail.TrailSeq ),
                                                []string{  trail.TrailName , trail.TrailPath , trail.Hostname , extract.Process.Name  }...)
            //========== extract_extract_trail_max_bytes_Metric     
            ch <- prometheus.MustNewConstMetric(collector.extract_trail_max_bytes_Metric, 
                                                prometheus.GaugeValue, 
                                                getMetricValue( trail.TrailMaxBytes ),
                                                []string{  trail.TrailName , trail.TrailPath , trail.Hostname , extract.Process.Name  }...)
                                                
        }
    }
    // ===== PUMP =======    
    for _,pump := range (*mpointsofpump){
        ch <- prometheus.MustNewConstMetric(collector.statusMetric, 
                                            prometheus.GaugeValue, 
                                            getMetricValue( pump.Process.Status),
                                            []string{ "GGserver1" ,  pump.Process.Name  , pump.Process.Type }...)
        // === Trail in                                      
        ch <- prometheus.MustNewConstMetric(collector.trail_rba_Metric, 
                                            prometheus.GaugeValue, 
                                            getMetricValue( pump.Process.TrailInput.TrailRba ),
                                            []string{  pump.Process.TrailInput.TrailName , pump.Process.TrailInput.TrailPath , "GGserver1" , pump.Process.Name  }...) 

        ch <- prometheus.MustNewConstMetric(collector.trail_seq_Metric,
                                            prometheus.GaugeValue, 
                                            getMetricValue( pump.Process.TrailInput.TrailSeq ),
                                            []string{  pump.Process.TrailInput.TrailName , pump.Process.TrailInput.TrailPath , "GGserver1" , pump.Process.Name  }...)

        // === Trail out 

        for _,trailout := range (pump.Process.TrailOutput){
            //========== extract_io_write_count     "trail_name","trail_path","hostname","group_name"
            ch <- prometheus.MustNewConstMetric(collector.extract_trail_iowc_Metric, 
                                                prometheus.GaugeValue, 
                                                getMetricValue( trailout.IoWriteCount ),
                                                []string{  trailout.TrailName , trailout.TrailPath , trailout.Hostname , pump.Process.Name  }...)
            //========== extract_io_write_bytespump            
            ch <- prometheus.MustNewConstMetric(collector.extract_trail_iowb_Metric, 
                                                prometheus.GaugeValue, 
                                                getMetricValue( trailout.IoWriteBytes ),
                                                []string{  trailout.TrailName , trailout.TrailPath , trailout.Hostname , pump.Process.Name  }...)
            //========== trail_rba_Metric     
            ch <- prometheus.MustNewConstMetric(collector.trail_rba_Metric, 
                                                prometheus.GaugeValue, 
                                                getMetricValue( trailout.TrailRba ),
                                                []string{  trailout.TrailName , trailout.TrailPath , trailout.Hostname , pump.Process.Name  }...)
            //========== trail_seq_Metric     
            ch <- prometheus.MustNewConstMetric(collector.trail_seq_Metric, 
                                                prometheus.GaugeValue, 
                                                getMetricValue( trailout.TrailSeq ),
                                                []string{  trailout.TrailName , trailout.TrailPath , trailout.Hostname , pump.Process.Name  }...)
            //========== extract_extract_trail_max_bytes_Metric     
            ch <- prometheus.MustNewConstMetric(collector.extract_trail_max_bytes_Metric, 
                                                prometheus.GaugeValue, 
                                                getMetricValue( trailout.TrailMaxBytes ),
                                                []string{  trailout.TrailName , trailout.TrailPath , trailout.Hostname , pump.Process.Name  }...)
                                                
        }
    }
    // ===== PMSRVR =======    
    ch <- prometheus.MustNewConstMetric(collector.statusMetric, 
                                                prometheus.GaugeValue, 
                                                getMetricValue( mpointsofpmsrvr.Process.Status),
                                                []string{ "GGserver1" ,  mpointsofpmsrvr.Process.Name  , mpointsofpmsrvr.Process.Type }...)
}

